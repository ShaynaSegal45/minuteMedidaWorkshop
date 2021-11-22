package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"workshop/exercises/ex8/fact"
	"workshop/exercises/ex8/facthttp"
	"workshop/exercises/ex8/inmem"
	"workshop/exercises/ex8/providers/mentalfloss"
)

func main() {
	factsRepo := inmem.NewFactRepository()
	mentalflossProvider := mentalfloss.NewProvider()
	service := fact.NewService(factsRepo, mentalflossProvider)

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc() //defined but not invoked

	fmt.Println("start here")
	service.UpdateFactsWithTicker(ctx, service.UpdateFacts)
	fmt.Println("end here")

	handler := facthttp.NewFactsHandler(factsRepo)

	http.HandleFunc("/ping", handler.Ping)
	http.HandleFunc("/api/facts", handler.Facts)

	log.Println("starting server")
	err := http.ListenAndServe(":9007", nil)
	if err != nil {
		log.Fatal(err)
	}
}
