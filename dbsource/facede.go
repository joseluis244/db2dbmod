package dbsource

import "github.com/joseluis244/db2dbmod/dbsource/mysqlv1"

var MySQLv1 = struct {
	New func() *mysqlv1.MySQLv1
}{
	New: mysqlv1.New,
}
