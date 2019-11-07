package models

type Circuit struct {
	CircuitID   string `json:"circuitId"`
	URL         string `json:"url"`
	CircuitName string `json:"circuitName"`
	Location    struct {
		Lat      float64 `json:"lat,string"`
		Long     float64 `json:"long,string"`
		Locality string  `json:"locality"`
		Country  string  `json:"country"`
	} `json:"Location"`
}
