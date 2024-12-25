package router

import (
	"github.com/agniBit/cryptonian/model/cfg"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app *fiber.App
	cfg *cfg.Config
}

func NewRouter(app *fiber.App, cfg *cfg.Config) *Router {

	return &Router{
		app: app,
		cfg: cfg,
	}
}

func (r *Router) RegisterRoutes() {
	// health check
	r.app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("alive!!!") })
}
