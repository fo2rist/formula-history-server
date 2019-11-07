package models

//ResponseMetadata contains common fields present in any ergast API response.
type ResponseMetadata struct {
	MRData struct {
		URL       string `json:"url"`
		Limit     int `json:"limit,string"`
		Offset    int `json:"offset,string"`
		Total     int `json:"total,string"`
	} `json:"MRData"`
}
