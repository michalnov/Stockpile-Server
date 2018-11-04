package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var dbconf string = "root:sql@tcp(127.0.0.1:3306)/michalstock"

var PersistList = make([]StockUnit, 0)

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

func InserStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := StockUnit{}
	_ = json.NewDecoder(r.Body).Decode(&req)
	if req.Token == "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "")
	} else if !scanUsers(req.Token) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "")
	} else {
		db, err := sql.Open("mysql", dbconf)
		if err != nil {
			fmt.Fprintf(w, "{\"Error\": 500}")
			panic(err.Error())
		}

		statement, err := db.Prepare("insert into stock(name, quantity, origin, recipient)values(?,?,?,?)")
		if err != nil {
			fmt.Fprintf(w, "{\"Error\": 500}")
			panic(err.Error())
		}
		_, err = statement.Exec(req.Name, req.Quantity, req.Origin, req.Recipient)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "")
			panic(err.Error())
		}
		PersistList = append(PersistList, StockUnit{Name: req.Name, Quantity: req.Quantity, Origin: req.Origin, Recipient: req.Recipient})
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "")
	}

}

func RemoveStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := StockUnit{}
	_ = json.NewDecoder(r.Body).Decode(&req)
	if req.Token == "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "")
	} else if !scanUsers(req.Token) {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "")
	} else {
		db, err := sql.Open("mysql", dbconf)
		if err != nil {
			fmt.Fprintf(w, "{\"Error\": 500}")
			panic(err.Error())
		}
		for i, elem := range PersistList {
			if elem.Name == req.Name && elem.Origin == req.Origin && elem.Recipient == req.Recipient {
				if elem.Quantity > req.Quantity {
					newquantity := elem.Quantity - req.Quantity
					statement, err := db.Prepare("update stock set quantity = ? where name = ? && origin = ? && recipient = ? && quantity = ?")
					if err != nil {
						fmt.Fprintf(w, "{\"Error\": 500}")
						panic(err.Error())
					}
					_, err = statement.Exec(newquantity, req.Name, req.Origin, req.Recipient, req.Quantity)
					if err != nil {
						w.WriteHeader(http.StatusForbidden)
						fmt.Fprintf(w, "")
						panic(err.Error())
					}
				} else if elem.Quantity == req.Quantity {
					statement, err := db.Prepare("delete from stock where name = ? && origin = ? && recipient = ? && quantity = ?")
					if err != nil {
						fmt.Fprintf(w, "{\"Error\": 500}")
						panic(err.Error())
					}
					_, err = statement.Exec(req.Name, req.Origin, req.Recipient, req.Quantity)
					if err != nil {
						w.WriteHeader(http.StatusForbidden)
						fmt.Fprintf(w, "")
						panic(err.Error())
					}
					PersistList = append(PersistList[:i], PersistList[i+1:]...)
				} else {
					w.WriteHeader(http.StatusForbidden)
					fmt.Fprintf(w, "")
				}
			}
		}
	}
}

func UpdateStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res StockList

	req := loginRequest{}
	_ = json.NewDecoder(r.Body).Decode(&req)
	if req.Token == "" {
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(res)
		//fmt.Fprintf(w, response)
	} else if !scanUsers(req.Token) {
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(res)
		//fmt.Fprintf(w, response)
	} else {
		w.WriteHeader(http.StatusOK)
		db, err := sql.Open("mysql", dbconf)
		if err != nil {
			fmt.Fprintf(w, "{\"Error\": 500}")
			panic(err.Error())
		}

		results, err := db.Query("select name, quantity, origin, recipient from stock")
		if err != nil {
			panic(err.Error())
		}

		for results.Next() {
			var swap StockUnit
			err = results.Scan(&swap.Name, &swap.Quantity, &swap.Origin, &swap.Recipient)
			res.Stockunits = append(res.Stockunits, swap)
		}
		PersistList = res.Stockunits
		_ = json.NewEncoder(w).Encode(res)
	}
}

func LoadStock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res StockList
	req := loginRequest{}
	_ = json.NewDecoder(r.Body).Decode(&req)
	if req.Token == "" {
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(res)
		//fmt.Fprintf(w, response)
	} else if !scanUsers(req.Token) {
		w.WriteHeader(http.StatusForbidden)
		_ = json.NewEncoder(w).Encode(res)
		//fmt.Fprintf(w, response)
	} else {
		w.WriteHeader(http.StatusOK)
		res.Stockunits = PersistList
		_ = json.NewEncoder(w).Encode(res)
	}

}
