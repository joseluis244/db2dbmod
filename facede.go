package db2dbmod

import (
	"github.com/joseluis244/db2dbmod/ortmysql2symongov1"
	"github.com/joseluis244/db2dbmod/ortmysql2symongov2"
)

var OrtMySQL2Mongo = struct {
	New func(DealerID string, ClientID string, BranchID string) *ortmysql2symongov2.BuilderStruct
}{
	New: ortmysql2symongov2.New,
}

var OrtMySQL2MongoV1 = struct {
	New func() *ortmysql2symongov1.BuilderStruct
}{
	New: ortmysql2symongov1.New,
}
