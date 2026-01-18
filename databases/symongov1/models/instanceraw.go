package models

type SyMongoV1InstanceType struct {
	Uuid      string                 `json:"Uuid" bson:"Uuid"`
	Ae        string                 `json:"Ae" bson:"Ae"`
	CloudSync int64                  `json:"CloudSync" bson:"CloudSync"`
	Hash      string                 `json:"Hash" bson:"Hash"`
	Id        int64                  `json:"Id" bson:"Id"`
	Path      string                 `json:"Path" bson:"Path"`
	SerieUuid string                 `json:"SerieUuid" bson:"SerieUuid"`
	Size      int64                  `json:"Size" bson:"Size"`
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	Update    int64                  `json:"Update" bson:"Update"`
	Tags      map[string]interface{} `json:"Tags" bson:"Tags"`
}

func NewSyMongoV1InstanceType(Uuid string, Ae string, CloudSync int64, Hash string, Id int64, Path string, SerieUuid string, Size int64, StudyUuid string, Update int64, Tags map[string]interface{}) SyMongoV1InstanceType {
	return SyMongoV1InstanceType{
		Uuid:      Uuid,
		Ae:        Ae,
		CloudSync: CloudSync,
		Hash:      Hash,
		Id:        Id,
		Path:      Path,
		SerieUuid: SerieUuid,
		Size:      Size,
		StudyUuid: StudyUuid,
		Update:    Update,
		Tags:      Tags,
	}
}
