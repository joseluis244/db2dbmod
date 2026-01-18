package models

type OrtMySQLv1StudyType struct {
	Id        int64                  `json:"Id" bson:"Id"`
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	Tags      map[string]interface{} `json:"Tags" bson:"Tags"`
}

type OrtStudyRaw struct {
	Id         int64  `json:"Id" bson:"Id"`
	StudyUuid  string `json:"StudyUuid" bson:"StudyUuid"`
	TagGroup   int    `json:"TagGroup" bson:"TagGroup"`
	TagElement int    `json:"TagElement" bson:"TagElement"`
	Value      string `json:"Value" bson:"Value"`
}

func NewOrtMySQLv1StudyType(Id int64, StudyUuid string, Tags map[string]interface{}) OrtMySQLv1StudyType {
	return OrtMySQLv1StudyType{
		Id:        Id,
		StudyUuid: StudyUuid,
		Tags:      Tags,
	}
}

func NewOrtStudyRaw(Id int64, StudyUuid string, TagGroup int, TagElement int, Value string) OrtStudyRaw {
	return OrtStudyRaw{
		Id:         Id,
		StudyUuid:  StudyUuid,
		TagGroup:   TagGroup,
		TagElement: TagElement,
		Value:      Value,
	}
}
