package main

import (
	"log"
	"net/http"

	"github.com/nkcraddock/mervis"
	"github.com/nkcraddock/mervis/res/client"
)

func main() {
	opts := getopts()

	var clientLocator client.ResourceLocator

	if opts.ClientRoot != "" {
		clientLocator = &client.FsLocator{opts.ClientRoot}
	} else {
		clientLocator = &client.BinDataLocator{}
	}

	mervis := mervis.New(clientLocator)

	server := &http.Server{
		Addr:    opts.Addr,
		Handler: logMiddleware(mervis),
	}

	log.Fatal(server.ListenAndServe())
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("REQUEST", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
