// Package httpserver 提供带健康检查与通用中间件的 Fiber 应用引导封装（骨架）。
package httpserver

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// New 创建一个 Fiber 应用，已挂载 recover/logger 中间件与 /healthz 健康检查。
func New(serviceName string) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:               serviceName,
		DisableStartupMessage: true,
	})

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("/healthz", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok", "service": serviceName})
	})

	return app
}
