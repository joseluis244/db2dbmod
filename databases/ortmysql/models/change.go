package models

type SourceMySQLv1ChangesType struct {
	Seq          int64  `json:"seq"`
	ChangeType   int    `json:"changeType"`
	InternalId   int64  `json:"internalId"`
	ResourceType string `json:"resourceType"`
	Date         string `json:"date"`
}
