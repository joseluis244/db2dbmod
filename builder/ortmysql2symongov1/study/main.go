package study

import (
	"github.com/joseluis244/db2dbmod/databases/ortmysql/models"
	modelssymongov1 "github.com/joseluis244/db2dbmod/databases/symongov1/models"
)

type StudyStruct struct {
}

func New() *StudyStruct {
	return &StudyStruct{}
}

func (s *StudyStruct) Build(study models.OrtMySQLv1StudyType) (modelssymongov1.SyMongoV1StudyType, error) {
	result := modelssymongov1.NewSyMongoV1StudyType(study.StudyUuid, study.Id, study.Tags)
	return result, nil
}

func (s *StudyStruct) BuildMany(studies []models.OrtMySQLv1StudyType) ([]modelssymongov1.SyMongoV1StudyType, error) {
	if len(studies) == 0 {
		return []modelssymongov1.SyMongoV1StudyType{}, nil
	}
	var studiesMongo []modelssymongov1.SyMongoV1StudyType
	for _, study := range studies {
		studyMongo, err := s.Build(study)
		if err != nil {
			return nil, err
		}
		studiesMongo = append(studiesMongo, studyMongo)
	}
	return studiesMongo, nil
}
