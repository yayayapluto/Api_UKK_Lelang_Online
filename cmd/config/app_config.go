package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yayayapluto/api-ukk-online/internal/api/handlers"
	"github.com/yayayapluto/api-ukk-online/internal/api/routes"
	"github.com/yayayapluto/api-ukk-online/internal/utils"
	objectType "github.com/yayayapluto/api-ukk-online/pkg/object-type"
	"github.com/yayayapluto/api-ukk-online/pkg/organizer"
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
	organizerRepository := organizer.NewOrganizerRepository(db)

	// Services
	objectTypeService := objectType.NewObjectTypeService(objectTypeRepository)
	organizerService := organizer.NewOrganizerService(organizerRepository)

	// Handlers
	objectTypeHandler := handlers.NewObjectTypeHandler(objectTypeService, validator)
	organizerHandler := handlers.NewOrganizerHandler(organizerService, validator)

	routeConfig := routes.RouteConfig{
		App:               app,
		ObjectTypeHandler: objectTypeHandler,
		OrganizerHandler:  organizerHandler,
	}
	routeConfig.Setup()

	return app, nil
}
