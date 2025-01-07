package model

// 2025;01;06;1;13.28;13.28;
type Price struct {
	Year   int     `json:"year"`
	Month  int     `json:"month"`
	Day    int     `json:"nonce"`
	Hour   int     `json:"hour"`
	Price1 float32 `json:"price1"`
	Price2 float32 `json:"price2"`
}
