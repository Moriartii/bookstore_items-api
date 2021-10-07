package items

type Item struct {
	Id               string      `json:"id"`
	Seller           int64       `json:"seller"`
	Title            string      `json:"title"`
	Description      Description `json: "description"`
	Picture          []Picture   `json:"picture"`
	Video            string      `json:"video"`
	Price            float32     `json:"price"`
	AvailableQantity int64       `json:"available_qantity"`
	SoldQantity      int64       `json:"sold_quantity"`
	ItemStatus       string      `json:"status"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}

type Picture struct {
	Id  int64  `json:"id"`
	Url string `json:"url"`
}
