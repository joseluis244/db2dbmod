package main

import (
	"github.com/joseluis244/db2dbmod/builder"
	"github.com/joseluis244/db2dbmod/destination"
)

func main() {
	clouddb := destination.MongoDB.New()
	clouddb.Connect("mongodb://localhost:27017", "test")
	defer clouddb.Disconnect()
	destdb := destination.MongoDB.New()
	destdb.Connect("mongodb://localhost:27017", "test")
	defer destdb.Disconnect()
	instance, err := builder.BuildInstanceRaw([]byte(`{
		"Uuid": "test",
		"Ae": "test",
		"Tags": {},
		"StudyUuid": "test",
		"SerieUuid": "test",
		"Hash": "test",
		"Size": 0,
		"Path": "test",
		"Store": "test"
	}`), 0)
	if err != nil {
		panic(err)
	}
	destdb.InsertInstanceRawModel(instance)
	clouddb.InsertInstanceRawModel(instance)
}
