package models

type DestinationSeriesType struct {
	DealerID  string                 `json:"DealerID" bson:"DealerID"`
	ClientID  string                 `json:"ClientID" bson:"ClientID"`
	BranchID  string                 `json:"BranchID" bson:"BranchID"`
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	SerieUuid string                 `json:"SerieUuid" bson:"SerieUuid"`
	Tags      map[string]interface{} `json:"Tags" bson:"Tags"`
	CreatedAt int64                  `json:"CreatedAt" bson:"CreatedAt"`
	UpdatedAt int64                  `json:"UpdatedAt" bson:"UpdatedAt"`
	Sync      SyncType               `json:"Sync" bson:"Sync"`
}

func NewDestinationSeriesRawType(DealerID string, ClientID string, BranchID string, StudyUuid string, SerieUuid string, Tags map[string]interface{}) DestinationSeriesType {
	return DestinationSeriesType{
		DealerID:  DealerID,
		ClientID:  ClientID,
		BranchID:  BranchID,
		StudyUuid: StudyUuid,
		SerieUuid: SerieUuid,
		Tags:      Tags,
		CreatedAt: 0,
		UpdatedAt: 0,
		Sync: SyncType{
			Status:   "pending",
			SyncTime: 0,
		},
	}
}
