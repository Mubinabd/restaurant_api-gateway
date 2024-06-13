package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/Mubinabd/restaurant_api-gateway/genproto"
	"github.com/gin-gonic/gin"
)

// @Router 				/admin/menu/create [POST]
// @Summary 			Create Menu
// @Description		 	This api create menu
// @Tags 				MENU
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param data 			body genproto.Menu true "Menu"
// @Success 201 		{object} string "menu created successfully"
// @Failure 400 		string Error
func (h *HandlerStruct) CreateMenu(c *gin.Context) {
	var req genproto.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.MenuClient.CreateMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "menu created successfully"})
}

// @Router 				/admin/menu/update [PUT]
// @Summary 			UPDATES MENU
// @Description		 	This api UPDATES menu 
// @Tags 				MENU
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param  reservation  body genproto.Menu true "MENU"
// @Success 200			{object} string "menu updated successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) UpdateMenu(c *gin.Context) {
	var req genproto.Menu
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.MenuClient.UpdateMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "menu updated successfully"})
}

// @Router 				/admin/menu/{id} [DELETE]
// @Summary 			DELETES Menu
// @Description		 	This api DELETES menu by id
// @Tags 				MENU
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "MENU ID"
// @Success 200			{object} string "menu deleted successfully"
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) DeleteMenu(c *gin.Context) {
	var req genproto.ById
	id := c.Param("id")
	req.Id = id
	_, err := h.Clients.MenuClient.DeleteMenu(context.Background(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "menu deleted successfully"})
}
// @Router 				/menu/{id} [GET]
// @Summary 			GET MENU
// @Description		 	This api GETS menu by id
// @Tags 				MENU
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    id path string true "MENU ID"
// @Success 200			{object} genproto.Menu
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetMenu(c *gin.Context) {
	var req genproto.ById
    id := c.Param("id")
    req.Id = id
    menu, err := h.Clients.MenuClient.GetMenu(context.Background(), &req)
    if err!= nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, menu)
}

// @Router 				/menu/all [GET]
// @Summary 			GETS ALL MENUS
// @Description		 	This api GETS ALL menus by id
// @Tags 				MENU
// @Accept 				json
// @Produce 			json
// @Security            BearerAuth
// @Param 			    rest_id query string false "RESTAURANT ID"
// @Param               price_from query string false "PriceFrom"
// @Param               price_to query string false "PriceTo"
// @Success 200			{object} genproto.Reservation
// @Failure 400 		string Error
// @Failure 404 		string Error
func (h *HandlerStruct) GetAllMenus(c *gin.Context) {
	rest_id := c.Query("rest_id")
	price_fromStr := c.Query("price_from")
	price_toStr := c.Query("price_to")

	var price_from, price_to float32
	var err error

	if price_fromStr != "" {
		var priceFromFloat64 float64
		priceFromFloat64, err = strconv.ParseFloat(price_fromStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price_from"})
			return
		}
		price_from = float32(priceFromFloat64)
	}

	if price_toStr != "" {
		var priceToFloat64 float64
		priceToFloat64, err = strconv.ParseFloat(price_toStr, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid price_to"})
			return
		}
		price_to = float32(priceToFloat64)
	}

	req := &genproto.MenuFilter{
		RestaurantId: rest_id,
		PriceFrom:    price_from,
		PriceTo:      price_to,
	}

	menus, err := h.Clients.MenuClient.GetMenus(context.Background(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, menus)
}