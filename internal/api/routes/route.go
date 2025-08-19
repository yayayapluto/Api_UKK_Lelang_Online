package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yayayapluto/api-ukk-online/internal/api/handlers"
)

type RouteConfig struct {
	App               *fiber.App
	ObjectTypeHandler handlers.ObjectTypeHandler
}

func (r *RouteConfig) Setup() {
	r.ping()
	r.ObjectType()
}

func (r *RouteConfig) ping() {
	r.App.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON("OK!")
	})
}

func (r *RouteConfig) ObjectType() {
	group := r.App.Group("/api/v1/objectTypes")
	group.Get("/", r.ObjectTypeHandler.List)
	group.Get("/:id", r.ObjectTypeHandler.Get)
}
