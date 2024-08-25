package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadRequestBody(r *http.Request, a interface{}) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, a)
	if err != nil {
		return
	}
}
