package study

import (
	"database/sql"

	"github.com/joseluis244/db2dbmod/databases/ortmysql/models"
	"github.com/joseluis244/db2dbmod/utils"
)

var intTags = map[string]bool{
	"0008,0020": true,
}

type StudyStruct struct {
	client *sql.DB
}

func New(client *sql.DB) *StudyStruct {
	return &StudyStruct{
		client: client,
	}
}

func (s *StudyStruct) GetStudyById(id int64) (models.OrtMySQLv1StudyType, error) {
	q := `SELECT 
	resourse.internalId as Id,
	resourse.publicId as StudyUuid,
	study.tagGroup as tagGroup,
	study.tagElement as tagElement,
	study.value as value
FROM Resources resourse
left join (select * from MainDicomTags) study on study.id = resourse.internalId
where internalId = ?;`
	rows, err := s.client.Query(q, id)
	if err != nil {
		return models.OrtMySQLv1StudyType{}, err
	}
	defer rows.Close()
	var result models.OrtMySQLv1StudyType = models.NewOrtMySQLv1StudyType(0, "", map[string]interface{}{})
	for rows.Next() {
		study := models.OrtStudyRaw{}
		if err := rows.Scan(&study.Id, &study.StudyUuid, &study.TagGroup, &study.TagElement, &study.Value); err != nil {
			return models.OrtMySQLv1StudyType{}, err
		}
		if result.StudyUuid == "" {
			result.Id = study.Id
			result.StudyUuid = study.StudyUuid
		}
		tag := utils.Dec2Hex(study.TagGroup, study.TagElement)
		result.Tags[tag] = utils.TagIntConverter(intTags, tag, study.Value)
	}
	return result, nil
}

func (s *StudyStruct) GetStudyByStudyUuid(uuid string) (models.OrtMySQLv1StudyType, error) {
	q := `SELECT 
	resourse.internalId as Id,
	resourse.publicId as StudyUuid,
	study.tagGroup as tagGroup,
	study.tagElement as tagElement,
	study.value as value
FROM Resources resourse
left join (select * from MainDicomTags) study on study.id = resourse.internalId
where publicId = ?;`
	rows, err := s.client.Query(q, uuid)
	if err != nil {
		return models.OrtMySQLv1StudyType{}, err
	}
	defer rows.Close()
	var result models.OrtMySQLv1StudyType = models.NewOrtMySQLv1StudyType(0, "", map[string]interface{}{})
	for rows.Next() {
		study := models.OrtStudyRaw{}
		if err := rows.Scan(&study.Id, &study.StudyUuid, &study.TagGroup, &study.TagElement, &study.Value); err != nil {
			return models.OrtMySQLv1StudyType{}, err
		}
		if study.StudyUuid == "" {
			result.Id = study.Id
			result.StudyUuid = study.StudyUuid
		}
		tag := utils.Dec2Hex(study.TagGroup, study.TagElement)
		result.Tags[tag] = utils.TagIntConverter(intTags, tag, study.Value)
	}
	return result, nil
}

func (s *StudyStruct) GetStudyByChangeRange(from int64, to int64) ([]models.OrtMySQLv1StudyType, error) {
	q := `SELECT 
	StudyResourse.internalId as Id,
	StudyResourse.publicId as StudyUuid,
	StudyTags.tagGroup as tagGroup,
	StudyTags.tagElement as tagElement,
	StudyTags.value as value
FROM Changes StudyChange
left join Resources StudyResourse on StudyResourse.internalId=StudyChange.internalId and StudyResourse.resourceType=1
left join MainDicomTags StudyTags on StudyResourse.internalId=StudyTags.id
where StudyChange.changeType=5 and (StudyChange.seq>=? and StudyChange.seq<=?)
order by StudyResourse.publicId;`
	rows, err := s.client.Query(q, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []models.OrtMySQLv1StudyType
	var currentStudy models.OrtMySQLv1StudyType = models.NewOrtMySQLv1StudyType(0, "", map[string]interface{}{})
	for rows.Next() {
		study := models.OrtStudyRaw{}
		if err := rows.Scan(&study.Id, &study.StudyUuid, &study.TagGroup, &study.TagElement, &study.Value); err != nil {
			return nil, err
		}
		if study.StudyUuid != currentStudy.StudyUuid {
			currentStudy = models.NewOrtMySQLv1StudyType(study.Id, study.StudyUuid, map[string]interface{}{})
			result = append(result, currentStudy)
		}
		tag := utils.Dec2Hex(study.TagGroup, study.TagElement)
		result[len(result)-1].Tags[tag] = utils.TagIntConverter(intTags, tag, study.Value)
	}
	return result, nil
}
