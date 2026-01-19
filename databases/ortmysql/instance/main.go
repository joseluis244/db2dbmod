package instance

import (
	"database/sql"

	"github.com/joseluis244/db2dbmod/databases/ortmysql/models"
	"github.com/joseluis244/db2dbmod/utils"
)

var intTags = map[string]bool{
	"0020,0013": true,
}

type InstanceStruct struct {
	client *sql.DB
}

func New(client *sql.DB) *InstanceStruct {
	return &InstanceStruct{
		client: client,
	}
}

func (i *InstanceStruct) GetInstanceById(id int64) (models.OrtMySQLv1InstanceType, error) {
	q := `SELECT 
resourseInstance.internalId as Id,
resourseStudy.publicId as StudyUuid,
resourseSerie.publicId as SerieUuid,
resourseInstance.publicId as InstanceUuid,
InstanceFile.uuid as FileUuid,
InstanceFile.uncompressedSize as FileSize,
InstanceFile.uncompressedHash as FileHash,
InstanceTags.tagGroup as tagGroup,
InstanceTags.tagElement as tagElement,
InstanceTags.value as value
FROM Resources resourseInstance
left join Resources resourseSerie on resourseSerie.internalId = resourseInstance.parentId AND resourseSerie.resourceType = 2
left join Resources resourseStudy on resourseStudy.internalId = resourseSerie.parentId AND resourseStudy.resourceType = 1
left join MainDicomTags InstanceTags on InstanceTags.id = resourseInstance.internalId
left join AttachedFiles InstanceFile on InstanceFile.id = resourseInstance.internalId AND InstanceFile.fileType = 1
where resourseInstance.internalId=?;`
	rows, err := i.client.Query(q, id)
	if err != nil {
		return models.OrtMySQLv1InstanceType{}, err
	}
	defer rows.Close()
	var result models.OrtMySQLv1InstanceType = models.NewOrtMySQLv1InstanceType(0, "", "", "", "", "", 0, "", map[string]interface{}{})
	for rows.Next() {
		instance := models.OrtInstanceRaw{}
		if err := rows.Scan(&instance.Id, &instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.FileUuid, &instance.Size, &instance.Hash, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return models.OrtMySQLv1InstanceType{}, err
		}
		if result.StudyUuid == "" || result.SerieUuid == "" || result.InstanceUuid == "" || result.Hash == "" || result.Size == 0 {
			result.StudyUuid = instance.StudyUuid
			result.SerieUuid = instance.SerieUuid
			result.InstanceUuid = instance.InstanceUuid
			result.Hash = instance.Hash
			result.Size = instance.Size
		}
		tag := utils.Dec2Hex(instance.TagGroup, instance.TagElement)
		result.Tags[tag] = utils.TagIntConverter(intTags, tag, instance.Value)
	}
	return result, nil
}

func (i *InstanceStruct) GetInstanceByInstanceUuid(uuid string) (models.OrtMySQLv1InstanceType, error) {
	q := `SELECT 
resourseInstance.internalId as Id,
resourseStudy.publicId as StudyUuid,
resourseSerie.publicId as SerieUuid,
resourseInstance.publicId as InstanceUuid,
InstanceFile.uuid as FileUuid,
InstanceFile.uncompressedSize as FileSize,
InstanceFile.uncompressedHash as FileHash,
InstanceTags.tagGroup as tagGroup,
InstanceTags.tagElement as tagElement,
InstanceTags.value as value
FROM Resources resourseInstance
left join Resources resourseSerie on resourseSerie.internalId = resourseInstance.parentId AND resourseSerie.resourceType = 2
left join Resources resourseStudy on resourseStudy.internalId = resourseSerie.parentId AND resourseStudy.resourceType = 1
left join MainDicomTags InstanceTags on InstanceTags.id = resourseInstance.internalId
left join AttachedFiles InstanceFile on InstanceFile.id = resourseInstance.internalId AND InstanceFile.fileType = 1
where resourseInstance.publicId=?;`
	rows, err := i.client.Query(q, uuid)
	if err != nil {
		return models.OrtMySQLv1InstanceType{}, err
	}
	defer rows.Close()
	var result models.OrtMySQLv1InstanceType = models.NewOrtMySQLv1InstanceType(0, "", "", "", "", "", 0, "", map[string]interface{}{})
	for rows.Next() {
		instance := models.OrtInstanceRaw{}
		if err := rows.Scan(&instance.Id, &instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.FileUuid, &instance.Size, &instance.Hash, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return models.OrtMySQLv1InstanceType{}, err
		}
		if result.StudyUuid == "" || result.SerieUuid == "" || result.InstanceUuid == "" || result.Hash == "" || result.Size == 0 {
			result.StudyUuid = instance.StudyUuid
			result.SerieUuid = instance.SerieUuid
			result.InstanceUuid = instance.InstanceUuid
			result.Hash = instance.Hash
			result.Size = instance.Size
		}
		tag := utils.Dec2Hex(instance.TagGroup, instance.TagElement)
		result.Tags[tag] = utils.TagIntConverter(intTags, tag, instance.Value)
	}
	return result, nil
}

