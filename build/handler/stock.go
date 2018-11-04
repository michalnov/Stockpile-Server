package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var dbconf string = "root:sql@tcp(127.0.0.1:3306)/michalstock"

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
	Stockunits []StockUnit `json:"stockunits,omitempty"`
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res StockList

	req := loginRequest{}
	_ = json.NewDecoder(r.Body).Decode(&req)
	if req.Token == "" {
		w.WriteHeader(http.StatusForbidden)
		response := json.NewEncoder(w).Encode(res)
		fmt.Fprintf(w, response)
	} else if !scanUsers(req.Token) {
		w.WriteHeader(http.StatusForbidden)
		response := json.NewEncoder(w).Encode(res)
		fmt.Fprintf(w, response)
	} else {
		w.WriteHeader(http.StatusOK)
		db, err := sql.Open("mysql", dbconf)
		if err != nil {
			fmt.Fprintf(w, "{\"Error\": 500}")
			panic(err.Error())
		}

		response := json.NewEncoder(w).Encode(res)
		fmt.Fprintf(w, response)
	}

}
