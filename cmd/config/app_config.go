package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yayayapluto/api-ukk-online/internal/api/handlers"
	"github.com/yayayapluto/api-ukk-online/internal/api/routes"
	"github.com/yayayapluto/api-ukk-online/internal/utils"
	objectType "github.com/yayayapluto/api-ukk-online/pkg/object-type"
	"gorm.io/gorm"
)

func NewApp(db *gorm.DB) (*fiber.App, error) {
	utils.InitValidator()
	//validator := utils.Validate

	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
	})

	validator := utils.Validate

	// Repositories
	objectTypeRepository := objectType.NewObjectTypeRepository(db)

	// Services
	objectTypeService := objectType.NewObjectTypeService(objectTypeRepository)

	// Handlers
	objectTypeHandler := handlers.NewObjectTypeHandler(objectTypeService, validator)

	routeConfig := routes.RouteConfig{
		App:               app,
		ObjectTypeHandler: objectTypeHandler,
	}
	routeConfig.Setup()

	return app, nil
}
