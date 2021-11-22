package mentalfloss

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"workshop/exercises/ex8/fact"
)

type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

func (mf Provider) Facts() ([]fact.Fact, error) {
	resp, err := http.Get("https://mentalfloss.com/api/facts")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var mfFacts []struct {
		FactText     string `json:"Fact"`
		PrimaryImage string `json:"primaryImage"`
	}

	unmarshalErr := json.Unmarshal(body, &mfFacts)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}

	var facts []fact.Fact
	for _, v := range mfFacts {
		newFact := fact.Fact{
			Image:       v.PrimaryImage,
			Description: v.FactText,
		}
		facts = append(facts, newFact)
	}
	return facts, nil
}

