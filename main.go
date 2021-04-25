package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/gauravtripathi001/go-microservices/handlers"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//Create the handlers
	//hh := handlers.NewHello(l)
	//gh := handlers.NewGoodbye(l)
	ph := handlers.NewProducts(l)

	sm := http.NewServeMux()
	//sm.Handle("/", hh)
	//sm.Handle("/goodbye",gh)
	sm.Handle("/",ph)

	s := &http.Server{
		Addr: ":9090",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func(){
		l.Println("Starting server ")

		err := s.ListenAndServe()
		if err != nil{
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan,os.Interrupt)
	signal.Notify(sigChan,os.Kill)

	sig := <- sigChan
	l.Println("Received terminate, graceful shutdown",sig)
	tc, _ := context.WithTimeout(context.Background(),30*time.Second)
	s.Shutdown(tc)

}
