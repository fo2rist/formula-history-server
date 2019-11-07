package models

type Race struct {
	Year     int     `json:"season,string"`
	Round    int     `json:"round,string"`
	URL      string  `json:"url"`
	RaceName string  `json:"raceName"`
	Circuit  Circuit `json:"Circuit"`
	Date     string  `json:"date"`
	Time     string  `json:"time"`
}

