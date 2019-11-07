package models

type CurrentSeasonResponse struct {
	MRData struct {
		Limit  int    `json:"limit,string"`
		Offset int    `json:"offset,string"`
		Total  int    `json:"total,string"`
		Data   Season `json:"RaceTable"`
	} `json:"MRData"`
}
