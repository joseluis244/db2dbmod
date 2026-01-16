package main

import (
	"fmt"
	"time"

	"github.com/joseluis244/db2dbmod/builder"
	"github.com/joseluis244/db2dbmod/destination"
	"github.com/joseluis244/db2dbmod/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// func main() {
// 	t := time.Now()
// 	mysql := dbsource.MySQLv1.New()
// 	err := mysql.Connect("medicaresoftmysql:MedicareSoft203$@tcp(127.0.0.1:3308)/symphony")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer mysql.Disconnect()
// 	mongo := destination.MongoDB.New()
// 	err = mongo.Connect("mongodb://localhost:27018", "test1")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer mongo.Disconnect()
// 	Builder := builder.Mysqlv12Mogo.New("MedicareSoftMongo", "test1", "sucursal1")
// 	lastChange, err := mysql.Change.LastChange()
// 	if err != nil {
// 		panic(err)
// 	}
// 	//procesar instances
// 	instances, err := mysql.Instance.GetInstanceByChangeRange(0, lastChange)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var ModelsInstances []models.DestinationInstanceType
// 	InstancesCreated := time.Now().Unix()
// 	for _, instance := range instances {
// 		model, err := Builder.Instance.Move2Mongo(instance)
// 		if err != nil {
// 			panic(err)
// 		}
// 		model.CreatedAt = InstancesCreated
// 		model.UpdatedAt = InstancesCreated
// 		ModelsInstances = append(ModelsInstances, model)
// 	}
// 	err = mongo.Instance.UpsertInstances(ModelsInstances)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(len(ModelsInstances))
// 	//procesar series
// 	series, err := mysql.Serie.GetSerieByChangeRange(0, lastChange)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var ModelsSeries []models.DestinationSeriesType
// 	SeriesCreated := time.Now().Unix()
// 	for _, serie := range series {
// 		model, err := Builder.Series.Move2Mongo(serie)
// 		if err != nil {
// 			panic(err)
// 		}
// 		model.CreatedAt = SeriesCreated
// 		model.UpdatedAt = SeriesCreated
// 		ModelsSeries = append(ModelsSeries, model)
// 	}
// 	err = mongo.Series.UpsertSeries(ModelsSeries)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(len(ModelsSeries))
// 	//procesar studies
// 	studies, err := mysql.Study.GetStudyByChangeRange(0, lastChange)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var ModelsStudies []models.DestinationStudyType
// 	StudiesCreated := time.Now().Unix()
// 	for _, study := range studies {
// 		model, err := Builder.Study.Move2Mongo(study)
// 		if err != nil {
// 			panic(err)
// 		}
// 		model.CreatedAt = StudiesCreated
// 		model.UpdatedAt = StudiesCreated
// 		ModelsStudies = append(ModelsStudies, model)
// 	}
// 	err = mongo.Study.UpsertStudies(ModelsStudies)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(len(ModelsStudies))
// 	//time consume
// 	fmt.Println(time.Since(t))
// }

func main() {
	t := time.Now()
	mongo := destination.MongoDB.New()
	err := mongo.Connect("mongodb://localhost:27018", "test1")
	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect()
	tobuild, err := mongo.Study.GetToBuild(bson.M{})
	if err != nil {
		panic(err)
	}
	V3s := []models.DestinationV3Type{}
	for _, study := range tobuild {
		V3 := builder.V3Builder(study.Study, study.Series, study.Instances)
		V3s = append(V3s, V3)
	}
	err = mongo.V3.UpsertV3s(V3s)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(V3s))
	fmt.Println(time.Since(t))
}
