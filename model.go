package main 

type product struct {
	ID  int  `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
}