package request

type Order struct {
	ProductId string `json:"productId"`
	Qty       int    `json:"qty"`
}
