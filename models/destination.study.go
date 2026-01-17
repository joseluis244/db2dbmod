package models

var DestinationStudyCollection string = "StudyRaw"

type DestinationStudyType struct {
	DealerID  string                 `json:"DealerID" bson:"DealerID"`
	ClientID  string                 `json:"ClientID" bson:"ClientID"`
	BranchID  string                 `json:"BranchID" bson:"BranchID"`
	StudyUuid string                 `json:"StudyUuid" bson:"StudyUuid"`
	Tags      map[string]interface{} `json:"Tags" bson:"Tags"`
	CreatedAt int64                  `json:"CreatedAt" bson:"CreatedAt"`
	UpdatedAt int64                  `json:"UpdatedAt" bson:"UpdatedAt"`
	BuildTime int64                  `json:"BuildTime" bson:"BuildTime"`
	Sync      SyncType               `json:"Sync" bson:"Sync"`
}

func NewDestinationStudyType(DealerID string, ClientID string, BranchID string, StudyUuid string, Tags map[string]interface{}) DestinationStudyType {
	return DestinationStudyType{
		DealerID:  DealerID,
		ClientID:  ClientID,
		BranchID:  BranchID,
		StudyUuid: StudyUuid,
		Tags:      Tags,
		CreatedAt: 0,
		UpdatedAt: 0,
		BuildTime: 0,
		Sync: SyncType{
			Status:   "pending",
			SyncTime: 0,
		},
	}
}
