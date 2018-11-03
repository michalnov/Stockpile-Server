package authentication

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const charset = "GHIJK456789abOPQRScdeUVmXYZfghijkLMNT0123lmnopYZ012qrsCDEtuv7xyzABF"

func NotImplemented() {

}

func randomString(lenght int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make([]byte, lenght)
	for i := range out {
		out[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(out)
}
