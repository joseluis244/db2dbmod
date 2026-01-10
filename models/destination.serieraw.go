package models

var DestinationSeriesRawCollection string = "SeriesRaw"

type DestinationSeriesRawType struct {
	StudyUuid string                 `json:"StudyUuid"`
	SerieUuid string                 `json:"SerieUuid"`
	Tags      map[string]interface{} `json:"Tags"`
	CreatedAt int64                  `json:"CreatedAt"`
	UpdatedAt int64                  `json:"UpdatedAt"`
}
