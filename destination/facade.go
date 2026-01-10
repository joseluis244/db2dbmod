package destination

import (
	"github.com/joseluis244/db2dbmod/destination/mongodb"
)

// MongoDB is a facade to access mongodb package functions
var MongoDB = struct {
	New func() *mongodb.MongoDB
}{
	New: mongodb.New,
}
