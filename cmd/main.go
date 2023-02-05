package main

import (
	"flag"
	"net/http"

	"github.com/cicconee/squash/team"
	"github.com/go-chi/chi/v5"
)

var service = flag.String("s", "", "specify the microservice to run")

func main() {
	teamHandler := team.NewHandler()
	r := chi.NewRouter()
	r.Get("/teams", teamHandler.Create)
	http.ListenAndServe(":5000", r)
}
