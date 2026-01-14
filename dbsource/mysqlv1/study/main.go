package study

import (
	"database/sql"
	"strconv"

	"github.com/joseluis244/db2dbmod/utils"
)

type SourceMySQLv1StudyType struct {
	StudyUuid string                 `json:"StudyUuid"`
	Tags      map[string]interface{} `json:"Tags"`
}

type raw struct {
	StudyUuid  string `json:"StudyUuid"`
	TagGroup   int    `json:"TagGroup"`
	TagElement int    `json:"TagElement"`
	Value      string `json:"Value"`
}

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

func (s *StudyStruct) GetStudyById(id int64) (SourceMySQLv1StudyType, error) {
	q := `SELECT 
resourse.publicId as StudyUuid,
study.tagGroup as tagGroup,
study.tagElement as tagElement,
study.value as value
FROM Resources resourse
left join (select * from MainDicomTags) study on study.id = resourse.internalId
where internalId = ?;`
	rows, err := s.client.Query(q, id)
	if err != nil {
		return SourceMySQLv1StudyType{}, err
	}
	defer rows.Close()
	var result SourceMySQLv1StudyType = SourceMySQLv1StudyType{
		StudyUuid: "",
		Tags:      map[string]interface{}{},
	}
	for rows.Next() {
		study := raw{}
		if err := rows.Scan(&study.StudyUuid, &study.TagGroup, &study.TagElement, &study.Value); err != nil {
			return SourceMySQLv1StudyType{}, err
		}
		if result.StudyUuid == "" {
			result.StudyUuid = study.StudyUuid
		}
		tag := utils.Dec2Hex(study.TagGroup, study.TagElement)
		result.Tags[tag] = study.Value
		if _, ok := intTags[tag]; ok {
			intValue, err := strconv.Atoi(study.Value)
			if err == nil {
				result.Tags[tag] = int64(intValue)
			}
		}
	}
	return result, nil
}

func (s *StudyStruct) GetStudyByStudyUuid(uuid string) (SourceMySQLv1StudyType, error) {
	q := `SELECT 
resourse.publicId as StudyUuid,
study.tagGroup as tagGroup,
study.tagElement as tagElement,
study.value as value
FROM Resources resourse
left join (select * from MainDicomTags) study on study.id = resourse.internalId
where publicId = ?;`
	rows, err := s.client.Query(q, uuid)
	if err != nil {
		return SourceMySQLv1StudyType{}, err
	}
	defer rows.Close()
	var result SourceMySQLv1StudyType = SourceMySQLv1StudyType{
		StudyUuid: "",
		Tags:      map[string]interface{}{},
	}
	for rows.Next() {
		study := raw{}
		if err := rows.Scan(&study.StudyUuid, &study.TagGroup, &study.TagElement, &study.Value); err != nil {
			return SourceMySQLv1StudyType{}, err
		}
		if study.StudyUuid == "" {
			result.StudyUuid = study.StudyUuid
		}
		tag := utils.Dec2Hex(study.TagGroup, study.TagElement)
		result.Tags[tag] = study.Value
		if _, ok := intTags[tag]; ok {
			intValue, err := strconv.Atoi(study.Value)
			if err == nil {
				result.Tags[tag] = int64(intValue)
			}
		}
	}
	return result, nil
}
