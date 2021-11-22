package facthttp

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"workshop/exercises/ex8/fact"
)

var newsTemplate = `<!DOCTYPE html>
<html>
  <head><style>/* copy coolfacts/styles.css for some color 🎨*/</style></head>
  <body>
  <h1>Facts List</h1>
  <div>
    {{ range . }}
       <article>
            <h3>{{.Description}}</h3>
            <img src="{{.Image}}" width="50%" />
       </article>
    {{ end }}
  <div>
  </body>
</html>`

type FactRepository interface {
	Add(f fact.Fact)
	GetAll() []fact.Fact
}

type factsHandler struct {
	factRepo FactRepository
}

func NewFactsHandler(factRepo FactRepository) *factsHandler {
	return &factsHandler{
		factRepo: factRepo,
	}
}
func (h *factsHandler) Ping(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "PONG")
	if err != nil {
		errMessage := fmt.Sprintf("error writing response: %v", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
	}
}

func (h *factsHandler) Facts(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.New("facts").Parse(newsTemplate)
		if err != nil {
			log.Fatal(err)
		}
		facts := h.factRepo.GetAll()
		err = tmpl.Execute(w, facts)
		if err != nil {
			errMessage := fmt.Sprintf("error executing html: %v", err)
			http.Error(w, errMessage, http.StatusInternalServerError)

		}
	}
}
