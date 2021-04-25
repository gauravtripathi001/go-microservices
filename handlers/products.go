package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Products struct{
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products{
	return &Products{l}
}

func (p*Products) ServeHTTP(rw http.ResponseWriter,h *http.ResponseWriter){
	lp := data.getProducts()
	d, err := json.Marshal(lp)
	if err != nil{
		http.Error(rw,"Unable to marshal json",http.StatusInternalServerError)
	}
}
