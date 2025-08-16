package routes

import "github.com/gofiber/fiber/v2"

type RouteConfig struct {
	App *fiber.App
}

func (r *RouteConfig) Setup() {
	r.ping()
}

func (r *RouteConfig) ping() {
	r.App.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("OK!")
	})
}
