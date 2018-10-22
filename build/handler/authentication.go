package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := loginRequest{}
	_ = json.NewDecoder(r.Body).Decode(&req)

}
