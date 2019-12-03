package models

//Location is the combination of Coordinates and the Country information.
type Location struct {
	Lat      float64
	Long     float64
	Locality string
	Country  string
}
