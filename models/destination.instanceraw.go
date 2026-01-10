package models

var DestinationInstanceRawCollection string = "InstanceRaw"

type DestinationInstanceRawType struct {
	Uuid      string                 `json:"Uuid"`
	Ae        string                 `json:"Ae"`
	Tags      map[string]interface{} `json:"Tags"`
	CreatedAt int64                  `json:"CreatedAt"`
	UpdatedAt int64                  `json:"UpdatedAt"`
	StudyUuid string                 `json:"StudyUuid"`
	SerieUuid string                 `json:"SerieUuid"`
	Hash      string                 `json:"Hash"`
	Size      int64                  `json:"Size"`
	Path      string                 `json:"Path"`
	Store     string                 `json:"Store"`
}
