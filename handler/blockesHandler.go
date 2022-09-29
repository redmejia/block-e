package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func (a *Api) BlockesHandler(w http.ResponseWriter, r *http.Request) {

	blockesByte, err := json.Marshal(a.Blockes.B)
	if err != nil {
		log.Println("err ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(blockesByte)
}
