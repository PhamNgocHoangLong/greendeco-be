package routes

import (
	"github.com/PhamNgocHoangLong/greendeco-be.git/app/controller"
	"github.com/PhamNgocHoangLong/greendeco-be.git/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

type couponRoutes struct {
	app fiber.Router
}

func NewCouponRouter(app fiber.Router) *couponRoutes {
	return &couponRoutes{
		app: app.Group("/coupon"),
	}
}

func (r *couponRoutes) RegisterRoutes() {
	r.publicCouponRoutes()
	r.privateCouponRoutes()
}

func (r *couponRoutes) publicCouponRoutes() {
	r.app.Get("/:id", controller.GetCouponById)
	r.app.Get("/code/:code", controller.GetCouponByCode)
}

func (r *couponRoutes) privateCouponRoutes() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Use(middlewares.AdminProtected)
	r.app.Post("/", controller.CreateCoupon)
	r.app.Put("/:id", controller.UpdateCouponById)
	r.app.Delete("/:id", controller.DeleteCouponById)
}
