package models

type SyMongoV1SystemType struct {
	VERSION string `json:"VERSION" bson:"VERSION"`
	LASTSEQ int64  `json:"LASTSEQ" bson:"LASTSEQ"`
	ID      int64  `json:"ID" bson:"ID"`
}

func NewSyMongoV1SystemType(VERSION string, LASTSEQ int64, ID int64) SyMongoV1SystemType {
	return SyMongoV1SystemType{
		VERSION: VERSION,
		LASTSEQ: LASTSEQ,
		ID:      ID,
	}
}
