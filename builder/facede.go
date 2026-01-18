package builder

import "github.com/joseluis244/db2dbmod/builder/ortmysql2symongov2"

var OrtMySQL2Mongo = struct {
	New func(DealerID string, ClientID string, BranchID string) *ortmysql2symongov2.BuilderStruct
}{
	New: ortmysql2symongov2.New,
}
