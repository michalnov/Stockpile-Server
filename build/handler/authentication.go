package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const charset = "GHIJK456789abOPQRScdeUVmXYZfghijkLMNT0123lmnopYZ012qrsCDEtuv7xyzABF"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := loginRequest{}
	_ = json.NewDecoder(r.Body).Decode(&req)
	if req.Username == "tester" && req.Password == "test123" {
		res := loginRequest{}
		res.Token = randomString(64)
		response := json.NewEncoder(w).Encode(res)
		fmt.Fprintf(w, response)
	} else {
		fmt.Fprintf(w, "\"token\":\"\"")
	}

}

func randomString(lenght int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]byte, lenght)
	for i := range out {
		out[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(out)
}
