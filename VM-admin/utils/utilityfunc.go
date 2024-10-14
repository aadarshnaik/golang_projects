package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aadarshnaik/vm-admin/models"
)

func ParseJSONBody(r *http.Request, vm *models.VM) error {

	err := json.NewDecoder(r.Body).Decode(&vm)
	if err != nil {
		log.Println("Invalid request body")
		return err
	}
	return nil
}

func HandleError(str string, err error) {
	if err != nil {
		log.Println(str)
		log.Println(err)
	}
}
