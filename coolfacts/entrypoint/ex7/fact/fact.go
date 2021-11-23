package fact

type Fact struct {
	Image       string
	Description string
}

type Repo struct {
	Facts []Fact
}

func (s *Repo) Add(f Fact) {
	s.Facts = append(s.Facts, f)
}

func (s *Repo) GetAll() []Fact {
	return s.Facts
}
