package routes

import (
	"github.com/PhamNgocHoangLong/greendeco-be.git/app/controller"
	"github.com/gofiber/fiber/v2"
)

type MediaRoutes struct {
	app fiber.Router
}

func NewMediaRouter(app fiber.Router) *MediaRoutes {
	return &MediaRoutes{app: app.Group("/media")}
}

func (r *MediaRoutes) RegisterRoutes() {
	r.publicRouter()
}

func (r *MediaRoutes) publicRouter() {
	r.app.Post("/upload", controller.PostMedia)
}
