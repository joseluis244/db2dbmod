package instance

import (
	"github.com/joseluis244/db2dbmod/databases/ortmysql/models"
	modelssymongov1 "github.com/joseluis244/db2dbmod/databases/symongov1/models"
)

type InstanceStruct struct {
}

func New() *InstanceStruct {
	return &InstanceStruct{}
}

func (i *InstanceStruct) Build(instance models.OrtMySQLv1InstanceType) (modelssymongov1.SyMongoV1InstanceType, error) {
	result := modelssymongov1.NewSyMongoV1InstanceType(instance.InstanceUuid, instance.AE, 0, instance.Hash, instance.Id, "", instance.SerieUuid, instance.Size, instance.StudyUuid, 0, instance.Tags)
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
