package models

//Race is a single race event happened on particular Circuit at particular Date and Time.
type Race struct {
	Year     int
	Round    int
	URL      string
	RaceName string
	Circuit  Circuit
	Date     string
	Time     string
}
