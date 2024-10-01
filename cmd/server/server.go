package server

import (
	"log"

	_ "github.com/PhamNgocHoangLong/greendeco-be.git/docs"
	"github.com/PhamNgocHoangLong/greendeco-be.git/pkg/configs"
	"github.com/PhamNgocHoangLong/greendeco-be.git/pkg/routes"
	"github.com/PhamNgocHoangLong/greendeco-be.git/pkg/validators"
	"github.com/PhamNgocHoangLong/greendeco-be.git/platform/database"
	"github.com/PhamNgocHoangLong/greendeco-be.git/platform/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Serve() {
	err := configs.LoadConfig()
	if err != nil {
		log.Fatal("error")
	}

	if err := database.ConnectDB(); err != nil {
		log.Panic(err)
	}

	if err := database.GetDB().Migrate(); err != nil {
		log.Panic(err)
	}

	if err := storage.ConnectStorage(); err != nil {
		log.Panic(err)
	}

	validators.AddProductQueryDecoder()
	validators.AddOrderQueryDecoder()

	app := fiber.New()
	app.Use(logger.New())

	corsApp := cors.ConfigDefault
	corsApp.AllowCredentials = true
	corsApp.AllowHeaders = "Origin, X-Requested-With, Content-Type, Accept, Authorization"
	app.Use(cors.New(corsApp))

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hellooo")
	})
	routes.SwaggerRoute(app)
	api := app.Group("/api/v1")
	routes.UserRoutes(api)
	routes.AuthRoutes(api)
	routes.NewMediaRouter(api).RegisterRoutes()
	routes.NewAdminRouter(api).RegisterRoutes()
	routes.CategoryRouter(api)

	routes.NewReviewRoutes(api).RegisterRoutes()
	routes.NewProductRouter(api).RegisterRoutes()
	routes.NewVariantRouter(api).RegisterRoute()
	routes.NewCartRouter(api).RegisterRoutes()
	routes.NewColorRouter(api).RegisterRoutes()
	routes.NewCouponRouter(api).RegisterRoutes()
	routes.NewOrderRouter(api).RegisterRoutes()
	routes.NewNotificationRouter(api).RegisterRoutes()
	routes.NewPaymentRouter(api).RegisterRoutes()
	if err := app.Listen(":8080"); err != nil {
		log.Fatal("not response")
	}
}
