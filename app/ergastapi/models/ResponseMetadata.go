package models

//ResponseMetadata contains common fields present in any ergast API response.
type ResponseMetadata struct {
	MRData struct {
		URL       string `json:"url"`
		Limit     string `json:"limit"`
		Offset    string `json:"offset"`
		Total     int    `json:"total,string"`
	} `json:"MRData"`
}
