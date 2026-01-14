package models

var DestinationStudyCollection string = "StudyRaw"

type SyncType struct {
	Status   string `json:"Status"` // "pending", "syncing", "synced"
	SyncTime int64  `json:"SyncTime"`
}

type DestinationStudyType struct {
	DealerID  string                 `json:"DealerID"`
	ClientID  string                 `json:"ClientID"`
	BranchID  string                 `json:"BranchID"`
	StudyUuid string                 `json:"StudyUuid"`
	Tags      map[string]interface{} `json:"Tags"`
	CreatedAt int64                  `json:"CreatedAt"`
	UpdatedAt int64                  `json:"UpdatedAt"`
	BuildTime int64                  `json:"BuildTime"`
	Sync      SyncType               `json:"Sync"`
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
