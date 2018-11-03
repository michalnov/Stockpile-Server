package handler

import (
	"encoding/json"
	_ "fmt"
	"math/rand"
	"net/http"
	"time"
)

const charset = "GHIJK456789abOPQRScdeUVmXYZfghijkLMNT0123lmnopYZ012qrsCDEtuv7xyzABF"

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := loginRequest{}
	_ = json.NewDecoder(r.Body).Decode(&req)

}

func randomString(lenght int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]byte, lenght)
	for i := range out {
		out[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(out)
}
