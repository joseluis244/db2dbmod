package serie

import (
	"database/sql"

	"github.com/joseluis244/db2dbmod/utils"
)

type SourceMySQLv1SerieType struct {
	StudyUuid string                 `json:"StudyUuid"`
	SerieUuid string                 `json:"SerieUuid"`
	Tags      map[string]interface{} `json:"Tags"`
}

type raw struct {
	StudyUuid  string `json:"StudyUuid"`
	SerieUuid  string `json:"SerieUuid"`
	TagGroup   int    `json:"TagGroup"`
	TagElement int    `json:"TagElement"`
	Value      string `json:"Value"`
}

type SerieStruct struct {
	client *sql.DB
}

func New(client *sql.DB) *SerieStruct {
	return &SerieStruct{
		client: client,
	}
}

func (s *SerieStruct) GetSerieById(id int64) (SourceMySQLv1SerieType, error) {
	q := `SELECT 
resourseStudy.publicId as StudyUuid,
resourseSerie.publicId as SerieUuid,
SerieTags.tagGroup as tagGroup,
SerieTags.tagElement as tagElement,
SerieTags.value as value
FROM Resources resourseSerie
left join(SELECT * FROM Resources where resourceType = 1) resourseStudy on resourseStudy.internalId = resourseSerie.parentId
left join(SELECT * FROM MainDicomTags) SerieTags on SerieTags.id = resourseSerie.internalId
where resourseSerie.internalId=?;`
	rows, err := s.client.Query(q, id)
	if err != nil {
		return SourceMySQLv1SerieType{}, err
	}
	defer rows.Close()
	var result SourceMySQLv1SerieType = SourceMySQLv1SerieType{
		StudyUuid: "",
		SerieUuid: "",
		Tags:      map[string]interface{}{},
	}
	for rows.Next() {
		serie := raw{}
		if err := rows.Scan(&serie.StudyUuid, &serie.SerieUuid, &serie.TagGroup, &serie.TagElement, &serie.Value); err != nil {
			return SourceMySQLv1SerieType{}, err
		}
		if result.StudyUuid == "" || result.SerieUuid == "" {
			result.StudyUuid = serie.StudyUuid
			result.SerieUuid = serie.SerieUuid
		}
		tag := utils.Dec2Hex(serie.TagGroup, serie.TagElement)
		result.Tags[tag] = serie.Value
	}
	return result, nil
}

func (s *SerieStruct) GetSerieBySerieUuid(uuid string) (SourceMySQLv1SerieType, error) {
	q := `SELECT 
resourseStudy.publicId as StudyUuid,
resourseSerie.publicId as SerieUuid,
SerieTags.tagGroup as tagGroup,
SerieTags.tagElement as tagElement,
SerieTags.value as value
FROM Resources resourseSerie
left join(SELECT * FROM Resources where resourceType = 1) resourseStudy on resourseStudy.internalId = resourseSerie.parentId
left join(SELECT * FROM MainDicomTags) SerieTags on SerieTags.id = resourseSerie.internalId
where resourseSerie.publicId=?;`
	rows, err := s.client.Query(q, uuid)
	if err != nil {
		return SourceMySQLv1SerieType{}, err
	}
	defer rows.Close()
	var result SourceMySQLv1SerieType = SourceMySQLv1SerieType{
		StudyUuid: "",
		SerieUuid: "",
		Tags:      map[string]interface{}{},
	}
	for rows.Next() {
		serie := raw{}
		if err := rows.Scan(&serie.StudyUuid, &serie.SerieUuid, &serie.TagGroup, &serie.TagElement, &serie.Value); err != nil {
			return SourceMySQLv1SerieType{}, err
		}
		if result.StudyUuid == "" || result.SerieUuid == "" {
			result.StudyUuid = serie.StudyUuid
			result.SerieUuid = serie.SerieUuid
		}
		tag := utils.Dec2Hex(serie.TagGroup, serie.TagElement)
		result.Tags[tag] = serie.Value
	}
	return result, nil
}

func (s *SerieStruct) GetSerieByStudyUuid(uuid string) ([]SourceMySQLv1SerieType, error) {
	q := `SELECT 
resourseStudy.publicId as StudyUuid,
resourseSerie.publicId as SerieUuid,
SerieTags.tagGroup as tagGroup,
SerieTags.tagElement as tagElement,
SerieTags.value as value
FROM Resources resourseSerie
left join(SELECT * FROM Resources where resourceType = 1) resourseStudy on resourseStudy.internalId = resourseSerie.parentId
left join(SELECT * FROM MainDicomTags) SerieTags on SerieTags.id = resourseSerie.internalId
where resourseStudy.publicId=?
ORDER BY resourseSerie.publicId;`
	rows, err := s.client.Query(q, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []SourceMySQLv1SerieType
	var currentSerie SourceMySQLv1SerieType = SourceMySQLv1SerieType{
		StudyUuid: "",
		SerieUuid: "",
		Tags:      map[string]interface{}{},
	}
	for rows.Next() {
		serie := raw{}
		if err := rows.Scan(&serie.StudyUuid, &serie.SerieUuid, &serie.TagGroup, &serie.TagElement, &serie.Value); err != nil {
			return nil, err
		}
		if serie.SerieUuid != currentSerie.SerieUuid {
			currentSerie = SourceMySQLv1SerieType{
				StudyUuid: serie.StudyUuid,
				SerieUuid: serie.SerieUuid,
				Tags:      map[string]interface{}{},
			}
			result = append(result, currentSerie)
		}
		tag := utils.Dec2Hex(serie.TagGroup, serie.TagElement)
		result[len(result)-1].Tags[tag] = serie.Value
	}
	return result, nil
}
