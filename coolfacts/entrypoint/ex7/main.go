package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"workshop/exercises/ex7/fact"
	"workshop/exercises/ex7/facthttp"
	"workshop/exercises/ex7/mentalfloss"
)

var factRepo fact.Repo

func mfUpdate() error {

	mfFacts, err := mentalfloss.Mentalfloss{}.Facts()
	if err != nil {
		log.Fatal(err)
	}
	factRepo = fact.Repo{}
	for _, val := range mfFacts {
		factRepo.Add(val)
	}
	return nil
}

func updateFactsWithTicker(ctx context.Context, updateFunc func() error) {
	ticker := time.NewTicker(5 * time.Millisecond)
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				err := updateFunc()
				fmt.Println(err)

			}
		}
	}(ctx)

}

func main() {

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc() //defined but not invoked

	fmt.Println("start here")
	updateFactsWithTicker(ctx, mfUpdate)
	fmt.Println("end here")

	handlerer := facthttp.FactsHandler{
		FactRepo: &factRepo,
	}

	http.HandleFunc("/ping", handlerer.Ping)
	http.HandleFunc("/api/facts", handlerer.Facts)

	log.Println("starting server")
	err := http.ListenAndServe(":9006", nil)
	if err != nil {
		log.Fatal(err)
	}
}
