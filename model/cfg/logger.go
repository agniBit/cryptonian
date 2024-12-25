package cfg

type Logger struct {
	LogLevel     string `mapstructure:"LOG_LEVEL"`
	SyncLogsToS3 bool   `mapstructure:"SYNC_LOGS_TO_S3"`
}
