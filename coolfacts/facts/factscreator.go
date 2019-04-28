package facts

import (
	"io/ioutil"
	"net/http"
)

type Parser interface {
	ParseFromPolling(b []byte) ([]Fact, error)
	ParseFromCreate(b []byte) (Fact, error)
}

type Store interface {
	Get() []Fact
	Set(data []Fact)
}

type WriteError func (w http.ResponseWriter)

type FactCreator struct {
	writeError WriteError
	parser Parser
	store Store
}

func NewFactCreator(we WriteError, p Parser, s Store) *FactCreator {
	return &FactCreator{we, p, s}
}

func (f *FactCreator) PostFactHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		f.writeError(w)
		return
	}
	fact, err := f.parser.ParseFromCreate(b)
	if err != nil {
		f.writeError(w)
		return
	}
	f.WriteToCache(fact)
}

func (f *FactCreator) WriteToCache(fact Fact) {
	data := f.store.Get()
	data = append([]Fact{fact}, data...)
	f.store.Set(data)
}