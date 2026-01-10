package builder

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/joseluis244/db2dbmod/models"
)

func BuildInstanceRaw(data []byte, buildTime int64) (models.DestinationInstanceRawType, error) {
	var instance map[string]interface{}
	err := json.Unmarshal(data, &instance)
	if err != nil {
		return models.DestinationInstanceRawType{}, err
	}

	requiredKeys := []string{"Uuid", "Ae", "Tags", "StudyUuid", "SerieUuid", "Hash", "Size", "Path", "Store"}
	for _, key := range requiredKeys {
		if val, ok := instance[key]; !ok || val == nil {
			return models.DestinationInstanceRawType{}, fmt.Errorf("missing or nil field: %s", key)
		}
	}

	size, ok := instance["Size"].(float64)
	if !ok {
		return models.DestinationInstanceRawType{}, fmt.Errorf("invalid type for field: Size, expected number")
	}

	return models.DestinationInstanceRawType{
		Uuid:      instance["Uuid"].(string),
		Ae:        instance["Ae"].(string),
		Tags:      instance["Tags"].(map[string]interface{}),
		CreatedAt: buildTime,
		UpdatedAt: buildTime,
		StudyUuid: instance["StudyUuid"].(string),
		SerieUuid: instance["SerieUuid"].(string),
		Hash:      instance["Hash"].(string),
		Size:      int64(size),
		Path:      instance["Path"].(string),
		Store:     instance["Store"].(string),
	}, nil
}

func InstanceRawSortByPosition(instances []models.DestinationInstanceRawType) map[string][]models.DestinationInstanceRawType {
	kv := make(map[string][]models.DestinationInstanceRawType)
	currentkey := ""
	for _, instance := range instances {
		if instance.StudyUuid+"_"+instance.SerieUuid != currentkey {
			currentkey = instance.StudyUuid + "_" + instance.SerieUuid
			kv[currentkey] = []models.DestinationInstanceRawType{}
		}
		kv[currentkey] = append(kv[currentkey], instance)
	}
	for _, instances := range kv {
		go func(INSTANCES []models.DestinationInstanceRawType) {
			sort.Slice(INSTANCES, func(i, j int) bool {
				return INSTANCES[i].Tags["0020,0013"].(int) < INSTANCES[j].Tags["0020,0013"].(int)
			})
			key := INSTANCES[0].StudyUuid + "_" + INSTANCES[0].SerieUuid
			kv[key] = INSTANCES
		}(instances)
	}
	return kv
}
