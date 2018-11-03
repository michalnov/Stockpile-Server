package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	auth "github.com/michalnov/Stockpile-Server/authentication"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := loginRequest{}
	_ = json.NewDecoder(r.Body).Decode(&req)

}
