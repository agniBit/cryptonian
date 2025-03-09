package server

import (
	"context"
	"fmt"
	"time"

	"github.com/agniBit/cryptonian/internal/logger"
	"github.com/agniBit/cryptonian/internal/router"
	"github.com/agniBit/cryptonian/internal/storage/s3"
	"github.com/agniBit/cryptonian/internal/websocket"
	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

const shutdownTimeout = 5 * time.Second

type deps struct {
	binanceWS *websocket.BinanceWebSocket
}

type Server struct {
	app  *fiber.App
	cfg  *cfg.Config
	deps deps
}

func New(cfg *cfg.Config) *Server {
	s3.InitS3(cfg)

	return &Server{
		cfg:  cfg,
		app:  setupFiberApp(cfg),
		deps: deps{},
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.deps.binanceWS = websocket.NewBinanceWebSocket(ctx, s.cfg)
	s.deps.binanceWS.Start()

	router := router.NewRouter(s.app, s.cfg)
	router.RegisterRoutes()

	return s.app.Listen(":" + s.cfg.Server.Port)
}

func (s *Server) Shutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
	defer cancel()

	if err := s.app.ShutdownWithContext(ctx); err != nil {
		return fmt.Errorf("server shutdown error: %w", err)
	}

	s.deps.binanceWS.Close()
	return nil
}

func setupFiberApp(cfg *cfg.Config) *fiber.App {
	app := fiber.New()
	setupMiddlewares(app, cfg)
	return app
}

func setupMiddlewares(app *fiber.App, cfg *cfg.Config) {
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			err, ok := e.(error)
			if !ok {
				err = fmt.Errorf("panic: %v", e)
			}
			logger.Error(c.Context(), "Recovered from panic", err, nil)
		},
	}))

	app.Use(cors.New(getCORSConfig(cfg)))
}

func getCORSConfig(cfg *cfg.Config) cors.Config {
	corsCfg := cors.Config{
		AllowOrigins: "https://algosignal.in, https://www.algosignal.in, https://pre-prod.algosignal.in",
		AllowHeaders: "Content-Type, Authorization",
	}

	if cfg.Server.Environment == "pre-prod" || cfg.Server.Environment == "dev" {
		corsCfg.AllowOrigins += ", http://localhost:5173"
	}

	return corsCfg
}
