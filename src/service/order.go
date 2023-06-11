package service

import (
	"order-inbound/src/model"
	"order-inbound/src/repository"
	"order-inbound/src/util"
	"strconv"
)

type OrderService interface {
	ReceiveTokopediaOrderNotif(order *model.TokopediaOrderNotif) error
	ReceiveTokopediaOrderChangeStatus(order *model.TokopediaOrderChangeStatus) error
	ReceiveShopeeOrderNotif(order *model.ShopeeOrderNotif) error
}

type orderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{orderRepository: orderRepository}
}

func (o *orderService) ReceiveTokopediaOrderNotif(order *model.TokopediaOrderNotif) error {
	kafkaOrderMessage := util.ConvertTokopediaOrderNotifToKafkaOrderMessage(*order)
	err := o.orderRepository.PublishToKafka(strconv.Itoa(order.OrderID), &kafkaOrderMessage)

	return err
}

func (o *orderService) ReceiveTokopediaOrderChangeStatus(order *model.TokopediaOrderChangeStatus) error {
	kafkaOrderMessage := util.ConvertTokopediaOrderChangeStatusToKafkaOrderMessage(*order)
	err := o.orderRepository.PublishToKafka(strconv.Itoa(order.OrderID), &kafkaOrderMessage)

	return err
}

func (o *orderService) ReceiveShopeeOrderNotif(order *model.ShopeeOrderNotif) error {
	partnerID := 1 // Hard coded

	shopeeOrderDetail, err := o.orderRepository.GetShopeeOrderDetail(order.ShopID, partnerID, order.Data.Ordersn)
	if err != nil {
		return err
	}

	kafkaOrderMessage := util.ConvertShopeeOrderDetailToKafkaOrderMessage(*shopeeOrderDetail, order.ShopID)
	err = o.orderRepository.PublishToKafka(order.Data.Ordersn, &kafkaOrderMessage)

	return err
}
