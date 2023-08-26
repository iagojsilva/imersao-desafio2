package repository

import (
	"database/sql"
	"encoding/json"

	"github.com/iagojsilva/imersao-desafio2/internal/routes/entity"
)

type RouteRepositoryMysql struct {
	db *sql.DB
}

func NewRouteRepository(db *sql.DB) *RouteRepositoryMysql {
	return &RouteRepositoryMysql{
		db: db,
	}
}

func (r *RouteRepositoryMysql) Create(route *entity.Route) error {

	sourceJSON, sourceJsonErr := json.Marshal(route.Source)
	destinationJSON, destJsonErr := json.Marshal(route.Destination)

	if sourceJsonErr != nil || destJsonErr != nil {
		panic("sourceJsonErr, destJsonErr has an error")
	}

	query := "INSERT INTO routes (id, name, source, destination) VALUES (?, ?, ?, ?);"
	_, err := r.db.Exec(query, route.ID, route.Name, sourceJSON, destinationJSON)

	if err != nil {
		return err
	}
	return nil
}

func (r *RouteRepositoryMysql) FindAll() (*[]entity.Route, error) {
	query := "SELECT id, name, source, destination FROM routes;"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var routes []entity.Route

	for rows.Next() {
		var route entity.Route
		var sourceJSON []byte      // To hold the raw JSON data
		var destinationJSON []byte // To hold the raw JSON data

		err := rows.Scan(&route.ID, &route.Name, &sourceJSON, &destinationJSON)
		if err != nil {
			return nil, err
		}

		// Unmarshal the JSON data into the Coords struct
		err = json.Unmarshal(sourceJSON, &route.Source)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(destinationJSON, &route.Destination)

		if err != nil {
			return nil, err
		}

		routes = append(routes, route)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &routes, nil
}
