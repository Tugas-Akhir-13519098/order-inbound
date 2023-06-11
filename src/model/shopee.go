package model

type ShopeeOrderNotif struct {
	Data   ShopeeOrderNotifData `json:"data"`
	ShopID int                  `json:"shop_id"`
}

type ShopeeOrderNotifData struct {
	Ordersn string `json:"ordersn"`
	Status  string `json:"status"`
}

type ShopeeOrderDetail struct {
	RequestID string          `json:"request_id"`
	Error     string          `json:"error"`
	Message   string          `json:"message"`
	Response  ShopeeOrderList `json:"response"`
	Warning   string          `json:"warning"`
}

type ShopeeOrderList struct {
	OrderList []ShopeeOrder `json:"order_list"`
}

type ShopeeOrder struct {
	OrderSN          string          `json:"order_sn"`
	TotalAmount      float32         `json:"total_amount"`
	OrderStatus      string          `json:"order_status"`
	RecipientAddress ShopeeRecipient `json:"recipient_address"`
	ItemList         []ShopeeItem    `json:"item_list"`
}

type ShopeeRecipient struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	District    string `json:"district"`
	City        string `json:"city"`
	State       string `json:"state"`
	Region      string `json:"region"`
	Zipcode     string `json:"zipcode"`
	FullAddress string `json:"full_address"`
}

type ShopeeItem struct {
	ItemID                 int     `json:"item_id"`
	ItemName               string  `json:"item_name"`
	ModelQuantityPurchased int     `json:"model_quantity_purchased"`
	ModelOriginalPrice     float32 `json:"model_quantity_price"`
}
