package routes

import (
	"github.com/PhamNgocHoangLong/greendeco-be.git/app/controller"
	"github.com/PhamNgocHoangLong/greendeco-be.git/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app fiber.Router) {
	auth := app.Group("/auth")
	publicAuthRouter(auth)
	privateAuthRouter(auth)
}

func publicAuthRouter(app fiber.Router) {
	app.Post("/register", controller.CreateUser)
	app.Post("/login", controller.Login)
	app.Post("/forgot-password", controller.ForgotPassword)
}

func privateAuthRouter(app fiber.Router) {
	app.Use(middlewares.JWTProtected())
	app.Put("/password", controller.UpdatePassword)
}
