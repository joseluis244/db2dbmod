package v3

import (
	"sort"
	"strconv"

	"github.com/joseluis244/sydatabasesmod/symongov1/models"
)

type V3Struct struct {
}

func New() *V3Struct {
	return &V3Struct{}
}

func (v *V3Struct) V3Builder(study models.SyMongoV1StudyType, series []models.SyMongoV1SeriesType, instances []models.SyMongoV1InstanceType) models.SyMongoV1V3Type {
	// 1. Agrupar instancias por SerieUuid
	instancesMap := make(map[string][]models.SyMongoV1InstanceType)
	for _, instance := range instances {
		instancesMap[instance.SerieUuid] = append(instancesMap[instance.SerieUuid], instance)
	}

	// 2. Construir v3series
	v3series := make([]models.SyMongoV1V3SeriesType, len(series))
	for i, serie := range series {
		// Obtener y ordenar instancias de esta serie
		serieInstances := instancesMap[serie.SerieUuid]
		sortInstancesByInstanceNumber(serieInstances)

		v3instances := make([]models.SyMongoV1V3InstanceType, len(serieInstances))
		for j, instance := range serieInstances {
			v3instances[j] = models.NewSyMongoV1V3InstanceType(
				instance.Uuid, instance.StudyUuid, instance.SerieUuid,
				instance.Size, instance.Hash, instance.Path,
				instance.Id, instance.Tags, instance.Ae,
			)
		}

		v3series[i] = models.NewSyMongoV1V3SeriesType(
			serie.StudyUuid, serie.SerieUuid, serie.Id, serie.Tags, v3instances,
		)
	}

	return models.NewSyMongoV1V3Type(study.StudyUuid, true, study.Id, 0, study.UpdateAt, v3series, study.Tags)
}

func sortInstancesByInstanceNumber(instances []models.SyMongoV1InstanceType) {
	if len(instances) <= 1 {
		return
	}

	sort.Slice(instances, func(i, j int) bool {
		return getInstanceNumber(instances[i]) < getInstanceNumber(instances[j])
	})
}

func getInstanceNumber(inst models.SyMongoV1InstanceType) int {
	val, ok := inst.Tags["0020,0013"]
	if !ok {
		return 0
	}

	switch v := val.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float64:
		return int(v)
	case string:
		n, _ := strconv.Atoi(v)
		return n
	default:
		return 0
	}
}
