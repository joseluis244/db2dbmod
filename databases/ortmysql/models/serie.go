package models

type OrtMySQLv1SerieType struct {
	Id        int64                  `json:"Id" bson:"Id"`
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	SerieUuid string                 `json:"SerieUuid" bson:"SerieUuid"`
	Tags      map[string]interface{} `json:"tags" bson:"tags"`
}

type OrtSerieRaw struct {
	Id         int64  `json:"Id" bson:"Id"`
	StudyUuid  string `json:"StudyUuid" bson:"StudyUuid"`
	SerieUuid  string `json:"SerieUuid" bson:"SerieUuid"`
	TagGroup   int    `json:"TagGroup" bson:"TagGroup"`
	TagElement int    `json:"TagElement" bson:"TagElement"`
	Value      string `json:"Value" bson:"Value"`
}

func NewOrtMySQLv1SerieType(Id int64, StudyUuid string, SerieUuid string, Tags map[string]interface{}) OrtMySQLv1SerieType {
	return OrtMySQLv1SerieType{
		Id:        Id,
		StudyUuid: StudyUuid,
		SerieUuid: SerieUuid,
		Tags:      Tags,
	}
}

func NewOrtSerieRaw(Id int64, StudyUuid string, SerieUuid string, TagGroup int, TagElement int, Value string) OrtSerieRaw {
	return OrtSerieRaw{
		Id:         Id,
		StudyUuid:  StudyUuid,
		SerieUuid:  SerieUuid,
		TagGroup:   TagGroup,
		TagElement: TagElement,
		Value:      Value,
	}
}
