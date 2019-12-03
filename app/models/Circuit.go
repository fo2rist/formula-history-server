package models

//Circuit is the race track at particular Location.
type Circuit struct {
	CircuitID   string
	URL         string
	CircuitName string
	Location    Location
}

