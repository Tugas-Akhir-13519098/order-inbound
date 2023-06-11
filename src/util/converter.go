package util

import (
	"encoding/json"
	"io"
	"order-inbound/src/model"
)

func ConvertTokopediaOrderNotifToKafkaOrderMessage(order model.TokopediaOrderNotif) model.KafkaOrderMessage {
	var products []model.KafkaProduct
	totalPrice := float32(0)

	for _, product := range order.Products {
		totalPrice += float32(product.TotalPrice)
		kafkaProduct := model.KafkaProduct{
			TokopediaProductID: product.ID,
			ProductName:        product.Name,
			ProductPrice:       float32(product.TotalPrice),
			ProductQuantity:    product.Quantity,
		}
		products = append(products, kafkaProduct)
	}

	kafkaOrderMessage := model.KafkaOrderMessage{
		Method:             model.CREATE,
		Products:           products,
		TokopediaOrderID:   order.OrderID,
		TokopediaShopID:    order.ShopID,
		TotalPrice:         totalPrice,
		CustomerName:       order.Recipient.Name,
		CustomerPhone:      order.Recipient.Phone,
		CustomerAddress:    order.Recipient.Address.AddressFull,
		CustomerDistrict:   order.Recipient.Address.District,
		CustomerCity:       order.Recipient.Address.City,
		CustomerProvince:   order.Recipient.Address.Province,
		CustomerCountry:    order.Recipient.Address.Country,
		CustomerPostalCode: order.Recipient.Address.PostalCode,
		OrderStatus:        model.ACCEPTED,
	}

	return kafkaOrderMessage
}

func ConvertTokopediaOrderChangeStatusToKafkaOrderMessage(order model.TokopediaOrderChangeStatus) model.KafkaOrderMessage {
	orderStatus := ConvertTokopediaOrderStatus(order.OrderStatus)
	kafkaOrderMessage := model.KafkaOrderMessage{
		Method:           model.UPDATE,
		TokopediaOrderID: order.OrderID,
		TokopediaShopID:  order.ShopID,
		OrderStatus:      orderStatus,
	}

	return kafkaOrderMessage
}

func ConvertTokopediaOrderStatus(orderStatus int) model.OrderStatus {
	if orderStatus < 100 {
		return model.CANCELLED
	} else if orderStatus >= 100 && orderStatus < 400 {
		return model.RECEIVED
	} else if orderStatus >= 400 && orderStatus < 700 {
		return model.ACCEPTED
	} else {
		return model.DONE
	}
}

func ConvertResponseToShopeeOrderDetail(body io.ReadCloser) (model.ShopeeOrderDetail, error) {
	respBody, _ := io.ReadAll(body)
	var shopeeOrderDetail model.ShopeeOrderDetail
	err := json.Unmarshal(respBody, &shopeeOrderDetail)

	return shopeeOrderDetail, err
}

func ConvertShopeeOrderDetailToKafkaOrderMessage(order model.ShopeeOrderDetail, shopID int) model.KafkaOrderMessage {
	method := model.CREATE
	orderStatus := ConvertShopeeOrderStatus(order.Response.OrderList[0].OrderStatus)
	if orderStatus != model.RECEIVED {
		method = model.UPDATE
	}

	var products []model.KafkaProduct
	for _, product := range order.Response.OrderList[0].ItemList {
		kafkaProduct := model.KafkaProduct{
			ShopeeProductID: product.ItemID,
			ProductName:     product.ItemName,
			ProductPrice:    product.ModelOriginalPrice,
			ProductQuantity: product.ModelQuantityPurchased,
		}
		products = append(products, kafkaProduct)
	}

	kafkaOrderMessage := model.KafkaOrderMessage{
		Method:             method,
		Products:           products,
		ShopeeOrderID:      order.Response.OrderList[0].OrderSN,
		ShopeeShopID:       shopID,
		TotalPrice:         order.Response.OrderList[0].TotalAmount,
		CustomerName:       order.Response.OrderList[0].RecipientAddress.Name,
		CustomerPhone:      order.Response.OrderList[0].RecipientAddress.Phone,
		CustomerAddress:    order.Response.OrderList[0].RecipientAddress.FullAddress,
		CustomerDistrict:   order.Response.OrderList[0].RecipientAddress.District,
		CustomerCity:       order.Response.OrderList[0].RecipientAddress.City,
		CustomerProvince:   order.Response.OrderList[0].RecipientAddress.State,
		CustomerCountry:    order.Response.OrderList[0].RecipientAddress.Region,
		CustomerPostalCode: order.Response.OrderList[0].RecipientAddress.Zipcode,
		OrderStatus:        orderStatus,
	}

	return kafkaOrderMessage
}

func ConvertShopeeOrderStatus(orderStatus string) model.OrderStatus {
	if orderStatus == "UNPAID" {
		return model.RECEIVED
	} else if orderStatus == "CANCELLED" {
		return model.CANCELLED
	} else if orderStatus == "COMPLETED" {
		return model.DONE
	} else {
		return model.ACCEPTED
	}
}
