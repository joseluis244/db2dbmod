package models

type DestinationV3InstanceType struct {
	Uuid  string                 `json:"Uuid" bson:"Uuid"`
	Ae    string                 `json:"Ae" bson:"Ae"`
	Hash  string                 `json:"Hash" bson:"Hash"`
	Size  int64                  `json:"Size" bson:"Size"`
	Path  string                 `json:"Path" bson:"Path"`
	Store string                 `json:"Store" bson:"Store"`
	Tags  map[string]interface{} `json:"Tags" bson:"Tags"`
}

type DestinationV3SeriesType struct {
	SerieUuid string                      `json:"SerieUuid" bson:"SerieUuid"`
	Tags      map[string]interface{}      `json:"Tags" bson:"Tags"`
	Instances []DestinationV3InstanceType `json:"Instances" bson:"Instances"`
}

type DestinationV3Type struct {
	DealerID  string                    `json:"DealerID" bson:"DealerID"`
	ClientID  string                    `json:"ClientID" bson:"ClientID"`
	BranchID  string                    `json:"BranchID" bson:"BranchID"`
	StudyUuid string                    `json:"StudyUuid" bson:"StudyUuid"`
	Tags      map[string]interface{}    `json:"Tags" bson:"Tags"`
	Series    []DestinationV3SeriesType `json:"Series" bson:"Series"`
	CreatedAt int64                     `json:"CreatedAt" bson:"CreatedAt"`
	UpdatedAt int64                     `json:"UpdatedAt" bson:"UpdatedAt"`
	Sync      SyncType                  `json:"Sync" bson:"Sync"`
}

func NewDestinationV3InstanceType(Uuid string, Ae string, Hash string, Size int64, Path string, Store string, Tags map[string]interface{}) DestinationV3InstanceType {
	return DestinationV3InstanceType{
		Uuid:  Uuid,
		Ae:    Ae,
		Hash:  Hash,
		Size:  Size,
		Path:  Path,
		Store: Store,
		Tags:  Tags,
	}
}

func NewDestinationV3SeriesType(SerieUuid string, Tags map[string]interface{}, Instances []DestinationV3InstanceType) DestinationV3SeriesType {
	return DestinationV3SeriesType{
		SerieUuid: SerieUuid,
		Tags:      Tags,
		Instances: Instances,
	}
}

func NewDestinationV3Type(DealerID string, ClientID string, BranchID string, StudyUuid string, Tags map[string]interface{}, CreatedAt int64, UpdatedAt int64, Series []DestinationV3SeriesType) DestinationV3Type {
	return DestinationV3Type{
		DealerID:  DealerID,
		ClientID:  ClientID,
		BranchID:  BranchID,
		StudyUuid: StudyUuid,
		Tags:      Tags,
		Series:    Series,
		CreatedAt: CreatedAt,
		UpdatedAt: UpdatedAt,
		Sync: SyncType{
			Status:   "pending",
			SyncTime: 0,
		},
	}
}
