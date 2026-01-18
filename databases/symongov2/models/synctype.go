package models

type SyncType struct {
	Status   string `json:"Status" bson:"Status"` // "pending", "syncing", "synced"
	SyncTime int64  `json:"SyncTime" bson:"SyncTime"`
}
