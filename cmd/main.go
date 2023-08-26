package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"strings"

	"net/http"

	"github.com/iagojsilva/imersao-desafio2/internal/routes/entity"
	"github.com/iagojsilva/imersao-desafio2/internal/routes/infra/repository"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	count := 1

	db, sqlCoonectError := sql.Open("mysql", "root:root@tcp(localhost:3306)/routes")

	if sqlCoonectError != nil {
		panic(sqlCoonectError)
	}

	routeRepository := repository.NewRouteRepository(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/routes", func(w http.ResponseWriter, r *http.Request) {
		routesStruct, err := routeRepository.FindAll()

		if err != nil {
			http.Error(w, "Could not fetch routes", http.StatusBadRequest)
			return
		}

		routesJSON, _ := json.Marshal(routesStruct)
		w.Write([]byte(routesJSON))
	})

	r.Post("/routes", func(w http.ResponseWriter, r *http.Request) {
		var rawRoute *entity.RawRoute
		body, _ := io.ReadAll(r.Body)

		err := json.NewDecoder(strings.NewReader(string(body))).Decode(&rawRoute)

		if err != nil {
			http.Error(w, "Invalid JSON data", http.StatusBadRequest)
			return
		}

		route, err := routeRepository.Create(rawRoute, count)

		if err != nil {
			http.Error(w, "Could not create route", http.StatusInternalServerError)
			return
		}

		count++

		routeJSON, _ := json.Marshal(route)

		w.Header().Set("Content-Type", "application/json")
		w.Write(routeJSON)
	})
	println("Listening at 8080")

	http.ListenAndServe(":8080", r)
}
