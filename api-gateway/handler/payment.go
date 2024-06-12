package handler

import (
	"context"
	"net/http"
	"strconv"

	pb "github.com/Mubinabd/restaurant_api-gateway/genproto"
	"github.com/gin-gonic/gin"
)

// @Router 				/payment/create [POST]
// @Summary 			Creates Payment
// @Description		 	This api creates payment
// @Tags 				PAYMENT
// @Accept 				json
// @Produce 			json
// @Param data 			body pb.PaymentCreate true "PaymentCreate"
// @Success 201 		{object} pb.Payment
// @Failure 400 		string Error
func (h *HandlerStruct) CreatePayment(c *gin.Context) {
	var req pb.PaymentCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Clients.PaymentClient.CreatePayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, res)
}

// @Router 				/payment/{id} [GET]
// @Summary 			Gets Payment
// @Description		 	This api gets payment
// @Tags 				PAYMENT
// @Accept 				json
// @Produce 			json
// @Param id 			path string true "PaymentCreate"
// @Success 201 		{object} pb.Payment
// @Failure 400 		string Error
func (h *HandlerStruct) GetPayment(c *gin.Context) {
	var req pb.ById
	id := c.Param("id")
	req.Id = id
	payment, err := h.Clients.PaymentClient.GetPayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

// @Router 				/payment/update [POST]
// @Summary 			Updates Payment
// @Description		 	This api updates payment
// @Tags 				PAYMENT
// @Accept 				json
// @Produce 			json
// @Param data 			body pb.PaymentCreate true "PaymentCreate"
// @Success 201 		{object} pb.Payment
// @Failure 400 		string Error
func (h *HandlerStruct) UpdatePayment(c *gin.Context) {
	var req pb.PaymentCreate
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Clients.PaymentClient.UpdatePayment(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
// @Router 				/payment/all [GET]
// @Summary 			Gets all Payments
// @Description		 	This api gets all payments
// @Tags 				PAYMENT
// @Accept 				json
// @Produce 			json
// @Param payment_method  path string false "PaymentMethod"
// @Param amount_from   path string false "AmountFrom"
// @Param amount_to     path string false "AmountTo"
// @Param payment_status path string false "PaymentStatus"
// @Success 201 		{object} pb.Payment
// @Failure 400 		string Error
func (h *HandlerStruct) GetAllPayments(c *gin.Context) {
	pay_method := c.Query("payment_method")
	amount_fromStr := c.Query("amount_from")
	amount_toStr := c.Query("amount_to")
	payment_status := c.Query("payment_status")

	var amount_from, amount_to float32
	var err error

	if amount_fromStr != "" {
		var amountFromFloat64 float64
		amountFromFloat64, err = strconv.ParseFloat(amount_fromStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount_from"})
			return
		}
		amount_from = float32(amountFromFloat64)
	}

	if amount_toStr != "" {
		var amountToFloat64 float64
		amountToFloat64, err = strconv.ParseFloat(amount_toStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount_to"})
			return
		}
		amount_to = float32(amountToFloat64)
	}

	req := &pb.PaymentFilter{
		PaymentMethod: pay_method,
		AmountFrom:    amount_from,
		AmountTo:      amount_to,
		PaymentStatus: payment_status,
	}

	payments, err := h.Clients.PaymentClient.GetPayments(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, payments)
}
