package models

type Season struct {
	Year  int    `json:"season,string"`
	Races []Race `json:"Races"`
}
