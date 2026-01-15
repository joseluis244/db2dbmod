package main

import (
	"fmt"
	"time"

	"github.com/joseluis244/db2dbmod/builder"
	"github.com/joseluis244/db2dbmod/dbsource"
	"github.com/joseluis244/db2dbmod/models"
)

func main() {
	t := time.Now()
	mysql := dbsource.MySQLv1.New()
	err := mysql.Connect("medicaresoftmysql:MedicareSoft203$@tcp(127.0.0.1:3308)/symphony")
	if err != nil {
		panic(err)
	}
	defer mysql.Disconnect()
	Builder := builder.Mysqlv12Mogo.New("MedicareSoftMongo", "test1", "sucursal1")
	lastChange, err := mysql.Change.LastChange()
	if err != nil {
		panic(err)
	}
	//procesar instances
	instances, err := mysql.Instance.GetInstanceByChangeRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	var ModelsInstances []models.DestinationInstanceType
	for _, instance := range instances {
		model, err := Builder.Instance.Move2Mongo(instance)
		if err != nil {
			panic(err)
		}
		ModelsInstances = append(ModelsInstances, model)
	}
	fmt.Println(len(ModelsInstances))
	//procesar series
	series, err := mysql.Serie.GetSerieByChangeRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	var ModelsSeries []models.DestinationSeriesType
	for _, serie := range series {
		model, err := Builder.Series.Move2Mongo(serie)
		if err != nil {
			panic(err)
		}
		ModelsSeries = append(ModelsSeries, model)
	}
	fmt.Println(len(ModelsSeries))
	//procesar studies
	studies, err := mysql.Study.GetStudyByChangeRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	var ModelsStudies []models.DestinationStudyType
	for _, study := range studies {
		model, err := Builder.Study.Move2Mongo(study)
		if err != nil {
			panic(err)
		}
		ModelsStudies = append(ModelsStudies, model)
	}
	fmt.Println(len(ModelsStudies))
	//time consume
	fmt.Println(time.Since(t))
}
