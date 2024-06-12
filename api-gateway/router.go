package api

import (
	"github.com/Mubinabd/restaurant_api-gateway/api-gateway/handler"
	_ "github.com/Mubinabd/restaurant_api-gateway/docs"
	"github.com/gin-contrib/cors"
	// "google.golang.org/genproto/googleapis/cloud/bigquery/reservation/v1"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Title Online Voting System Swagger UI
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewGin(h *handler.HandlerStruct) *gin.Engine {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"http://localhost", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))


	restaurant := r.Group("/restaurant")
	{
		restaurant.POST("/create", h.CreateRestaurant)
		restaurant.GET("/:id", h.GetRestaurant)
		restaurant.GET("/all", h.GetAllRestaurants)
		restaurant.PUT("/:id", h.UpdateRestaurant)
		restaurant.DELETE("/:id", h.DeleteRestaurant)
	}
	reservation := r.Group("/reservation")
	{
		reservation.POST("/create", h.CreateReservation)
        reservation.GET("/:id", h.GetReservation)
        reservation.GET("/all", h.GetAllReservations)
        reservation.PUT("/update", h.UpdateReservation)
        reservation.DELETE("/:id", h.DeleteReservation)
	}
	menu := r.Group("/menu")
	{
		menu.POST("/create", h.CreateMenu)
        menu.GET("/:id", h.GetMenu)
        menu.GET("/all", h.GetAllMenus)
        menu.PUT("/update", h.UpdateMenu)
        menu.DELETE("/:id", h.DeleteMenu)
	}
	order := r.Group("/order")
	{
		order.POST("/create", h.CreateOrder)
        order.GET("/:id", h.GetOrder)
        order.GET("/all", h.GetAllOrders)
        order.PUT("/:id", h.UpdateOrder)
        order.DELETE("/:id", h.DeleteOrder)
	}


	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
