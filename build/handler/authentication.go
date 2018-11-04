package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type user struct {
	Username string
	Token    string
}

var Users = make([]user, 0)

const charset = "GHIJK456789abOPQRScdeUVmXYZfghijkLMNT0123lmnopYZ012qrsCDEtuv7xyzABF"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	req := loginRequest{}
	_ = json.NewDecoder(r.Body).Decode(&req)
	if req.Username == "tester" && req.Password == "test123" {
		res := loginRequest{}
		res.Token = randomString(64)
		Users = append(Users, user{Username: req.Username, Token: res.Token})
		response := json.NewEncoder(w).Encode(res)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, response)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "\"token\":\"\"")
	}

}

func scanUsers(toekn string) bool {
	for index, element := range Users {
		if element.Token == toekn {
			return true
		}
	}
	return false
}

func randomString(lenght int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]byte, lenght)
	for i := range out {
		out[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(out)
}
