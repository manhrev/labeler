package payos

type CreatePaymentLinkParams struct {
	OrderCode    int32
	Amount       int64
	Description  string
	CancelUrl    string
	ReturnUrl    string
	Items        []Item
	BuyerName    string
	BuyerEmail   string
	BuyerPhone   string
	BuyerAddress string
	ExpiredAt    int64
}

type Item struct {
	Name     string
	Quantity int
	Price    int
}
