package models

type OrtMySQLv1InstanceType struct {
	Id           int64                  `json:"Id"`
	StudyUuid    string                 `json:"StudyUuid"`
	SerieUuid    string                 `json:"SerieUuid"`
	InstanceUuid string                 `json:"InstanceUuid"`
	Hash         string                 `json:"Hash"`
	Size         int64                  `json:"Size"`
	AE           string                 `json:"AE"`
	Tags         map[string]interface{} `json:"Tags"`
}

type OrtInstanceRaw struct {
	Id           int64  `json:"Id"`
	StudyUuid    string `json:"StudyUuid"`
	SerieUuid    string `json:"SerieUuid"`
	InstanceUuid string `json:"InstanceUuid"`
	Hash         string `json:"Hash"`
	Size         int64  `json:"Size"`
	AE           string `json:"AE"`
	TagGroup     int    `json:"TagGroup"`
	TagElement   int    `json:"TagElement"`
	Value        string `json:"Value"`
}

func NewOrtMySQLv1InstanceType(id int64, studyUuid string, serieUuid string, instanceUuid string, hash string, size int64, ae string, tags map[string]interface{}) *OrtMySQLv1InstanceType {
	return &OrtMySQLv1InstanceType{
		Id:           id,
		StudyUuid:    studyUuid,
		SerieUuid:    serieUuid,
		InstanceUuid: instanceUuid,
		Hash:         hash,
		Size:         size,
		AE:           ae,
		Tags:         tags,
	}
}

func NewOrtInstanceRaw(id int64, studyUuid string, serieUuid string, instanceUuid string, hash string, size int64, ae string, tagGroup int, tagElement int, value string) *OrtInstanceRaw {
	return &OrtInstanceRaw{
		Id:           id,
		StudyUuid:    studyUuid,
		SerieUuid:    serieUuid,
		InstanceUuid: instanceUuid,
		Hash:         hash,
		Size:         size,
		AE:           ae,
		TagGroup:     tagGroup,
		TagElement:   tagElement,
		Value:        value,
	}
}
