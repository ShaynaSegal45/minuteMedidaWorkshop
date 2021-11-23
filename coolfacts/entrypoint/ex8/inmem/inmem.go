package inmem

import "workshop/exercises/ex8/fact"

type factRepository struct {
	Facts []fact.Fact
}

func NewFactRepository() *factRepository {
	return &factRepository{}
}

func (r *factRepository) Add(f fact.Fact)  {
	r.Facts=append([]fact.Fact{f},r.Facts...)
}

func (r *factRepository) GetAll() []fact.Fact {
	return r.Facts
}
