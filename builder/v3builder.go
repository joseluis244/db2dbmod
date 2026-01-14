package builder

import (
	"sort"
	"strconv"
	"sync"

	"github.com/joseluis244/db2dbmod/models"
)

func V3Builder(study models.DestinationStudyType, series []models.DestinationSeriesType, instances []models.DestinationInstanceType) models.DestinationV3Type {
	instancesMap := sortInstancesBySerieUuid(instances)
	v3series := make([]models.DestinationV3SeriesType, len(series))
	for i, serie := range series {
		v3instances := make([]models.DestinationV3InstanceType, len(instancesMap[serie.SerieUuid]))
		for j, instance := range instancesMap[serie.SerieUuid] {
			v3instance := models.NewDestinationV3InstanceType(instance.Uuid, instance.Ae, instance.Hash, instance.Size, instance.Path, instance.Store, instance.Tags)
			v3instances[j] = v3instance
		}
		v3serie := models.NewDestinationV3SeriesType(serie.SerieUuid, serie.StudyUuid, serie.Tags, v3instances)
		v3series[i] = v3serie
	}
	v3study := models.NewDestinationV3Type(study.DealerID, study.ClientID, study.BranchID, study.StudyUuid, study.Tags, study.CreatedAt, study.UpdatedAt, v3series)
	return v3study
}

func sortInstancesBySerieUuid(instances []models.DestinationInstanceType) map[string][]models.DestinationInstanceType {
	instancesMap := make(map[string][]models.DestinationInstanceType)
	for _, instance := range instances {
		instancesMap[instance.SerieUuid] = append(instancesMap[instance.SerieUuid], instance)
	}
	wg := sync.WaitGroup{}
	wg.Add(len(instancesMap))
	for serieUuid, seriesInstances := range instancesMap {
		go func(uuid string, instances []models.DestinationInstanceType) {
			instancesMap[uuid] = sortInstancesByInstanceNumber(instances)
			wg.Done()
		}(serieUuid, seriesInstances)
	}
	wg.Wait()
	return instancesMap
}

func sortInstancesByInstanceNumber(instances []models.DestinationInstanceType) []models.DestinationInstanceType {
	// Pre-calcular los números para evitar conversiones repetidas
	type instanceWithNumber struct {
		instance models.DestinationInstanceType
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
