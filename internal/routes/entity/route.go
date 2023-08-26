package entity

type Coords struct {
	Lat uint8 `json:"lat"`
	Lng uint8 `json:"lng"`
}

func NewCoords(lat uint8, lng uint8) *Coords {
	return &Coords{
		Lat: lat,
		Lng: lng,
	}
}

type Route struct {
	ID          string
	Name        string
	Source      *Coords
	Destination *Coords
}

func NewRoute(id string, name string, source *Coords, destination *Coords) *Route {
	return &Route{
		ID:          id,
		Name:        name,
		Source:      source,
		Destination: destination,
	}
}

type RouteRepository interface {
	FindAll() ([]*Route, error)
	CreateRoute(route *Route) error
}
