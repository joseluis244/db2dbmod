package models

var DestinationStudyCollection string = "StudyRaw"

type DestinationStudyType struct {
	StudyUuid string                 `json:"StudyUuid"`
	Tags      map[string]interface{} `json:"Tags"`
	CreatedAt int64                  `json:"CreatedAt"`
	UpdatedAt int64                  `json:"UpdatedAt"`
	BuildTime int64                  `json:"BuildTime"`
}
