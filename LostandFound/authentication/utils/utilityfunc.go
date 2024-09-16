package utils

import (
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
	"time"
)

func ParseJSONBody(r *http.Request, a interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, a)
	if err != nil {
		return
	}
}

func GenerateSalt(length int) []byte {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	salt := make([]byte, length)
	for i := range salt {
		salt[i] = letters[r.Intn(len(letters))]
	}
	return salt
}
