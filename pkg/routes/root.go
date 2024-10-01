package routes

import (
	_ "github.com/PhamNgocHoangLong/greendeco-be.git/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoute(a fiber.Router) {
	a.Get("/docs/*", swagger.HandlerDefault)
	a.Get("/api/v1/", func(c *fiber.Ctx) error {
		return c.Redirect("/docs")
	})
}
