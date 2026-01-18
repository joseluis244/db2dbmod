package models

type DestinationV3InstanceType struct {
	Uuid      string                 `json:"Uuid" bson:"Uuid"`
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	SerieUuid string                 `json:"SerieUuid" bson:"SerieUuid"`
	Size      int64                  `json:"Size" bson:"Size"`
	Hash      string                 `json:"Hash" bson:"Hash"`
	Path      string                 `json:"Path" bson:"Path"`
	Id        int64                  `json:"Id" bson:"Id"`
	Tags      map[string]interface{} `json:"tags" bson:"tags"`
	Ae        string                 `json:"Ae" bson:"Ae"`
}

type DestinationV3SeriesType struct {
	StudyUuid string                      `json:"StudyUuid" bson:"StudyUuid"`
	SerieUuid string                      `json:"SerieUuid" bson:"SerieUuid"`
	Id        int64                       `json:"Id" bson:"Id"`
	Tags      map[string]interface{}      `json:"tags" bson:"tags"`
	Instances []DestinationV3InstanceType `json:"instances" bson:"instances"`
}

type DestinationV3Type struct {
	StudyUuid  string                    `json:"StudyUuid" bson:"StudyUuid"`
	Complete   bool                      `json:"Complete" bson:"Complete"`
	Id         int64                     `json:"Id" bson:"Id"`
	LastSync   int64                     `json:"LastSync" bson:"LastSync"`
	LastUpdate int64                     `json:"LastUpdate" bson:"LastUpdate"`
	Series     []DestinationV3SeriesType `json:"series" bson:"series"`
	Tags       map[string]interface{}    `json:"tags" bson:"tags"`
}

func NewDestinationV3InstanceType(Uuid string, StudyUuid string, SerieUuid string, Size int64, Hash string, Path string, Id int64, Tags map[string]interface{}, Ae string) DestinationV3InstanceType {
	return DestinationV3InstanceType{
		Uuid:      Uuid,
		StudyUuid: StudyUuid,
		SerieUuid: SerieUuid,
		Size:      Size,
		Hash:      Hash,
		Path:      Path,
		Id:        Id,
		Tags:      Tags,
		Ae:        Ae,
	}
}

func NewDestinationV3SeriesType(StudyUuid string, SerieUuid string, Id int64, Tags map[string]interface{}, Instances []DestinationV3InstanceType) DestinationV3SeriesType {
	return DestinationV3SeriesType{
		StudyUuid: StudyUuid,
		SerieUuid: SerieUuid,
		Id:        Id,
		Tags:      Tags,
		Instances: Instances,
	}
}

func NewDestinationV3Type(StudyUuid string, Complete bool, Id int64, LastSync int64, LastUpdate int64, Series []DestinationV3SeriesType, Tags map[string]interface{}) DestinationV3Type {
	return DestinationV3Type{
		StudyUuid:  StudyUuid,
		Complete:   Complete,
		Id:         Id,
		LastSync:   LastSync,
		LastUpdate: LastUpdate,
		Series:     Series,
		Tags:       Tags,
	}
}
