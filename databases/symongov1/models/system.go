package models

type SystemType struct {
	VERSION string `json:"VERSION" bson:"VERSION"`
	LASTSEQ int64  `json:"LASTSEQ" bson:"LASTSEQ"`
	ID      int64  `json:"ID" bson:"ID"`
}

func NewSystemType(VERSION string, LASTSEQ int64, ID int64) SystemType {
	return SystemType{
		VERSION: VERSION,
		LASTSEQ: LASTSEQ,
		ID:      ID,
	}
}
