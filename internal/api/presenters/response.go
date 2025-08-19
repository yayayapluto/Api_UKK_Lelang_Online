package presenters

import "github.com/gofiber/fiber/v2"

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Content *T     `json:"content"`
	Error   error  `json:"error"`
}

func SuccessResponse[T any](c *fiber.Ctx, statusCode int, message string, content *T) error {
	return c.Status(statusCode).JSON(Response[T]{
		Success: true,
		Message: message,
		Content: content,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string, error error) error {
	return c.Status(statusCode).JSON(Response[any]{
		Success: true,
		Message: message,
		Error:   error,
	})
}
