package mysqlv12mogo

import (
	"github.com/joseluis244/db2dbmod/builder/mysqlv12mogo/instance"
	"github.com/joseluis244/db2dbmod/builder/mysqlv12mogo/serie"
	"github.com/joseluis244/db2dbmod/builder/mysqlv12mogo/study"
	mysqlv1instance "github.com/joseluis244/db2dbmod/dbsource/mysqlv1/instance"
	mysqlv1serie "github.com/joseluis244/db2dbmod/dbsource/mysqlv1/serie"
	mysqlv1study "github.com/joseluis244/db2dbmod/dbsource/mysqlv1/study"
)

type ChanelChangesToBuild struct {
	Studies    []mysqlv1study.SourceMySQLv1StudyType
	Series     []mysqlv1serie.SourceMySQLv1SerieType
	Instances  []mysqlv1instance.SourceMySQLv1InstanceType
	LastChange int64
}

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
