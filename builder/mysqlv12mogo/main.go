package mysqlv12mogo

import (
	"github.com/joseluis244/db2dbmod/builder/mysqlv12mogo/instance"
	"github.com/joseluis244/db2dbmod/builder/mysqlv12mogo/serie"
	"github.com/joseluis244/db2dbmod/builder/mysqlv12mogo/study"
)

type BuilderStruct struct {
	DealerID string
	ClientID string
	BranchID string
	Instance *instance.InstanceStruct
	Study    *study.StudyStruct
	Series   *serie.SerieStruct
}

func New(DealerID string, ClientID string, BranchID string) *BuilderStruct {
	return &BuilderStruct{
		DealerID: DealerID,
		ClientID: ClientID,
		BranchID: BranchID,
		Instance: instance.New(DealerID, ClientID, BranchID),
		Study:    study.New(DealerID, ClientID, BranchID),
		Series:   serie.New(DealerID, ClientID, BranchID),
	}
}
