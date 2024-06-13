package handler

import (
	"context"
	"net/http"

	genproto "github.com/Mubinabd/restaurant_api-gateway/genproto"
	"github.com/gin-gonic/gin"
)

// @Router 				/reservation/create [POST]
// @Summary 			Create Reservation
// @Description		 	This api create reservation
// @Tags 				RESERVATION
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param data 			body genproto.ReservationCreate true "ReservationCreate"
// @Success 201 		{object} string "reservation created successfully"
// @Failure 400 		string Error
func (h *HandlerStruct) CreateReservation(c *gin.Context) {
	var req genproto.ReservationCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.ReservationClient.CreateReservation(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "reservation created successfully"})
}

// @Router 				/reservation/{id} [GET]
// @Summary 			GET RESERVATION
// @Description		 	This api GETS reservation by id
// @Tags 				RESERVATION
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "RESERVATION ID"
// @Success 200			{object} genproto.Reservation
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetReservation(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	reservation, err := h.Clients.ReservationClient.GetReservation(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, reservation)
}

// @Router 				/admin/reservation/{id} [PUT]
// @Summary 			UPDATES RESERVATION
// @Description		 	This api UPDATES reservation by id
// @Tags 				RESERVATION
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param  id           path string true "Reservation ID"
// @Param  reservation  body genproto.Reservation true "RESERVATION"
// @Success 200			{object} string "reservation updated successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) UpdateReservation(c *gin.Context) {
	id := c.Param("id")
	var req genproto.Reservation
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	reser := genproto.ReservationCreate{
		Id: id,
		UserId: req.UserId,
        Status: req.Status,
        RestaurantId: req.RestaurantId,
		ReservationTime: req.ReservationTime,
	}
	_, err := h.Clients.ReservationClient.UpdateReservation(context.Background(), &reser)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "reservation updated successfully"})
}

// @Router 				/admin/reservation/{id} [DELETE]
// @Summary 			DELETES RESERVATION
// @Description		 	This api DELETES reservation by id
// @Tags 				RESERVATION
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "RESERVATION ID"
// @Success 200			{object} string "reservation deleted successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) DeleteReservation(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	_, err := h.Clients.ReservationClient.DeleteReservation(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "reservation deleted successfully"})
}

// @Router 				/reservation/all [GET]
// @Summary 			GETS ALL RESERVATION
// @Description		 	This api GETS ALL reservation by id
// @Tags 				RESERVATION
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    rest_id query string false "RESTAURANT ID"
// @Param               time_from query string false "ReservationFrom"
// @Param               time_to query string false "ReservationTo"
// @Success 200			{object} genproto.Reservation
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetAllReservations(c *gin.Context) {
	rest_id := c.Query("rest_id")
	time_from := c.Query("time_from")
	time_to := c.Query("time_to")

	req := &genproto.FilterByTime{
		RestaurantId:    rest_id,
		ReservationFrom: time_from,
		ReservationTo:   time_to,
	}

	reservations, err := h.Clients.ReservationClient.GetAllReservation(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

// @Router 				/reservation/total/{id} [GET]
// @Summary 			GET TOTAL SUM FOR  RESERVATION
// @Description		 	This api GETS TOTAL PAYCHECK FOR reservation by id
// @Tags 				RESERVATION
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "RESERVATION ID"
// @Success 200			{object} genproto.Total
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetTotalSum(c *gin.Context) {
	var req genproto.ById
    id := c.Param("id")
    req.Id = id
    total, err := h.Clients.ReservationClient.GetTotalSum(context.Background(), &req)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, total)
}
// @Router 				/reservation/{id}/check [POST]
// @Summary 			CHECKS IF THE IT IS AVAILABLE FOR RESERVATION
// @Description		 	This api CHECKS IF THE IT IS AVAILABLE FOR RESERVATION
// @Tags 				RESERVATION
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    DATA body genproto.ResrvationTime true "RESERVATION Check"
// @Success 200			{object} string "this reservation time is not reserved yet!"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) CheckReservation(c *gin.Context) {
	var req genproto.ResrvationTime
	if err := c.ShouldBindJSON(&req); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	_, err := h.Clients.ReservationClient.CheckReservation(context.Background(), &req)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"message": "this reservation time is not reserved yet!"})
}
// @Router 				/reservation/{id}/add-order [POST]
// @Summary 			ADDS ORDER TO RESERVATION
// @Description		 	This api DELETES reservation by id
// @Tags 				RESERVATION
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "RESERVATION ID"
// @Param  data         body genproto.ReservationOrder true "order data"
// @Success 200			{object} string "order created successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) AddOrder(c *gin.Context) {
	id := c.Param("id")
	var ord genproto.ReservationOrder
	if err := c.ShouldBindJSON(&ord); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	req := genproto.Order{
		ReservationId: id,
		MenuItemId: ord.MenuItemId,
		Quantity: ord.Quantity,
	}
	_, err := h.Clients.OrderClient.CreateOrder(context.Background(),&req)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"message": "order created successfully"})
}

// @Router 				/reservation/{id}/payment [POST]
// @Summary 			PAYS FOR RESERVATION
// @Description		 	This api PAYS for reservation
// @Tags 				RESERVATION
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "RESERVATION ID"
// @Param  data         body genproto.ReservationPayment true "PAYMENT data"
// @Success 200			{object} genproto.Payment
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) ReservationPayment(c *gin.Context) {
	id := c.Param("id")
	var respayment genproto.ReservationPayment
	if err := c.ShouldBindJSON(&respayment); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	req := genproto.PaymentCreate{
        ReservationId: id,
		PaymentMethod: respayment.PaymentMethod,
        Amount: respayment.Amount,
    }
	res, err := h.Clients.PaymentClient.CreatePayment(context.Background(),&req)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	c.JSON(http.StatusOK, res)
}