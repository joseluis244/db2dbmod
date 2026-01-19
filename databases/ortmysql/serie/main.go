package serie

import (
	"database/sql"

	"github.com/joseluis244/db2dbmod/databases/ortmysql/models"
	"github.com/joseluis244/db2dbmod/utils"
)

type SerieStruct struct {
	client *sql.DB
}

func New(client *sql.DB) *SerieStruct {
	return &SerieStruct{
		client: client,
	}
}

func (s *SerieStruct) GetSerieById(id int64) (models.OrtMySQLv1SerieType, error) {
	q := `SELECT 
resourseSerie.internalId as Id,
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
		return models.OrtMySQLv1SerieType{}, err
	}
	defer rows.Close()
	var result models.OrtMySQLv1SerieType = models.NewOrtMySQLv1SerieType(0, "", "", map[string]interface{}{})
	for rows.Next() {
		serie := models.OrtSerieRaw{}
		if err := rows.Scan(&serie.Id, &serie.StudyUuid, &serie.SerieUuid, &serie.TagGroup, &serie.TagElement, &serie.Value); err != nil {
			return models.OrtMySQLv1SerieType{}, err
		}
		if result.StudyUuid == "" || result.SerieUuid == "" {
			result.Id = serie.Id
			result.StudyUuid = serie.StudyUuid
			result.SerieUuid = serie.SerieUuid
		}
		tag := utils.Dec2Hex(serie.TagGroup, serie.TagElement)
		result.Tags[tag] = serie.Value
	}
	return result, nil
}

func (s *SerieStruct) GetSerieBySerieUuid(uuid string) (models.OrtMySQLv1SerieType, error) {
	q := `SELECT 
resourseSerie.internalId as Id,
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
		return models.OrtMySQLv1SerieType{}, err
	}
	defer rows.Close()
	var result models.OrtMySQLv1SerieType = models.NewOrtMySQLv1SerieType(0, "", "", map[string]interface{}{})
	for rows.Next() {
		serie := models.OrtSerieRaw{}
		if err := rows.Scan(&serie.Id, &serie.StudyUuid, &serie.SerieUuid, &serie.TagGroup, &serie.TagElement, &serie.Value); err != nil {
			return models.OrtMySQLv1SerieType{}, err
		}
		if result.StudyUuid == "" || result.SerieUuid == "" {
			result.Id = serie.Id
			result.StudyUuid = serie.StudyUuid
			result.SerieUuid = serie.SerieUuid
		}
		tag := utils.Dec2Hex(serie.TagGroup, serie.TagElement)
		result.Tags[tag] = serie.Value
	}
	return result, nil
}

func (s *SerieStruct) GetSerieByStudyUuid(uuid string) ([]models.OrtMySQLv1SerieType, error) {
	q := `SELECT 
resourseSerie.internalId as Id,
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
	var result []models.OrtMySQLv1SerieType
	var currentSerie models.OrtMySQLv1SerieType = models.NewOrtMySQLv1SerieType(0, "", "", map[string]interface{}{})
	for rows.Next() {
		serie := models.OrtSerieRaw{}
		if err := rows.Scan(&serie.Id, &serie.StudyUuid, &serie.SerieUuid, &serie.TagGroup, &serie.TagElement, &serie.Value); err != nil {
			return nil, err
		}
		if serie.SerieUuid != currentSerie.SerieUuid {
			currentSerie = models.NewOrtMySQLv1SerieType(serie.Id, serie.StudyUuid, serie.SerieUuid, map[string]interface{}{})
			result = append(result, currentSerie)
		}
		tag := utils.Dec2Hex(serie.TagGroup, serie.TagElement)
		result[len(result)-1].Tags[tag] = serie.Value
	}
	return result, nil
}

func (s *SerieStruct) GetSerieByChangeRange(from int64, to int64) ([]models.OrtMySQLv1SerieType, error) {
	q := `SELECT 
SeriesResourse.internalId as Id,
StudyResourse.publicId as StudyUuid,
SeriesResourse.publicId as SerieUuid,
SeriesTags.tagGroup as tagGroup,
SeriesTags.tagElement as tagElement,
SeriesTags.value as value
FROM Changes SeriesChange
left join Resources SeriesResourse on SeriesResourse.internalId=SeriesChange.internalId and SeriesResourse.resourceType=2
left join Resources StudyResourse on StudyResourse.internalId=SeriesResourse.parentId and StudyResourse.resourceType=1
left join MainDicomTags SeriesTags on SeriesResourse.internalId=SeriesTags.id
where SeriesChange.changeType=4 and (SeriesChange.seq>=? and SeriesChange.seq<=?)
order by StudyResourse.publicId;`
	rows, err := s.client.Query(q, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []models.OrtMySQLv1SerieType
	var currentSerie models.OrtMySQLv1SerieType = models.NewOrtMySQLv1SerieType(0, "", "", map[string]interface{}{})
	for rows.Next() {
		serie := models.OrtSerieRaw{}
		if err := rows.Scan(&serie.Id, &serie.StudyUuid, &serie.SerieUuid, &serie.TagGroup, &serie.TagElement, &serie.Value); err != nil {
			return nil, err
		}
		if serie.SerieUuid != currentSerie.SerieUuid {
			currentSerie = models.NewOrtMySQLv1SerieType(serie.Id, serie.StudyUuid, serie.SerieUuid, map[string]interface{}{})
			result = append(result, currentSerie)
		}
		tag := utils.Dec2Hex(serie.TagGroup, serie.TagElement)
		result[len(result)-1].Tags[tag] = serie.Value
	}
	return result, nil
}
