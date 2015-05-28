package mervis

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nkcraddock/mervis/res/client"
)

type Mervis struct {
	handler http.Handler
}

type Handler interface {
	Handle(r *mux.Route)
}

func New(clientResource client.ResourceLocator) *Mervis {
	m := mux.NewRouter()

	c := client.NewHandler(clientResource)

	c.Handle(m.NewRoute().Subrouter())

	return &Mervis{m}
}

func (m *Mervis) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.handler.ServeHTTP(w, r)
}
