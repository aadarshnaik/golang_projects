package controllers

import (
	"log"
	"net/http"

	"github.com/aadarshnaik/vm-admin/models"
	"github.com/aadarshnaik/vm-admin/utils"
)

func GetHeartbeat(w http.ResponseWriter, r *http.Request) {

	var vm models.VM
	err := utils.ParseJSONBody(r, &vm)
	utils.HandleError("", err)

	log.Println("Heartbeat received for VM: " + vm.Hostname)
	log.Println("Details: ", vm)

	log.Println("Heartbeat received")
	w.WriteHeader(http.StatusOK)
}
