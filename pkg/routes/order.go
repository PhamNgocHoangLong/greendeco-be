package routes

import (
	"github.com/PhamNgocHoangLong/greendeco-be.git/app/controller"
	"github.com/PhamNgocHoangLong/greendeco-be.git/pkg/middlewares"
	"github.com/gofiber/fiber/v2"
)

type orderRoutes struct {
	app fiber.Router
}

func NewOrderRouter(app fiber.Router) *orderRoutes {
	return &orderRoutes{
		app: app.Group("/order"),
	}
}

func (r *orderRoutes) RegisterRoutes() {
	r.publicOrderRoutes()
	r.privateOrderRoutes()
}

func (r *orderRoutes) publicOrderRoutes() {
}

func (r *orderRoutes) privateOrderRoutes() {
	r.app.Use(middlewares.JWTProtected())
	r.app.Post("/", controller.CreateOrderFromCart)
	r.app.Get("/all/", middlewares.AdminProtected, controller.GetAllOrders)
	r.app.Get("/:id", controller.GetOrderById)
	r.app.Get("/:id/product/", controller.GetOrderProductByOrderId)
	r.app.Get("/", controller.GetOrderByToken)
	r.app.Get("/:id/total", controller.GetTotalOrderById)
	r.app.Put("/:id", controller.UpdateOrderStatus)
}
