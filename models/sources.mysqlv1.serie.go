package models

type SourceMySQLv1SerieType struct {
	StudyUuid string                 `json:"StudyUuid"`
	SerieUuid string                 `json:"SerieUuid"`
	Tags      map[string]interface{} `json:"Tags"`
}
