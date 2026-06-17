// Package router 注册 API 网关的 HTTP 路由（骨架）。
package router

import "github.com/gofiber/fiber/v3"

// Register 在 Fiber 应用上挂载网关路由。
// 真实实现里这里会做协议转换并转发到后端 gRPC 服务，骨架阶段返回占位响应。
func Register(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/ping", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "pong", "gateway": "api-gateway"})
	})
}
