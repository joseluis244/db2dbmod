package models

type SyMongoV1StudyType struct {
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	Id        int64                  `json:"Id" bson:"Id"`
	Tags      map[string]interface{} `json:"Tags" bson:"Tags"`
}

func NewSyMongoV1StudyType(StudyUuid string, Id int64, Tags map[string]interface{}) SyMongoV1StudyType {
	return SyMongoV1StudyType{
		StudyUuid: StudyUuid,
		Id:        Id,
		Tags:      Tags,
	}
}
