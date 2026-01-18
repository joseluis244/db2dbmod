package models

type SyMongoV1SeriesType struct {
	SerieUuid string                 `json:"SerieUuid" bson:"SerieUuid"`
	Id        int64                  `json:"Id" bson:"Id"`
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	Tags      map[string]interface{} `json:"Tags" bson:"Tags"`
}

func NewSyMongoV1SeriesType(SerieUuid string, Id int64, StudyUuid string, Tags map[string]interface{}) SyMongoV1SeriesType {
	return SyMongoV1SeriesType{
		StudyUuid: StudyUuid,
		SerieUuid: SerieUuid,
		Id:        Id,
		Tags:      Tags,
	}
}
