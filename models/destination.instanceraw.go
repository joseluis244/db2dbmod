package models

var DestinationInstanceRawCollection string = "InstanceRaw"

type DestinationInstanceType struct {
	DealerID     string                 `json:"DealerID" bson:"DealerID"`
	ClientID     string                 `json:"ClientID" bson:"ClientID"`
	BranchID     string                 `json:"BranchID" bson:"BranchID"`
	StudyUuid    string                 `json:"StudyUuid" bson:"StudyUuid"`
	SerieUuid    string                 `json:"SerieUuid" bson:"SerieUuid"`
	InstanceUuid string                 `json:"InstanceUuid" bson:"InstanceUuid"`
	Tags         map[string]interface{} `json:"Tags" bson:"Tags"`
	Hash         string                 `json:"Hash" bson:"Hash"`
	Size         int64                  `json:"Size" bson:"Size"`
	Path         string                 `json:"Path" bson:"Path"`
	Store        string                 `json:"Store" bson:"Store"` // "local", "s3", "r2"
	Ae           string                 `json:"Ae" bson:"Ae"`
	Sync         SyncType               `json:"Sync" bson:"Sync"`
	CreatedAt    int64                  `json:"CreatedAt" bson:"CreatedAt"`
	UpdatedAt    int64                  `json:"UpdatedAt" bson:"UpdatedAt"`
}

func NewDestinationInstanceRawType(DealerID string, ClientID string, BranchID string, InstanceUuid string, Ae string, Tags map[string]interface{}, StudyUuid string, SerieUuid string, Hash string, Size int64, Path string, Store string) DestinationInstanceType {
	return DestinationInstanceType{
		DealerID:     DealerID,
		ClientID:     ClientID,
		BranchID:     BranchID,
		InstanceUuid: InstanceUuid,
		Ae:           Ae,
		Tags:         Tags,
		CreatedAt:    0,
		UpdatedAt:    0,
		StudyUuid:    StudyUuid,
		SerieUuid:    SerieUuid,
		Hash:         Hash,
		Size:         Size,
		Path:         Path,
		Store:        Store,
		Sync: SyncType{
			Status:   "pending",
			SyncTime: 0,
		},
	}
}
