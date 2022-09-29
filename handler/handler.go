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

	newBlock := a.Blocke.GenerateNewBlock(&a.Blockes.B[len(a.Blockes.B)-1], bpm.Bpm)

	if blocke.CheckBlock(newBlock, &a.Blockes.B[len(a.Blockes.B)-1]) {
		newBlockChain := append(a.Blockes.B, *newBlock)
		a.Blockes.ReplaceChain(newBlockChain)
	}

	a.Blockes.AddBlock(*newBlock)

	log.Println(newBlock)

	newBlockByte, err := json.Marshal(&newBlock)
	if err != nil {
		log.Println("error ", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(newBlockByte))
}
