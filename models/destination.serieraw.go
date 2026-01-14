package models

var DestinationSeriesRawCollection string = "SeriesRaw"

type DestinationSeriesType struct {
	DealerID  string                 `json:"DealerID"`
	ClientID  string                 `json:"ClientID"`
	BranchID  string                 `json:"BranchID"`
	StudyUuid string                 `json:"StudyUuid"`
	SerieUuid string                 `json:"SerieUuid"`
	Tags      map[string]interface{} `json:"Tags"`
	CreatedAt int64                  `json:"CreatedAt"`
	UpdatedAt int64                  `json:"UpdatedAt"`
	Sync      SyncType               `json:"Sync"`
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
