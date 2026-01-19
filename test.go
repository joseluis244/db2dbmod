package main

import (
	"fmt"
	"time"

	"github.com/joseluis244/db2dbmod/databases"
)

var t = time.Now()

func main() {
	localsource := databases.OrtMySql.New()
	err := localsource.Connect("medicaltecmysql:Medicaltec310188$@tcp(127.0.0.1:3306)/medicaltec")
	if err != nil {
		panic(err)
	}
	defer localsource.Disconnect()
	lastChange, err := localsource.Change.LastChange()
	if err != nil {
		panic(err)
	}
	fmt.Println("Last change: ", lastChange)
	studies, err := localsource.Study.GetStudyByChangeRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	fmt.Println("Studies: ", studies[0])
	series, err := localsource.Serie.GetSerieByChangeRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	fmt.Println("Series: ", series[0])
	instances, err := localsource.Instance.GetInstanceByChangeRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	fmt.Println("Instances: ", instances[0])
	fmt.Println("Time consume: ", time.Since(t))
}

// func findChanges(localsource *ortmysql.OrtMySQL, chBuilded chan<- OrtMySQL2Mongo.ChanelChangesToBuild) {

// 	lastChange, err := localsource.Change.LastChange()
// 	if err != nil {
// 		panic(err)
// 	}
// 	instances, err := localsource.Instance.GetInstanceByChangeRange(0, lastChange)
// 	if err != nil {
// 		panic(err)
// 	}
// 	series, err := localsource.Serie.GetSerieByChangeRange(0, lastChange)
// 	if err != nil {
// 		panic(err)
// 	}
// 	studies, err := localsource.Study.GetStudyByChangeRange(0, lastChange)
// 	if err != nil {
// 		panic(err)
// 	}
// 	chBuilded <- OrtMySQL2Mongo.ChanelChangesToBuild{
// 		Studies:    studies,
// 		Series:     series,
// 		Instances:  instances,
// 		LastChange: lastChange,
// 	}
// }

// func processChanges(builder *OrtMySQL2Mongo.BuilderStruct, chBuilded <-chan OrtMySQL2Mongo.ChanelChangesToBuild, localdestination *symongov2.SyMongoV2) {
// 	for tobuild := range chBuilded {
// 		studyMongo, err := builder.Study.MoveMany2Mongo(tobuild.Studies)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(len(studyMongo))
// 		seriesMongo, err := builder.Series.MoveMany2Mongo(tobuild.Series)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(len(seriesMongo))
// 		instanceMongo, err := builder.Instance.MoveMany2Mongo(tobuild.Instances)
// 		if err != nil {
// 			panic(err)
// 		}
// 		fmt.Println(len(instanceMongo))
// 		BUILDED, err := localdestination.Study.GetToBuild(bson.M{})
// 		if err != nil {
// 			panic(err)
// 		}
// 		V3s := []models.DestinationV3Type{}
// 		for _, study := range BUILDED {
// 			V3 := builder.V3.V3Builder(study.Study, study.Series, study.Instances)
// 			V3s = append(V3s, V3)
// 		}
// 		fmt.Println(len(V3s))
// 		fmt.Println(tobuild.LastChange)
// 		fmt.Println("Time consume: ", time.Since(t))
// 	}
// }
