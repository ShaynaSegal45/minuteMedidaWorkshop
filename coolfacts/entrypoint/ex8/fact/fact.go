package fact

import (
	"context"
	"fmt"
	"log"
	"reflect"
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

	factSet :=make(map[Fact]bool)
	fmt.Println(factSet)

	allFacts:=s.repository.GetAll()
	for _,val :=range allFacts {
		if _, ok := factSet[val];ok{
			break
		}
		factSet[val]=true
	}

	for _,val :=range mfFacts{
		if _, ok := factSet[val];ok{
			break
		}
		s.repository.Add(val)
		factSet[val]=true
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

func clear(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}
func contains(facts []Fact, f Fact) bool {
	for _, v := range facts {
		if v == f {
			return true
		}
	}

	return false
}
