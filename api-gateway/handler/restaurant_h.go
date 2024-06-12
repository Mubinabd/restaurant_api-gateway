package handler

import (
	"context"
	"log"

	pb "github.com/Mubinabd/restaurant_api-gateway/genproto"

	"github.com/gin-gonic/gin"
)

// @Router 				/restaurant/create [POST]
// @Summary 			Create Restaurant
// @Description		 	This api create restaurant
// @Tags 				RESTAURANT
// @Accept 				json
// @Produce 			json
// @Param data 			body pb.CreateRestaurantReq true "CreateRestaurantReq"
// @Success 201 		{object} pb.Void
// @Failure 400 		string Error
func (h *HandlerStruct)CreateRestaurant(c *gin.Context) {
	var req pb.CreateRestaurantReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.RestaurantClient.CreateRestaurant(context.Background(),&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Restaurant created")
}

// @Router 				/restaurant/{id} [GET]
// @Summary 			GET RESTAURANT
// @Description		 	This api GETS restaurant by id
// @Tags 				RESTAURANT
// @Accept 				json
// @Produce 			json
// @Security    		BearerAuth
// @Param 			    id path string true "RESTAURANT ID"
// @Success 200			{object} pb.Restaurant
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct)GetRestaurant(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")

	req.Id = id
	// log.Println(id,req)
	user, err := h.Clients.RestaurantClient.GetRestaurant(context.Background(),&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}

// @Router 				/restaurant/{id} [PUT]
// @Summary 			Update Restaurant
// @Description		 	This api logs  restaurant in
// @Tags 				RESTAURANT
// @Accept 				json
// @Produce 			json
// @Param id 			path string true "Restaurant ID"
// @Param restaurant    body pb.Restaurant true "Restaurant "
// @Success 201 		{object} pb.Void
// @Failure 400 		string Error
func (h *HandlerStruct)UpdateRestaurant(c *gin.Context) {
	id := c.Param("id")
	var rest_req pb.Restaurant
	if err := c.ShouldBindJSON(&rest_req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
        return
	}
	req := pb.CreateRestaurantReq{
		Id: id,
        Name: rest_req.Name,
        Address: rest_req.Address,
        PhoneNumber: rest_req.PhoneNumber,
        Description: rest_req.Description,
	}

	log.Println(id,req)
	_, err := h.Clients.RestaurantClient.UpdateRestaurant(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"Error when updating restaurant": err.Error()})
		return
	}
	c.JSON(200, "Restaurant updated")
}

// @Router 				/restaurant/{id} [DELETE]
// @Summary 			Delete Restaurant
// @Description		 	This api logs  restaurant in
// @Tags 				RESTAURANT
// @Accept 				json
// @Produce 			json
// @Param 			    id path string true "RESTAURANT ID"
// @Success 201 		{object} pb.Void
// @Failure 400 		string Error
func (h *HandlerStruct)DeleteRestaurant(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")

	req.Id = id
	log.Println(id,&req)
	user, err := h.Clients.RestaurantClient.DeleteRestaurant(context.Background(),&req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, user)
}
// @Router 				/restaurant/all [GET]
// @Summary 			Get all Restaurants
// @Description		 	This api logs restaurant in
// @Tags 				RESTAURANT
// @Accept 				json
// @Produce 			json
// @Param 			    address query string false "RESTAURANT Address"
// @Success 200 		{object} pb.Restaurants
// @Failure 400 		string Error
func (h *HandlerStruct) GetAllRestaurants(c *gin.Context) {
	var addressFilter pb.AddressFilter
	address := c.Query("address")
	addressFilter.Address = address
	// log.Println(address,addressFilter)
	restaurants, err := h.Clients.RestaurantClient.GetAllRestaurants(c.Request.Context(), &addressFilter)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, restaurants)
}
