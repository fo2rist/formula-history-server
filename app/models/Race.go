package models

type Race struct {
	Year     int
	Round    int
	URL      string
	RaceName string
	Circuit  Circuit
	Date     string
	Time     string
}

