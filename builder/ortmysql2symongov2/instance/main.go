package instance

import (
	ortmysqlv1model "github.com/joseluis244/db2dbmod/databases/ortmysql/models"
	"github.com/joseluis244/db2dbmod/databases/symongov2/models"
)

type InstanceStruct struct {
	DealerID string
	ClientID string
	BranchID string
}

func New(DealerID string, ClientID string, BranchID string) *InstanceStruct {
	return &InstanceStruct{
		DealerID: DealerID,
		ClientID: ClientID,
		BranchID: BranchID,
	}
}

func (i *InstanceStruct) Move2Mongo(instance ortmysqlv1model.OrtMySQLv1InstanceType) (models.DestinationInstanceType, error) {
	Ae := ""
	Path := ""
	Store := "local"
	var instanceMongo models.DestinationInstanceType = models.NewDestinationInstanceRawType(i.DealerID, i.ClientID, i.BranchID, instance.InstanceUuid, Ae, instance.Tags, instance.StudyUuid, instance.SerieUuid, instance.Hash, instance.Size, Path, Store)
	return instanceMongo, nil
}

func (i *InstanceStruct) MoveMany2Mongo(instances []ortmysqlv1model.OrtMySQLv1InstanceType) ([]models.DestinationInstanceType, error) {
	if len(instances) == 0 {
		return []models.DestinationInstanceType{}, nil
	}
	var instancesMongo []models.DestinationInstanceType
	for _, instance := range instances {
		instanceMongo, err := i.Move2Mongo(instance)
		if err != nil {
			return nil, err
		}
		instancesMongo = append(instancesMongo, instanceMongo)
	}
	return instancesMongo, nil
}
