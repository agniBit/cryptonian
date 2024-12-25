package server

import (
	"context"
	"os"
	"time"

	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/agniBit/cryptonian/internal/router"
	"github.com/agniBit/cryptonian/internal/storage/s3"
	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/robfig/cron"
)

func ListenAndServe(ctx context.Context, cfg *cfg.Config, quit chan os.Signal) {

	logger.InitNewRelic(cfg)

	app := fiber.New()
	addCorsAndRecoverMiddleware(cfg, app)

	logger.Init(cfg)
	s3.InitS3(cfg)

	cron := cron.NewWithLocation(time.UTC)

	defer func() {
		defer cron.Stop()

		logger.Info(ctx, "Flushing logs...", nil)
		logger.Flush()
		s3.SyncLogsToS3(ctx, "./temp/")
	}()

	routerService := router.NewRouter(app, cfg)
	routerService.RegisterRoutes()

	time.Sleep(time.Second * 1)

	// start the server
	err := app.Listen(":" + cfg.Server.Port)
	if err != nil {
		logger.Error(ctx, "error in starting server", err, nil)
		panic(err)
	}

	<-quit
	logger.Warn(ctx, "Shutting down server...", nil)

	// Attempt a graceful shutdown
	if err := app.Shutdown(); err != nil {
		logger.Error(ctx, "Server forced to shutdown", err, nil)
	} else {
		logger.Info(ctx, "Server shutdown gracefully", nil)
	}

	ctx.Done()
	// wait for websocket to close
	<-time.After(time.Second * 5)
	logger.Info(ctx, "Server exiting", nil)
}

func addCorsAndRecoverMiddleware(cfg *cfg.Config, app *fiber.App) {
	// recover stack trace handler on panic
	stackTraceHandler := func(c *fiber.Ctx, e interface{}) {
		if err, ok := e.(error); ok {
			logger.Error(c.Context(), "recovered from panic", err, nil)
		} else {
			logger.Error(c.Context(), "recovered from panic", nil, map[string]interface{}{"panic": e})
		}
	}

	recoverCfg := recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: stackTraceHandler,
	}
	app.Use(recover.New(recoverCfg))

	corsCfg := cors.Config{
		AllowOrigins: "https://algosignal.in, https://www.algosignal.in, https://pre-prod.algosignal.in, https://pre-release.daos0x4tnrymg.amplifyapp.com",
		AllowHeaders: "Content-Type, Authorization",
	}

	if cfg.Server.Enviroment == "pre-prod" || cfg.Server.Enviroment == "dev" {
		corsCfg.AllowOrigins += ", http://localhost:5173"
	}

	app.Use(cors.New(corsCfg))
}
