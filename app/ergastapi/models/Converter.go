package models

import (
	domain "github.com/fo2rist/formula-history/app/models"
)

type Vertex struct {
	X, Y float64
}

func (season *Season) ToDomainModel() *domain.Season {
	result := domain.Season{
		Year:  season.Year,
		Races: make([]domain.Race, len(season.Races)),
	}
	for i := range result.Races {
		sourceModel := season.Races[i]
		result.Races[i] = domain.Race{
			Year:     sourceModel.Year,
			Round:    sourceModel.Round,
			URL:      sourceModel.URL,
			RaceName: sourceModel.RaceName,
			Date:     sourceModel.Date,
			Time:     sourceModel.Time,
			Circuit:  domain.Circuit{
				CircuitID:   sourceModel.Circuit.CircuitID,
				URL:         sourceModel.Circuit.URL,
				CircuitName: sourceModel.Circuit.CircuitName,
				Location: domain.Location {
					Lat:      sourceModel.Circuit.Location.Lat,
					Long:     sourceModel.Circuit.Location.Long,
					Locality: sourceModel.Circuit.Location.Locality,
					Country:  sourceModel.Circuit.Location.Country,
				},
			},
		}
	}
	return &result
}
