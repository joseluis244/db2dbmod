package serie

import (
	ortmysqlv1 "github.com/joseluis244/db2dbmod/databases/ortmysql/models"
	symongov2model "github.com/joseluis244/db2dbmod/databases/symongov2/models"
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

func (s *SerieStruct) Move2Mongo(serie ortmysqlv1.OrtMySQLv1SerieType) (symongov2model.SyMongoV2SeriesType, error) {
	return symongov2model.NewSyMongoV2SeriesType(s.DealerID, s.ClientID, s.BranchID, serie.StudyUuid, serie.SerieUuid, serie.Tags), nil
}

func (s *SerieStruct) MoveMany2Mongo(series []ortmysqlv1.OrtMySQLv1SerieType) ([]symongov2model.SyMongoV2SeriesType, error) {
	if len(series) == 0 {
		return []symongov2model.SyMongoV2SeriesType{}, nil
	}
	var seriesMongo []symongov2model.SyMongoV2SeriesType
	for _, serie := range series {
		serieMongo, err := s.Move2Mongo(serie)
		if err != nil {
			return nil, err
		}
		seriesMongo = append(seriesMongo, serieMongo)
	}
	return seriesMongo, nil
}
