package ortmysql2symongov2

import (
	"github.com/joseluis244/db2dbmod/builder/ortmysql2symongov2/instance"
	"github.com/joseluis244/db2dbmod/builder/ortmysql2symongov2/serie"
	"github.com/joseluis244/db2dbmod/builder/ortmysql2symongov2/study"
	v3 "github.com/joseluis244/db2dbmod/builder/ortmysql2symongov2/v3"
)

type BuilderStruct struct {
	DealerID string
	ClientID string
	BranchID string
	Instance *instance.InstanceStruct
	Study    *study.StudyStruct
	Series   *serie.SerieStruct
	V3       *v3.V3Struct
}

func New(DealerID string, ClientID string, BranchID string) *BuilderStruct {
	return &BuilderStruct{
		DealerID: DealerID,
		ClientID: ClientID,
		BranchID: BranchID,
		Instance: instance.New(DealerID, ClientID, BranchID),
		Study:    study.New(DealerID, ClientID, BranchID),
		Series:   serie.New(DealerID, ClientID, BranchID),
		V3:       v3.New(),
	}
}
