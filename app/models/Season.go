package models

//Season is the calendar of Races for particular Year.
type Season struct {
	Year  int
	Races []Race
}
