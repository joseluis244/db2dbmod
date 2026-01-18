package databases

import (
	"github.com/joseluis244/db2dbmod/databases/ortmysql"
	"github.com/joseluis244/db2dbmod/databases/symongov2"
)

var SyMongoV2 = struct {
	New func() *symongov2.SyMongoV2
}{
	New: symongov2.New,
}

var OrtMySql = struct {
	New func() *ortmysql.OrtMySQL
}{
	New: ortmysql.New,
}
