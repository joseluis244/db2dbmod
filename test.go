package main

import (
	"fmt"
	"time"

	"github.com/joseluis244/db2dbmod/builder"
	"github.com/joseluis244/db2dbmod/builder/mysqlv12mogo"
	"github.com/joseluis244/db2dbmod/dbsource"
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1"
)

func main() {
	t := time.Now()
	mysql := dbsource.MySQLv1.New()
	mysql.Connect("medicaresoftmysql:MedicareSoft203$@tcp(127.0.0.1:3308)/symphony")
	defer mysql.Disconnect()
	Builder := builder.Mysqlv12Mogo.New("MedicareSoftMongo", "test1", "sucursal1")
	lastChange, err := mysql.Change.LastChange()
	if err != nil {
		panic(err)
	}
	changes, err := mysql.Change.ChangesRange(0, lastChange)
	if err != nil {
		panic(err)
	}
	for _, change := range changes {
		switch change.ResourceType {
		case "1":
			estudio(change.InternalId, mysql, Builder)
		case "2":
			serie(change.InternalId, mysql, Builder)
		case "3":
			instance(change.InternalId, mysql, Builder)
		}
	}
	fmt.Println(time.Since(t))
}

func estudio(internalId int64, mysql *mysqlv1.MySQLv1, Builder *mysqlv12mogo.BuilderStruct) {
	study, err := mysql.Study.GetStudyById(internalId)
	if err != nil {
		panic(err)
	}
	model, err := Builder.Study.Move2Mongo(study)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)
}
func serie(internalId int64, mysql *mysqlv1.MySQLv1, Builder *mysqlv12mogo.BuilderStruct) {
	serie, err := mysql.Serie.GetSerieById(internalId)
	if err != nil {
		panic(err)
	}
	model, err := Builder.Series.Move2Mongo(serie)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)
}
func instance(internalId int64, mysql *mysqlv1.MySQLv1, Builder *mysqlv12mogo.BuilderStruct) {
	instance, err := mysql.Instance.GetInstanceById(internalId)
	if err != nil {
		panic(err)
	}
	model, err := Builder.Instance.Move2Mongo(instance)
	if err != nil {
		panic(err)
	}
	fmt.Println(model)
}
