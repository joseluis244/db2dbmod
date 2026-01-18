package models

type DestinationStudyType struct {
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	Id        int64                  `json:"Id" bson:"Id"`
	Tags      map[string]interface{} `json:"Tags" bson:"Tags"`
}

func NewDestinationStudyType(StudyUuid string, Id int64, Tags map[string]interface{}) DestinationStudyType {
	return DestinationStudyType{
		StudyUuid: StudyUuid,
		Id:        Id,
		Tags:      Tags,
	}
}
