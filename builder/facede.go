package builder

import "github.com/joseluis244/db2dbmod/builder/mysqlv12mogo"

var Mysqlv12Mogo = struct {
	New func(DealerID string, ClientID string, BranchID string) *mysqlv12mogo.BuilderStruct
}{
	New: mysqlv12mogo.New,
}
