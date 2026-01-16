package models

var DestinationV3Collection = "V3"

type DestinationV3InstanceType struct {
	Uuid  string                 `json:"Uuid"`
	Ae    string                 `json:"Ae"`
	Hash  string                 `json:"Hash"`
	Size  int64                  `json:"Size"`
	Path  string                 `json:"Path"`
	Store string                 `json:"Store"`
	Tags  map[string]interface{} `json:"Tags"`
}

type DestinationV3SeriesType struct {
	SerieUuid string                      `json:"SerieUuid"`
	Tags      map[string]interface{}      `json:"Tags"`
	Instances []DestinationV3InstanceType `json:"Instances"`
}

type DestinationV3Type struct {
	DealerID  string                    `json:"DealerID"`
	ClientID  string                    `json:"ClientID"`
	BranchID  string                    `json:"BranchID"`
	StudyUuid string                    `json:"StudyUuid"`
	Tags      map[string]interface{}    `json:"Tags"`
	Series    []DestinationV3SeriesType `json:"Series"`
	CreatedAt int64                     `json:"CreatedAt"`
	UpdatedAt int64                     `json:"UpdatedAt"`
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
	}
}