func (i *InstanceStruct) GetInstanceBySerieUuid(uuid string) ([]models.OrtMySQLv1InstanceType, error) {
	q := `SELECT 
resourseInstance.internalId as Id,
resourseStudy.publicId as StudyUuid,
resourseSerie.publicId as SerieUuid,
resourseInstance.publicId as InstanceUuid,
InstanceFile.uuid as FileUuid,
InstanceFile.uncompressedSize as FileSize,
InstanceFile.uncompressedHash as FileHash,
InstanceTags.tagGroup as tagGroup,
InstanceTags.tagElement as tagElement,
InstanceTags.value as value
FROM Resources resourseInstance
left join Resources resourseSerie on resourseSerie.internalId = resourseInstance.parentId AND resourseSerie.resourceType = 2
left join Resources resourseStudy on resourseStudy.internalId = resourseSerie.parentId AND resourseStudy.resourceType = 1
left join MainDicomTags InstanceTags on InstanceTags.id = resourseInstance.internalId
left join AttachedFiles InstanceFile on InstanceFile.id = resourseInstance.internalId AND InstanceFile.fileType = 1
where resourseSerie.publicId=?
ORDER BY resourseInstance.publicId;`
	rows, err := i.client.Query(q, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []models.OrtMySQLv1InstanceType
	var currentInstance models.OrtMySQLv1InstanceType = models.NewOrtMySQLv1InstanceType(0, "", "", "", "", "", 0, "", map[string]interface{}{})
	for rows.Next() {
		instance := models.OrtInstanceRaw{}
		if err := rows.Scan(&instance.Id, &instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.FileUuid, &instance.Size, &instance.Hash, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return nil, err
		}
		if currentInstance.InstanceUuid != instance.InstanceUuid {
			currentInstance = models.NewOrtMySQLv1InstanceType(instance.Id, instance.StudyUuid, instance.SerieUuid, instance.InstanceUuid, instance.FileUuid, instance.Hash, instance.Size, "", map[string]interface{}{})
			result = append(result, currentInstance)
		}
		tag := utils.Dec2Hex(instance.TagGroup, instance.TagElement)
		result[len(result)-1].Tags[tag] = utils.TagIntConverter(intTags, tag, instance.Value)
	}
	return result, nil
}

func (i *InstanceStruct) GetInstanceByStudyUuid(uuid string) ([]models.OrtMySQLv1InstanceType, error) {
	q := `SELECT 
resourseInstance.internalId as Id,
resourseStudy.publicId as StudyUuid,
resourseSerie.publicId as SerieUuid,
resourseInstance.publicId as InstanceUuid,
InstanceFile.uuid as FileUuid,
InstanceFile.uncompressedSize as FileSize,
InstanceFile.uncompressedHash as FileHash,
InstanceTags.tagGroup as tagGroup,
InstanceTags.tagElement as tagElement,
InstanceTags.value as value
FROM Resources resourseInstance
left join Resources resourseSerie on resourseSerie.internalId = resourseInstance.parentId AND resourseSerie.resourceType = 2
left join Resources resourseStudy on resourseStudy.internalId = resourseSerie.parentId AND resourseStudy.resourceType = 1
left join MainDicomTags InstanceTags on InstanceTags.id = resourseInstance.internalId
left join AttachedFiles InstanceFile on InstanceFile.id = resourseInstance.internalId AND InstanceFile.fileType = 1
where resourseStudy.publicId=?
ORDER BY resourseInstance.publicId;`
	rows, err := i.client.Query(q, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []models.OrtMySQLv1InstanceType
	var currentInstance models.OrtMySQLv1InstanceType = models.NewOrtMySQLv1InstanceType(0, "", "", "", "", "", 0, "", map[string]interface{}{})
	for rows.Next() {
		instance := models.OrtInstanceRaw{}
		if err := rows.Scan(&instance.Id, &instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.FileUuid, &instance.Size, &instance.Hash, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return nil, err
		}
		if instance.InstanceUuid != currentInstance.InstanceUuid {
			currentInstance = models.NewOrtMySQLv1InstanceType(instance.Id, instance.StudyUuid, instance.SerieUuid, instance.InstanceUuid, instance.FileUuid, instance.Hash, instance.Size, "", map[string]interface{}{})
			result = append(result, currentInstance)
		}
		tag := utils.Dec2Hex(instance.TagGroup, instance.TagElement)
		result[len(result)-1].Tags[tag] = utils.TagIntConverter(intTags, tag, instance.Value)
	}
	return result, nil
}

func (i *InstanceStruct) GetInstanceByChangeRange(from int64, to int64) ([]models.OrtMySQLv1InstanceType, error) {
	q := `SELECT 
InstanceResourse.internalId as Id,
StudyResourse.publicId as StudyUuid,
SeriesResourse.publicId as SerieUuid,
InstanceResourse.publicId as InstanceUuid,
InstanceFile.uuid as FileUuid,
InstanceFile.uncompressedSize as FileSize,
InstanceFile.uncompressedHash as FileHash,
InstanceTags.tagGroup as tagGroup,
InstanceTags.tagElement as tagElement,
InstanceTags.value as value
FROM Changes InstanceChange
left join Resources InstanceResourse on InstanceResourse.internalId=InstanceChange.internalId and InstanceResourse.resourceType=3
left join Resources SeriesResourse on InstanceResourse.parentId=SeriesResourse.internalId and SeriesResourse.resourceType=2
left join Resources StudyResourse on StudyResourse.internalId=SeriesResourse.parentId and StudyResourse.resourceType=1
left join MainDicomTags InstanceTags on InstanceResourse.internalId=InstanceTags.id
left join AttachedFiles InstanceFile on  InstanceFile.id = InstanceResourse.internalId and InstanceFile.fileType=1
where InstanceChange.changeType=2 and (InstanceChange.seq>=? and InstanceChange.seq<=?)
order by StudyResourse.publicId,SeriesResourse.publicId,InstanceResourse.publicId;`
	rows, err := i.client.Query(q, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []models.OrtMySQLv1InstanceType
	var currentInstance models.OrtMySQLv1InstanceType = models.NewOrtMySQLv1InstanceType(0, "", "", "", "", "", 0, "", map[string]interface{}{})
	for rows.Next() {
		instance := models.OrtInstanceRaw{}
		if err := rows.Scan(&instance.Id, &instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.FileUuid, &instance.Size, &instance.Hash, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return nil, err
		}
		if instance.InstanceUuid != currentInstance.InstanceUuid {
			currentInstance = models.NewOrtMySQLv1InstanceType(instance.Id, instance.StudyUuid, instance.SerieUuid, instance.InstanceUuid, instance.FileUuid, instance.Hash, instance.Size, "", map[string]interface{}{})
			result = append(result, currentInstance)
		}
		tag := utils.Dec2Hex(instance.TagGroup, instance.TagElement)
		result[len(result)-1].Tags[tag] = utils.TagIntConverter(intTags, tag, instance.Value)
	}
	return result, nil
}
