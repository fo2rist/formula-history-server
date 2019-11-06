package models

type CurrentSeasonResponse struct {
	MRData struct {
		Total int    `json:"total,string"`
		Data  Season `json:"RaceTable"`
	} `json:"MRData"`
}

type Season struct {
	Season string `json:"season"`
	Races  []struct {
		Season   string `json:"season"`
		Round    string `json:"round"`
		URL      string `json:"url"`
		RaceName string `json:"raceName"`
		Circuit  struct {
			CircuitID   string `json:"circuitId"`
			URL         string `json:"url"`
			CircuitName string `json:"circuitName"`
			Location    struct {
				Lat      string `json:"lat"`
				Long     string `json:"long"`
				Locality string `json:"locality"`
				Country  string `json:"country"`
			} `json:"Location"`
		} `json:"Circuit"`
		Date string `json:"date"`
		Time string `json:"time"`
	} `json:"Races"`
}
