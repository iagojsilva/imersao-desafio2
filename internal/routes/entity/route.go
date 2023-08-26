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

type RawRoute struct {
	Name        string  `json:"name"`
	Source      *Coords `json:"source"`
	Destination *Coords `json:"destination"`
}

func NewRawRoute(name string, source *Coords, destination *Coords) *RawRoute {
	return &RawRoute{
		Name:        name,
		Source:      source,
		Destination: destination,
	}
}

type Route struct {
	ID          int
	Name        string  `json:"name"`
	Source      *Coords `json:"source"`
	Destination *Coords `json:"destination"`
}

func NewRoute(id int, name string, source *Coords, destination *Coords) *Route {
	return &Route{
		ID:          id,
		Name:        name,
		Source:      source,
		Destination: destination,
	}
}

type RouteRepository interface {
	FindAll() ([]*Route, error)
	CreateRoute(route *RawRoute) error
}
