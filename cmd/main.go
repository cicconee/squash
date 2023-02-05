package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/cicconee/squash/team"
	"github.com/go-chi/chi/v5"
)

var service = flag.String("s", "", "specify the microservice to run")

func main() {
	teamHandler := team.NewHandler()
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]interface{}{
			"message": "success",
		}

		data, _ := json.Marshal(payload)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	})
	r.Get("/teams", teamHandler.Create)
	http.ListenAndServe(":3000", r)
}
