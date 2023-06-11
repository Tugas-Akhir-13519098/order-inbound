package controller

import (
	"net/http"
	"order-inbound/src/model"
	"order-inbound/src/service"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	ReceiveTokopediaOrderNotif(c *gin.Context)
	ReceiveTokopediaOrderChangeStatus(c *gin.Context)
	ReceiveShopeeOrderNotif(c *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) OrderController {
	return &orderController{orderService: orderService}
}

func (o *orderController) ReceiveTokopediaOrderNotif(c *gin.Context) {
	order := &model.TokopediaOrderNotif{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusCreated, gin.H{"message": err.Error(), "status": "failed"})
		return
	}

	err := o.orderService.ReceiveTokopediaOrderNotif(order)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusCreated, gin.H{"message": err.Error(), "status": "failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully received order!"})
}

func (o *orderController) ReceiveTokopediaOrderChangeStatus(c *gin.Context) {
	order := &model.TokopediaOrderChangeStatus{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusCreated, gin.H{"message": err.Error(), "status": "failed"})
		return
	}

	err := o.orderService.ReceiveTokopediaOrderChangeStatus(order)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusCreated, gin.H{"message": err.Error(), "status": "failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully received order change status!"})
}

func (o *orderController) ReceiveShopeeOrderNotif(c *gin.Context) {
	order := &model.ShopeeOrderNotif{}
	if err := c.ShouldBindJSON(&order); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		c.JSON(http.StatusCreated, gin.H{"message": err.Error(), "status": "failed"})
		return
	}

	err := o.orderService.ReceiveShopeeOrderNotif(order)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		c.JSON(http.StatusCreated, gin.H{"message": err.Error(), "status": "failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully received order change status!"})
}
