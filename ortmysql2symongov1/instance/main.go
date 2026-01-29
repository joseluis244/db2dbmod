package instance

import (
	"fmt"

	"github.com/joseluis244/sydatabasesmod/ortmysql/models"
	modelssymongov1 "github.com/joseluis244/sydatabasesmod/symongov1/models"
)

type InstanceStruct struct {
}

func New() *InstanceStruct {
	return &InstanceStruct{}
}

func (i *InstanceStruct) Build(instance models.OrtMySQLv1InstanceType) (modelssymongov1.SyMongoV1InstanceType, error) {
	path := fmt.Sprintf("%s/%s/%s", instance.FileUuid[:2], instance.FileUuid[2:4], instance.FileUuid)
	result := modelssymongov1.NewSyMongoV1InstanceType(instance.FileUuid, instance.AE, 0, instance.Hash, instance.Id, path, instance.SerieUuid, instance.Size, instance.StudyUuid, 0, instance.Tags)
	return result, nil
}

func (i *InstanceStruct) BuildMany(instances []models.OrtMySQLv1InstanceType) ([]modelssymongov1.SyMongoV1InstanceType, error) {
	if len(instances) == 0 {
		return []modelssymongov1.SyMongoV1InstanceType{}, nil
	}
	var instancesMongo []modelssymongov1.SyMongoV1InstanceType
	for _, instance := range instances {
		instanceMongo, err := i.Build(instance)
		if err != nil {
			return nil, err
		}
		instancesMongo = append(instancesMongo, instanceMongo)
	}
	return instancesMongo, nil
}
