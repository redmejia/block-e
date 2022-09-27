package handler

import (
	"cadena/blocke/blocke"
	"encoding/json"
	"log"
	"net/http"
)

func (a *Api) HandleNewBlock(w http.ResponseWriter, r *http.Request) {

	var bpm blocke.BPM

	err := json.NewDecoder(r.Body).Decode(&bpm)
	if err != nil {
		log.Println("error here ", err)
	}
	defer r.Body.Close()

}
