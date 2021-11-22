package fact

import (
	"context"
	"fmt"
	"log"
	"time"
)

type Fact struct {
	Image       string
	Description string
}

type Provider interface {
	Facts() ([]Fact, error)
}

type Repository interface {
	Add(f Fact)
	GetAll() []Fact
}

type service struct {
	provider Provider
	repository    Repository
}

func NewService(r Repository, p Provider) *service {
	serv:=service {provider: p,
		repository: r,
	}
	return &serv
}

func (s *service) UpdateFacts() error {

	mfFacts, err := s.provider.Facts()
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range mfFacts {
		s.repository.Add(val)
	}
	return nil
}
func (s *service)UpdateFactsWithTicker(ctx context.Context, updateFunc func() error) {
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


