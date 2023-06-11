package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"order-inbound/config"
	"order-inbound/src/model"
	"order-inbound/src/util"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/segmentio/kafka-go"
)

type OrderRepository interface {
	PublishToKafka(orderID string, order *model.KafkaOrderMessage) error
	GetShopeeOrderDetail(shopID int, partnerID int, ordersn string) (*model.ShopeeOrderDetail, error)
}

type orderRepository struct {
	writer      *kafka.Writer
	retryClient *retryablehttp.Client
}

func NewOrderRepository(writer *kafka.Writer, retryClient *retryablehttp.Client) OrderRepository {
	return &orderRepository{writer: writer, retryClient: retryClient}
}

func (o *orderRepository) PublishToKafka(orderID string, order *model.KafkaOrderMessage) error {
	// Convert order message to byte
	message, err := json.Marshal(order)
	if err != nil {
		return err
	}

	// Publish to order topic
	err = o.writer.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(orderID),
		Value: []byte(message),
	})
	if err != nil {
		return err
	}

	return nil
}

func (o *orderRepository) GetShopeeOrderDetail(shopID int, partnerID int, ordersn string) (*model.ShopeeOrderDetail, error) {
	cfg := config.Get()
	var shopeeOrderDetail model.ShopeeOrderDetail
	url := fmt.Sprintf("%sget_order_detail?order_sn_list=%s&shop_id=%d&partner_id=%d", cfg.ShopeeURL, ordersn, shopID, partnerID)

	resp, err := o.retryClient.Get(url)
	if err != nil {
		return &shopeeOrderDetail, err
	}
	shopeeOrderDetail, err = util.ConvertResponseToShopeeOrderDetail(resp.Body)
	if shopeeOrderDetail.Error != "" {
		err = fmt.Errorf(shopeeOrderDetail.Message)
	}

	return &shopeeOrderDetail, err
}
