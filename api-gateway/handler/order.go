package handler

import (
	"context"
	"log"
	"strconv"

	pb "github.com/Mubinabd/restaurant_api-gateway/genproto"

	"github.com/gin-gonic/gin"
)

// @Router 				/order/create [POST]
// @Summary 			Create Order
// @Description		 	This api create order
// @Tags 				ORDER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param data 			body pb.Order true "Void"
// @Success 201 		{object} pb.Void
// @Failure 400 		string Error
func (h *HandlerStruct) CreateOrder(c *gin.Context) {
	var req pb.Order
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.OrderClient.CreateOrder(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, "Order created")
}

// @Router 				/order/{id} [GET]
// @Summary 			GET ORDER
// @Description		 	This api GETS order by id
// @Tags 				ORDER
// @Accept 				json
// @Produce 			json
// @Security    		BearerAuth
// @Param 			    id path string true "ORDER ID"
// @Success 200			{object} pb.Order
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetOrder(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")

	req.Id = id
	// log.Println(id, req)
	order, err := h.Clients.OrderClient.GetOrder(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, order)
}

// @Router 				/order/{id} [PUT]
// @Summary 			Update Order
// @Description		 	This api logs order in
// @Tags 				ORDER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param id 			path string true "ById"
// @Param menu_item_id  query string false "Menu Item ID"
// @Param reservation_id query string false "Reservation ID"
// @Param quantity      query string false "Quantity"
// @Success 201 		{object} pb.Void
// @Failure 400 		string Error
func (h *HandlerStruct)UpdateOrder(c *gin.Context) {
	
	id := c.Param("id")
	MenuItemId := c.Query("menu_item_id")
	ReservationId := c.Query("reservation_id")
	QuantityStr := c.Query("quantity")
	quantity,_:= strconv.Atoi(QuantityStr)
	req := pb.Order{
		Id: id,
        MenuItemId: MenuItemId,
        ReservationId: ReservationId,
        Quantity: int32(quantity),
	}
	_, err := h.Clients.OrderClient.UpdateOrder(c.Request.Context(), &req)
	if err != nil {
		c.JSON(400, gin.H{"Error when updating order": err.Error()})
		return
	}
	c.JSON(200, "Order updated")
}


// @Router 				/order/{id} [DELETE]
// @Summary 			Delete Order
// @Description		 	This api logs order in
// @Tags 				ORDER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "ORDER ID"
// @Success 201 		{object} pb.Void
// @Failure 400 		string Error
func (h *HandlerStruct) DeleteOrder(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")

	req.Id = id
	log.Println(id, &req)
	order, err := h.Clients.OrderClient.DeleteOrder(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, order)
}

// @Router 				/order/all [GET]
// @Summary 			Get all Orders
// @Description		 	This api logs order in
// @Tags 				ORDER
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Success 201 		{object} pb.Orders
// @Failure 400 		string Error
func (h *HandlerStruct) GetAllOrders(c *gin.Context) {
	orders, err := h.Clients.OrderClient.GetAllOrders(context.Background(), &pb.Void{})
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, orders)
}
