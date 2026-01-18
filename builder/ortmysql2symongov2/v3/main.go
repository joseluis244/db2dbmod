package v3

import (
	"sort"
	"strconv"
	"sync"

	"github.com/joseluis244/db2dbmod/databases/symongov2/models"
)

type V3Struct struct {
}

func New() *V3Struct {
	return &V3Struct{}
}

func (v *V3Struct) V3Builder(study models.SyMongoV2StudyType, series []models.SyMongoV2SeriesType, instances []models.SyMongoV2InstanceType) models.SyMongoV2V3Type {
	instancesMap := sortInstancesBySerieUuid(instances)
	v3series := make([]models.SyMongoV2V3SeriesType, len(series))
	for i, serie := range series {
		v3instances := make([]models.SyMongoV2V3InstanceType, len(instancesMap[serie.SerieUuid]))
		for j, instance := range instancesMap[serie.SerieUuid] {
			v3instance := models.NewSyMongoV2V3InstanceType(instance.InstanceUuid, instance.Ae, instance.Hash, instance.Size, instance.Path, instance.Store, instance.Tags)
			v3instances[j] = v3instance
		}
		v3serie := models.NewSyMongoV2V3SeriesType(serie.SerieUuid, serie.Tags, v3instances)
		v3series[i] = v3serie
	}
	v3study := models.NewSyMongoV2V3Type(study.DealerID, study.ClientID, study.BranchID, study.StudyUuid, study.Tags, study.CreatedAt, study.UpdatedAt, v3series)
	return v3study
}

func sortInstancesBySerieUuid(instances []models.SyMongoV2InstanceType) map[string][]models.SyMongoV2InstanceType {
	instancesMap := make(map[string][]models.SyMongoV2InstanceType)
	for _, instance := range instances {
		instancesMap[instance.SerieUuid] = append(instancesMap[instance.SerieUuid], instance)
	}
	sortedMap := make(map[string][]models.SyMongoV2InstanceType)
	wg := sync.WaitGroup{}
	var mu sync.Mutex
	wg.Add(len(instancesMap))
	for serieUuid, seriesInstances := range instancesMap {
		go func(uuid string, instances []models.SyMongoV2InstanceType) {
			sorted := sortInstancesByInstanceNumber(instances)
			mu.Lock()
			sortedMap[uuid] = sorted
			mu.Unlock()
			wg.Done()
		}(serieUuid, seriesInstances)
	}
	wg.Wait()
	return sortedMap
}

func sortInstancesByInstanceNumber(instances []models.SyMongoV2InstanceType) []models.SyMongoV2InstanceType {
	// Pre-calcular los números para evitar conversiones repetidas
	type instanceWithNumber struct {
		instance models.SyMongoV2InstanceType
		number   int
	}

	items := make([]instanceWithNumber, len(instances))
	for i, inst := range instances {
		number := 0
		if val, ok := inst.Tags["0020,0013"].(int); ok {
			number = val
		} else if val, ok := inst.Tags["0020,0013"].(string); ok {
			number, _ = strconv.Atoi(val)
		}
		items[i] = instanceWithNumber{inst, number}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].number < items[j].number
	})

	// Extraer instances ordenadas
	for i, item := range items {
		instances[i] = item.instance
	}

	return instances
}
