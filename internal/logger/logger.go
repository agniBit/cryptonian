package logger

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/newrelic/go-agent/v3/integrations/logcontext-v2/nrzap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logFielRotater *lumberjack.Logger
var logger *zap.Logger
var initOnce = &sync.Once{}

func Init(cfg *cfg.Config) {
	if logger == nil {
		initOnce.Do(func() {
			if logger == nil {
				newLogger(cfg)
			}
		})
	}
}

func addCustomCallerInfo() zap.Field {
	// Skip 2 levels to get the caller of the logger function
	_, file, line, ok := runtime.Caller(4)
	if !ok {
		return zap.String("caller", "unknown")
	}
	return zap.String("caller", fmt.Sprintf("%s:%d", filepath.Base(file), line))
}

func newLogger(cfg *cfg.Config) {
	if cfg == nil {
		config := zap.NewDevelopmentEncoderConfig()
		logger = zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(config),
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		))
		return
	}

	// Configure encoder and log level
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	logLevel := zapcore.InfoLevel
	if cfg.Logger != nil && cfg.Logger.LogLevel != "" {
		if level, err := zapcore.ParseLevel(cfg.Logger.LogLevel); err == nil {
			logLevel = level
		}
	}

	// Set up console and file logging
	fileEncoder := zapcore.NewJSONEncoder(config)
	cores := []zapcore.Core{
		zapcore.NewCore(fileEncoder, zapcore.AddSync(os.Stdout), logLevel),
	}

	if cfg.Server.LogFile != "" {
		logFielRotater = &lumberjack.Logger{
			Filename:   cfg.Server.LogFile,
			MaxSize:    50,
			MaxBackups: 2,
			MaxAge:     10,
			Compress:   true,
		}
		cores = append(cores, zapcore.NewCore(fileEncoder, zapcore.AddSync(logFielRotater), zapcore.DebugLevel))
	}

	core := zapcore.NewTee(cores...)
	backgroundCore, err := nrzap.WrapBackgroundCore(core, GetApp())
	if err == nil || !errors.Is(err, nrzap.ErrNilApp) {
		core = backgroundCore
	} else {
		fmt.Printf("Failed to integrate New Relic: %v\n", err)
	}

	logger = zap.New(core, zap.AddStacktrace(zapcore.ErrorLevel))
}

func mergeFields(fields ...map[string]interface{}) map[string]interface{} {
	merged := make(map[string]interface{})
	for _, m := range fields {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}

// flattenStruct recursively flattens the JSON-like structure.
func flattenStruct(data interface{}, prefix string, result map[string]interface{}) {
	if data == nil {
		return
	}

	defer func(data interface{}, prefix string, result map[string]interface{}) {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic", zap.Any("panic", r), zap.Any("data", data), zap.String("prefix", prefix), zap.Any("result", result))
		}
	}(data, prefix, result)

	v := reflect.ValueOf(data)
	if v.IsZero() {
		return
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.IsZero() {
		return
	}
	switch v.Kind() {
	case reflect.Map:
		for _, key := range v.MapKeys() {
			flattenStruct(v.MapIndex(key).Interface(), prefix+key.String()+".", result)
		}
	case reflect.Struct:
		// handle time.Time separately
		if v.Type().PkgPath() == "time" && v.Type().Name() == "Time" {
			result[strings.TrimSuffix(prefix, ".")] = v.Interface().(time.Time).Format(time.RFC3339)
			return
		}
		for i := 0; i < v.NumField(); i++ {
			field := v.Type().Field(i)
			fieldName := field.Name

			// Skip unexported fields
			if fieldName[0:1] != strings.ToUpper(fieldName[0:1]) {
				continue
			}

			fieldValue := v.Field(i).Interface()
			if prefix[len(prefix)-1] != '.' {
				prefix += "."
			}
			flattenStruct(fieldValue, prefix+fieldName+".", result)
		}
	case reflect.Array, reflect.Slice:
		// if struct is a slice, iterate over each element and call flattenStruct
		// otherwise just use the value as is
		if v.Len() == 0 {
			break
		}

		// check if the slice is a slice of structs
		v0 := v.Index(0)
		if v0.Kind() == reflect.Ptr {
			v0 = v0.Elem()
		}
		if v0.Kind() == reflect.Struct {
			for i := 0; i < v.Len(); i++ {
				flattenStruct(v.Index(i).Interface(), strings.TrimSuffix(prefix, ".")+fmt.Sprintf("[%d].", i), result)
			}
			break
		} else {
			// if not a slice of structs, add the slice as is
			result[strings.TrimSuffix(prefix, ".")] = v.Interface()
		}
	default:
		// Base case: add the value to the map with the full prefixed key
		result[strings.TrimSuffix(prefix, ".")] = v.Interface()
	}
}

func ParseAndFlattenMessage(prefix, jsonString string) map[string]interface{} {
	defer func(prefix, jsonString string) {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic", zap.Any("panic", r))
		}
	}(prefix, jsonString)

	var jsonData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonString), &jsonData); err != nil {
		return map[string]interface{}{prefix: jsonString}
	}
	flattenedData := make(map[string]interface{})
	flattenStruct(jsonData, prefix, flattenedData)
	return flattenedData
}

func FlattenStruct(prefix string, data interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	flattenStruct(data, prefix, result)
	return result
}

func logMessage(ctx context.Context, level zapcore.Level, message string, err error, flatFields ...map[string]interface{}) {
	defer func(level zapcore.Level, message string, err error, flatFields ...map[string]interface{}) {
		if r := recover(); r != nil {
			logger.Error("Recovered from panic", zap.Any("panic", r), zap.String("message", message), zap.Error(err), zap.Any("level", level), zap.String("message", message), zap.Any("fields", flatFields))
		}
	}(level, message, err, flatFields...)

	fields := mergeFields(flatFields...)
	zapFields := getCommonFields(ctx)
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	if err != nil {
		zapFields = append(zapFields, zap.Error(err))
	}
	switch level {
	case zapcore.InfoLevel:
		logger.Info(message, zapFields...)
	case zapcore.WarnLevel:
		logger.Warn(message, zapFields...)
	case zapcore.ErrorLevel:
		logger.Error(message, zapFields...)
	case zapcore.FatalLevel:
		logger.Fatal(message, zapFields...)
	case zapcore.DebugLevel:
		logger.Debug(message, zapFields...)
	}
}

func Debug(ctx context.Context, message string, flatFields ...map[string]interface{}) {
	logMessage(ctx, zapcore.DebugLevel, message, nil, flatFields...)
}

func Info(ctx context.Context, message string, flatFields ...map[string]interface{}) {
	logMessage(ctx, zapcore.InfoLevel, message, nil, flatFields...)
}

func Warn(ctx context.Context, message string, flatFields ...map[string]interface{}) {
	logMessage(ctx, zapcore.WarnLevel, message, nil, flatFields...)
}

func Error(ctx context.Context, message string, err error, flatFields ...map[string]interface{}) {
	logMessage(ctx, zapcore.ErrorLevel, message, err, flatFields...)
}

func Fatal(ctx context.Context, message string, err error, fields map[string]interface{}) {
	Flush()
	logMessage(ctx, zapcore.FatalLevel, message, err, fields)
}

func getCommonFields(ctx context.Context) []zap.Field {
	zapFields := make([]zap.Field, 0)

	zapFields = append(zapFields, addCustomCallerInfo())

	return zapFields
}

func Flush() {
	_ = logger.Sync()
	logFielRotater.Rotate()
}
