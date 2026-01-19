package serie

import (
	"github.com/joseluis244/db2dbmod/databases/ortmysql/models"
	modelssymongov1 "github.com/joseluis244/db2dbmod/databases/symongov1/models"
)

type SerieStruct struct {
}

func New() *SerieStruct {
	return &SerieStruct{}
}

func (s *SerieStruct) Build(serie models.OrtMySQLv1SerieType) (modelssymongov1.SyMongoV1SeriesType, error) {
	result := modelssymongov1.NewSyMongoV1SeriesType(serie.SerieUuid, serie.Id, serie.StudyUuid, serie.Tags)
	return result, nil
}

func (s *SerieStruct) BuildMany(series []models.OrtMySQLv1SerieType) ([]modelssymongov1.SyMongoV1SeriesType, error) {
	if len(series) == 0 {
		return []modelssymongov1.SyMongoV1SeriesType{}, nil
	}
	var seriesMongo []modelssymongov1.SyMongoV1SeriesType
	for _, serie := range series {
		serieMongo, err := s.Build(serie)
		if err != nil {
			return nil, err
		}
		seriesMongo = append(seriesMongo, serieMongo)
	}
	return seriesMongo, nil
}
