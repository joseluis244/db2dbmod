package models

type SourceMySQLv1InstanceType struct {
	StudyUuid    string                 `json:"StudyUuid"`
	SerieUuid    string                 `json:"SerieUuid"`
	InstanceUuid string                 `json:"InstanceUuid"`
	Hash         string                 `json:"Hash"`
	Size         int64                  `json:"Size"`
	AE           string                 `json:"AE"`
	Tags         map[string]interface{} `json:"Tags"`
}
