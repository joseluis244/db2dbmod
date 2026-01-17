package study

import (
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1/study"
	"github.com/joseluis244/db2dbmod/models"
)

type StudyStruct struct {
	DealerID string
	ClientID string
	BranchID string
}

func New(DealerID string, ClientID string, BranchID string) *StudyStruct {
	return &StudyStruct{
		DealerID: DealerID,
		ClientID: ClientID,
		BranchID: BranchID,
	}
}

func (s *StudyStruct) Move2Mongo(st study.SourceMySQLv1StudyType) (models.DestinationStudyType, error) {
	return models.NewDestinationStudyType(s.DealerID, s.ClientID, s.BranchID, st.StudyUuid, st.Tags), nil
}

func (s *StudyStruct) MoveMany2Mongo(studies []study.SourceMySQLv1StudyType) ([]models.DestinationStudyType, error) {
	if len(studies) == 0 {
		return []models.DestinationStudyType{}, nil
	}
	var studiesMongo []models.DestinationStudyType
	for _, st := range studies {
		studyMongo, err := s.Move2Mongo(st)
		if err != nil {
			return nil, err
		}
		studiesMongo = append(studiesMongo, studyMongo)
	}
	return studiesMongo, nil
}
