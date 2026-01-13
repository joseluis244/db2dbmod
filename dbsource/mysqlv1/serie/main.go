package serie

import (
	"database/sql"

	"github.com/joseluis244/db2dbmod/models"
	"github.com/joseluis244/db2dbmod/utils"
)

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

func (s *SerieStruct) GetSerieById(id string) (models.SourceMySQLv1SerieType, error) {
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
		return models.SourceMySQLv1SerieType{}, err
	}
	defer rows.Close()
	var result models.SourceMySQLv1SerieType = models.SourceMySQLv1SerieType{
		StudyUuid: "",
		SerieUuid: "",
		Tags:      map[string]interface{}{},
	}
	for rows.Next() {
		serie := raw{}
		if err := rows.Scan(&serie.StudyUuid, &serie.SerieUuid, &serie.TagGroup, &serie.TagElement, &serie.Value); err != nil {
			return models.SourceMySQLv1SerieType{}, err
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

func (s *SerieStruct) GetSerieBySerieUuid(uuid string) (models.SourceMySQLv1SerieType, error) {
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
		return models.SourceMySQLv1SerieType{}, err
	}
	defer rows.Close()
	var result models.SourceMySQLv1SerieType = models.SourceMySQLv1SerieType{
		StudyUuid: "",
		SerieUuid: "",
		Tags:      map[string]interface{}{},
	}
	for rows.Next() {
		serie := raw{}
		if err := rows.Scan(&serie.StudyUuid, &serie.SerieUuid, &serie.TagGroup, &serie.TagElement, &serie.Value); err != nil {
			return models.SourceMySQLv1SerieType{}, err
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

func (s *SerieStruct) GetSerieByStudyUuid(uuid string) ([]models.SourceMySQLv1SerieType, error) {
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
	var result []models.SourceMySQLv1SerieType
	var currentSerie models.SourceMySQLv1SerieType = models.SourceMySQLv1SerieType{
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
			if currentSerie.SerieUuid != "" {
				result = append(result, currentSerie)
			}
			currentSerie = models.SourceMySQLv1SerieType{
				StudyUuid: serie.StudyUuid,
				SerieUuid: serie.SerieUuid,
				Tags:      map[string]interface{}{},
			}
		}
		tag := utils.Dec2Hex(serie.TagGroup, serie.TagElement)
		currentSerie.Tags[tag] = serie.Value
	}
	if currentSerie.SerieUuid != "" {
		result = append(result, currentSerie)
	}
	return result, nil
}
