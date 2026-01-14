package serie

import (
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1/serie"
	"github.com/joseluis244/db2dbmod/models"
)

type SerieStruct struct {
	DealerID string
	ClientID string
	BranchID string
}

func New(DealerID string, ClientID string, BranchID string) *SerieStruct {
	return &SerieStruct{
		DealerID: DealerID,
		ClientID: ClientID,
		BranchID: BranchID,
	}
}

func (s *SerieStruct) Move2Mongo(serie serie.SourceMySQLv1SerieType) (models.DestinationSeriesType, error) {
	return models.NewDestinationSeriesRawType(s.DealerID, s.ClientID, s.BranchID, serie.StudyUuid, serie.SerieUuid, serie.Tags), nil
}

func (s *SerieStruct) MoveMany2Mongo(series []serie.SourceMySQLv1SerieType) ([]models.DestinationSeriesType, error) {
	var seriesMongo []models.DestinationSeriesType
	for _, serie := range series {
		serieMongo, err := s.Move2Mongo(serie)
		if err != nil {
			return nil, err
		}
		seriesMongo = append(seriesMongo, serieMongo)
	}
	return seriesMongo, nil
}
