package ortmysql2symongov1

import (
	"github.com/joseluis244/db2dbmod/builder/ortmysql2symongov1/instance"
	"github.com/joseluis244/db2dbmod/builder/ortmysql2symongov1/serie"
	"github.com/joseluis244/db2dbmod/builder/ortmysql2symongov1/study"
	v3 "github.com/joseluis244/db2dbmod/builder/ortmysql2symongov1/v3"
)

type BuilderStruct struct {
	Instance *instance.InstanceStruct
	Study    *study.StudyStruct
	Series   *serie.SerieStruct
	V3       *v3.V3Struct
}

func New() *BuilderStruct {
	return &BuilderStruct{
		Instance: instance.New(),
		Study:    study.New(),
		Series:   serie.New(),
		V3:       v3.New(),
	}
}
