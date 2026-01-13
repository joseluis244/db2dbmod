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

func test() {
	arrCompleto := []interface{}{}
	lenSt := funcionparacontarestudios(arrCompleto)
	lenSer := funcionparacontarseries(arrCompleto)
	lenInst := funcionparacontarinstancias(arrCompleto)
	var Stchan = make(chan interface{}, lenSt)
	var Serchan = make(chan interface{}, lenSer)
	var Instchan = make(chan interface{}, lenInst)
	for _, item := range arrCompleto {
		switch item.T {
		case "Study":
			go func(item interface{}) {
				//hacer algo co el estudio
				Stchan <- item //solo ejemplo
			}(item)
		case "Series":
			go func(item interface{}) {
				//hacer algo co la serie
				Serchan <- item //solo ejemplo
			}(item)
		case "Instance":
			go func(item interface{}) {
				//hacer algo co la instancia
				Instchan <- item //solo ejemplo
			}(item)
		}
	}
	for i := len(arrCompleto); i > 0; i-- {
		select {
		case <-Stchan:
			//hacer algo co el estudio
		case <-Serchan:
			//hacer algo co la serie
		case <-Instchan:
			//hacer algo co la instancia
		}
	}
}
