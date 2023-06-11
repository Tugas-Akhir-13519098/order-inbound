package model

type TokopediaOrderNotif struct {
	FSID        int                `json:"fs_id"`
	OrderID     int                `json:"order_id"`
	Products    []TokopediaProduct `json:"products"`
	Recipient   TokopediaRecipient `json:"recipient"`
	ShopID      int                `json:"shop_id"`
	OrderStatus int                `json:"order_status"`
}

type TokopediaOrderChangeStatus struct {
	OrderStatus    int                `json:"order_status"`
	FSID           int                `json:"fs_id"`
	ShopID         int                `json:"shop_id"`
	OrderID        int                `json:"order_id"`
	ProductDetails []TokopediaProduct `json:"product_details"`
}

type TokopediaProduct struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Quantity   int    `json:"quantity"`
	TotalPrice int    `json:"total_price"`
}

type TokopediaRecipient struct {
	Name    string           `json:"name"`
	Phone   string           `json:"phone"`
	Address TokopediaAddress `json:"address"`
}

type TokopediaAddress struct {
	AddressFull string `json:"address_full"`
	District    string `json:"district"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	PostalCode  string `json:"postal_code"`
}
