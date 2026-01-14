package models

var DestinationInstanceRawCollection string = "InstanceRaw"

type DestinationInstanceType struct {
	DealerID  string                 `json:"DealerID"`
	ClientID  string                 `json:"ClientID"`
	BranchID  string                 `json:"BranchID"`
	Uuid      string                 `json:"Uuid"`
	Ae        string                 `json:"Ae"`
	Tags      map[string]interface{} `json:"Tags"`
	CreatedAt int64                  `json:"CreatedAt"`
	UpdatedAt int64                  `json:"UpdatedAt"`
	StudyUuid string                 `json:"StudyUuid"`
	SerieUuid string                 `json:"SerieUuid"`
	Hash      string                 `json:"Hash"`
	Size      int64                  `json:"Size"`
	Path      string                 `json:"Path"`
	Store     string                 `json:"Store"` // "local", "s3", "r2"
	Sync      SyncType               `json:"Sync"`
}

func NewDestinationInstanceRawType(DealerID string, ClientID string, BranchID string, Uuid string, Ae string, Tags map[string]interface{}, StudyUuid string, SerieUuid string, Hash string, Size int64, Path string, Store string) DestinationInstanceType {
	return DestinationInstanceType{
		DealerID:  DealerID,
		ClientID:  ClientID,
		BranchID:  BranchID,
		Uuid:      Uuid,
		Ae:        Ae,
		Tags:      Tags,
		CreatedAt: 0,
		UpdatedAt: 0,
		StudyUuid: StudyUuid,
		SerieUuid: SerieUuid,
		Hash:      Hash,
		Size:      Size,
		Path:      Path,
		Store:     Store,
		Sync: SyncType{
			Status:   "pending",
			SyncTime: 0,
		},
	}
}
