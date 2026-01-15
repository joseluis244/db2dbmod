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

func intconverter(tag string, value string) any {
	if _, ok := intTags[tag]; ok {
		intValue, err := strconv.Atoi(value)
		if err == nil {
			return int64(intValue)
		}
	}
	return value
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
		result.Tags[tag] = intconverter(tag, study.Value)
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
		result.Tags[tag] = intconverter(tag, study.Value)
	}
	return result, nil
}

func (s *StudyStruct) GetStudyByChangeRange(from int64, to int64) ([]SourceMySQLv1StudyType, error) {
	q := `SELECT 
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
	var result []SourceMySQLv1StudyType
	var currentStudy SourceMySQLv1StudyType = SourceMySQLv1StudyType{
		StudyUuid: "",
		Tags:      map[string]interface{}{},
	}
	for rows.Next() {
		study := raw{}
		if err := rows.Scan(&study.StudyUuid, &study.TagGroup, &study.TagElement, &study.Value); err != nil {
			return nil, err
		}
		if study.StudyUuid != currentStudy.StudyUuid {
			currentStudy = SourceMySQLv1StudyType{
				StudyUuid: study.StudyUuid,
				Tags:      map[string]interface{}{},
			}
			result = append(result, currentStudy)
		}
		tag := utils.Dec2Hex(study.TagGroup, study.TagElement)
		result[len(result)-1].Tags[tag] = intconverter(tag, study.Value)
	}
	return result, nil
}
