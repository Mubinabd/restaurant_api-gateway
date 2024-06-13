package api

import (
	"github.com/Mubinabd/restaurant_api-gateway/api-gateway/handler"
	_ "github.com/Mubinabd/restaurant_api-gateway/docs"
	"github.com/Mubinabd/restaurant_api-gateway/middleware"
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
		AllowOrigins:     []string{"http://localhost", "http://localhost:8090"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	r.Use(cors.New(corsConfig))

	restaurantAdmin := r.Group("/admin/restaurant")
	restaurantAdmin.Use(middleware.MiddlewareAdmin())
	{
		restaurantAdmin.POST("/create", h.CreateRestaurant)
		restaurantAdmin.PUT("/:id", h.UpdateRestaurant)
		restaurantAdmin.DELETE("/:id", h.DeleteRestaurant)
	}

	restaurant := r.Group("/restaurant")
	restaurant.Use(middleware.Middleware())
	{
		restaurant.GET("/:id", h.GetRestaurant)
		restaurant.GET("/all", h.GetAllRestaurants)
	}




	reservationAdmin := r.Group("/admin/reservation")
	reservationAdmin.Use(middleware.MiddlewareAdmin())
	{
        reservationAdmin.PUT("/:id", h.UpdateReservation)
        reservationAdmin.DELETE("/:id", h.DeleteReservation)
	}

	reservation := r.Group("/reservation")
	reservation.Use(middleware.Middleware())
	{
        reservation.GET("/:id", h.GetReservation)
        reservation.GET("/all", h.GetAllReservations)
		reservation.GET("/total/:id", h.GetTotalSum)
		reservation.POST("/:id/check", h.CheckReservation)
		reservation.POST("/:id/add-order", h.AddOrder)
		reservation.POST("/:id/payment", h.ReservationPayment)
	}


	menuAdmin := r.Group("/admin/menu")
	menuAdmin.Use(middleware.MiddlewareAdmin())
    {
		menuAdmin.POST("/create", h.CreateMenu)
		menuAdmin.PUT("/update", h.UpdateMenu)
        menuAdmin.DELETE("/:id", h.DeleteMenu)
	}
	
	menu := r.Group("/menu")
	menu.Use(middleware.Middleware())
	{
        menu.GET("/:id", h.GetMenu)
        menu.GET("/all", h.GetAllMenus)
	}



	order := r.Group("/order")
	order.Use(middleware.Middleware())
	{
		order.POST("/create", h.CreateOrder)
        order.GET("/:id", h.GetOrder)
        order.GET("/all", h.GetAllOrders)
        order.PUT("/:id", h.UpdateOrder)
        order.DELETE("/:id", h.DeleteOrder)
	}


	payment := r.Group("/payment")
	payment.Use(middleware.Middleware())
	{
		payment.POST("/create", h.CreatePayment)
        payment.GET("/:id", h.GetPayment)
        payment.GET("/all", h.GetAllPayments)
        payment.PUT("/update", h.UpdatePayment)
	}


	url := ginSwagger.URL("swagger/doc.json")
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	return r
}
