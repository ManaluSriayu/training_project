package talk_training

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func ReadTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	response := GetTalks()
	json.NewEncoder(w).Encode(response)
}

func WriteTalks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	response := "halo"
	json.NewEncoder(w).Encode(response)
}

func InsertTalking(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/vnd.api+json")
	productID := r.FormValue("product_id")
	message := r.FormValue("message")
	createTime := r.FormValue("createTime")
	userId := r.FormValue("userId")

	// pID := r.URL.Query().Get()
	productIDInt, errParse := strconv.ParseInt(productID, 10, 64)
	userIdInt, errParse := strconv.ParseInt(userId, 10, 64)

	if errParse != nil {
		log.Println(errParse)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	InsertTalks(productIDInt, userIdInt, message, createTime)
	w.WriteHeader(http.StatusOK)
	return
}
