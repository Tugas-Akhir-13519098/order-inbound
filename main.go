package main

import (
	"fmt"
	"net/http"
	"order-inbound/config"
	"order-inbound/src/controller"
	"order-inbound/src/repository"
	"order-inbound/src/service"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/segmentio/kafka-go"
)

func NewKafkaWriter(cfg *config.Config) *kafka.Writer {
	config := kafka.WriterConfig{
		Brokers: []string{fmt.Sprintf("%s:%s", cfg.KafkaHost, cfg.KafkaPort)},
		Topic:   cfg.KafkaOrderTopic,
	}
	writer := kafka.NewWriter(config)

	return writer
}

func NewRetryClient() *retryablehttp.Client {
	retryClient := retryablehttp.NewClient()
	retryClient.ErrorHandler = func(resp *http.Response, err error, numTries int) (*http.Response, error) {
		resp.Body.Close()
		return resp, err
	}

	return retryClient
}

func main() {
	cfg := config.Get()

	// kafka
	writer := NewKafkaWriter(&cfg)

	// retryable http
	retryClient := retryablehttp.NewClient()

	orderRepository := repository.NewOrderRepository(writer, retryClient)

	orderService := service.NewOrderService(orderRepository)

	orderController := controller.NewOrderController(orderService)

	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		// product route
		orderRoute := v1.Group("/order")

		orderRoute.POST("/tokopedia/notif/", orderController.ReceiveTokopediaOrderNotif)
		orderRoute.POST("/tokopedia/status/", orderController.ReceiveTokopediaOrderChangeStatus)
		orderRoute.POST("/shopee/", orderController.ReceiveShopeeOrderNotif)
	}

	router.Run(fmt.Sprintf("%s:%d", cfg.RESTHost, cfg.RESTPort))
}
