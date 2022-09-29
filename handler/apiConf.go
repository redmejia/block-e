package handler

import "cadena/blocke/blocke"

type Api struct {
	Blocke  *blocke.Blocke // optional not need
	Blockes *blocke.Blockes
}
