package models

type DestinationSeriesType struct {
	SerieUuid string                 `json:"SerieUuid" bson:"SerieUuid"`
	Id        int64                  `json:"Id" bson:"Id"`
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	Tags      map[string]interface{} `json:"Tags" bson:"Tags"`
}

func NewDestinationSeriesRawType(SerieUuid string, Id int64, StudyUuid string, Tags map[string]interface{}) DestinationSeriesType {
	return DestinationSeriesType{
		StudyUuid: StudyUuid,
		SerieUuid: SerieUuid,
		Id:        Id,
		Tags:      Tags,
	}
}
