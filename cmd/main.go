package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/iagojsilva/imersao-desafio2/internal/routes/entity"
	"github.com/iagojsilva/imersao-desafio2/internal/routes/infra/repository"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, sqlCoonectError := sql.Open("mysql", "root:root@tcp(localhost:3306)/routes")

	if sqlCoonectError != nil {
		panic(sqlCoonectError)
	}

	routeRepository := repository.NewRouteRepository(db)

	source := entity.NewCoords(46, 57)

	destination := entity.NewCoords(42, 57)

	route := entity.NewRoute(uuid.NewString(), "my route", source, destination)

	err := routeRepository.Create(route)

	if err != nil {
		panic(err)
	}
	fmt.Println("Route created", route)

	routes, findAllError := routeRepository.FindAll()

	if findAllError != nil {
		panic(findAllError)
	}
	fmt.Println(routes)
}
