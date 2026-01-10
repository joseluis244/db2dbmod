package models

type SourceMySQLv1StudyType struct {
	StudyUuid string                 `json:"StudyUuid"`
	Tags      map[string]interface{} `json:"Tags"`
	CreatedAt int64                  `json:"CreatedAt"`
	UpdatedAt int64                  `json:"UpdatedAt"`
	BuildTime int64                  `json:"BuildTime"`
}
