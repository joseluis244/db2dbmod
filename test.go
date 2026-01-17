package main

import (
	"fmt"
	"time"

	"github.com/joseluis244/db2dbmod/builder"
	"github.com/joseluis244/db2dbmod/builder/mysqlv12mogo"
	"github.com/joseluis244/db2dbmod/dbsource"
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1"
	"github.com/joseluis244/db2dbmod/destination"
)

func main() {
	t := time.Now()
	mysql := dbsource.MySQLv1.New()
	err := mysql.Connect("medicaresoftmysql:MedicareSoft203$@tcp(127.0.0.1:3308)/symphony")
	if err != nil {
		panic(err)
	}
	defer mysql.Disconnect()
	mongo := destination.MongoDB.New()
	err = mongo.Connect("mongodb://localhost:27018", "test1")
	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect()
	Builder := builder.Mysqlv12Mogo.New("MedicareSoftMongo", "test1", "sucursal1")
	//
	chBuilded := make(chan mysqlv12mogo.ChanelChangesToBuild, 1000)
	//
	go findChanges(mysql, chBuilded)
	//process
	go processChanges(Builder, chBuilded)
	//time consume
	fmt.Println(time.Since(t))
	select {}
}

func findChanges(mysql *mysqlv1.MySQLv1, chBuilded chan<- mysqlv12mogo.ChanelChangesToBuild) {

	lastChange, err := mysql.Change.LastChange()
	if err != nil {
		panic(err)
	}
	instances, err := mysql.Instance.GetInstanceByChangeRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	series, err := mysql.Serie.GetSerieByChangeRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	studies, err := mysql.Study.GetStudyByChangeRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	chBuilded <- mysqlv12mogo.ChanelChangesToBuild{
		Studies:    studies,
		Series:     series,
		Instances:  instances,
		LastChange: lastChange,
	}
}

func processChanges(builder *mysqlv12mogo.BuilderStruct, chBuilded <-chan mysqlv12mogo.ChanelChangesToBuild) {
	for tobuild := range chBuilded {
		studyMongo, err := builder.Study.MoveMany2Mongo(tobuild.Studies)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(studyMongo))
		seriesMongo, err := builder.Series.MoveMany2Mongo(tobuild.Series)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(seriesMongo))
		instanceMongo, err := builder.Instance.MoveMany2Mongo(tobuild.Instances)
		if err != nil {
			panic(err)
		}
		fmt.Println(len(instanceMongo))
		fmt.Println(tobuild.LastChange)
	}
}

// func v3builder(mongo *mongodb.MongoDB) {
// 	tobuild, err := mongo.Study.GetToBuild(bson.M{})
// 	if err != nil {
// 		panic(err)
// 	}
// 	V3s := []models.DestinationV3Type{}
// 	for _, study := range tobuild {
// 		V3 := builder.V3Builder(study.Study, study.Series, study.Instances)
// 		V3s = append(V3s, V3)
// 	}
// 	err = mongo.V3.UpsertV3s(V3s)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(len(V3s))
// }
