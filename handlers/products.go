package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gauravtripathi001/go-microservices/data"
)

type Products struct{
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products{
	return &Products{l}
}

func (p*Products) ServeHTTP(rw http.ResponseWriter,h *http.ResponseWriter){
	lp := data.GetProducts()
	d, err := json.Marshal(lp)
	if err != nil{
		http.Error(rw,"Unable to marshal json",http.StatusInternalServerError)
	}
}
