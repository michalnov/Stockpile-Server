package handler

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func notImplemented() {

}

type StockUnit struct {
	Token     string `json:"token,omitempty"`
	Name      string `json:"name,omitempty"`
	Quantity  int    `json:"quantity,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Origin    string `json:"origin,omitempty"`
}

type StockList struct {
	Stockunits StockUnit `json:"stockunits,omitempty"`
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {

}
