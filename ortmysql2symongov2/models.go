package ortmysql2symongov2

import (
	ortmysqlv1 "github.com/joseluis244/sydatabasesmod/ortmysql/models"
)

type ChanelChangesToBuild struct {
	Studies    []ortmysqlv1.OrtMySQLv1StudyType
	Series     []ortmysqlv1.OrtMySQLv1SerieType
	Instances  []ortmysqlv1.OrtMySQLv1InstanceType
	LastChange int64
}
