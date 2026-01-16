package instance

import (
	"database/sql"

	"github.com/joseluis244/db2dbmod/utils"
)

type SourceMySQLv1InstanceType struct {
	StudyUuid    string                 `json:"StudyUuid"`
	SerieUuid    string                 `json:"SerieUuid"`
	InstanceUuid string                 `json:"InstanceUuid"`
	Hash         string                 `json:"Hash"`
	Size         int64                  `json:"Size"`
	AE           string                 `json:"AE"`
	Tags         map[string]interface{} `json:"Tags"`
}

type raw struct {
	StudyUuid    string `json:"StudyUuid"`
	SerieUuid    string `json:"SerieUuid"`
	InstanceUuid string `json:"InstanceUuid"`
	Hash         string `json:"Hash"`
	Size         int64  `json:"Size"`
	AE           string `json:"AE"`
	TagGroup     int    `json:"TagGroup"`
	TagElement   int    `json:"TagElement"`
	Value        string `json:"Value"`
}

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

func (i *InstanceStruct) GetInstanceById(id int64) (SourceMySQLv1InstanceType, error) {
	q := `SELECT 
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
left join(SELECT * FROM Resources where resourceType = 2) resourseSerie on resourseSerie.internalId = resourseInstance.parentId
left join(SELECT * FROM Resources where resourceType = 1) resourseStudy on resourseStudy.internalId = resourseSerie.parentId
left join(SELECT * FROM MainDicomTags) InstanceTags on InstanceTags.id = resourseInstance.internalId
left join (SELECT * FROM AttachedFiles where fileType=1) InstanceFile on  InstanceFile.id = resourseInstance.internalId
where resourseInstance.internalId=?;`
	rows, err := i.client.Query(q, id)
	if err != nil {
		return SourceMySQLv1InstanceType{}, err
	}
	defer rows.Close()
	var result SourceMySQLv1InstanceType = SourceMySQLv1InstanceType{
		StudyUuid:    "",
		SerieUuid:    "",
		InstanceUuid: "",
		Hash:         "",
		Size:         0,
		AE:           "",
		Tags:         map[string]interface{}{},
	}
	for rows.Next() {
		instance := raw{}
		if err := rows.Scan(&instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.Hash, &instance.Size, &instance.AE, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return SourceMySQLv1InstanceType{}, err
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

func (i *InstanceStruct) GetInstanceByInstanceUuid(uuid string) (SourceMySQLv1InstanceType, error) {
	q := `SELECT 
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
left join(SELECT * FROM Resources where resourceType = 2) resourseSerie on resourseSerie.internalId = resourseInstance.parentId
left join(SELECT * FROM Resources where resourceType = 1) resourseStudy on resourseStudy.internalId = resourseSerie.parentId
left join(SELECT * FROM MainDicomTags) InstanceTags on InstanceTags.id = resourseInstance.internalId
left join (SELECT * FROM AttachedFiles where fileType=1) InstanceFile on  InstanceFile.id = resourseInstance.internalId
where resourseInstance.publicId=?;`
	rows, err := i.client.Query(q, uuid)
	if err != nil {
		return SourceMySQLv1InstanceType{}, err
	}
	defer rows.Close()
	var result SourceMySQLv1InstanceType = SourceMySQLv1InstanceType{
		StudyUuid:    "",
		SerieUuid:    "",
		InstanceUuid: "",
		Hash:         "",
		Size:         0,
		AE:           "",
		Tags:         map[string]interface{}{},
	}
	for rows.Next() {
		instance := raw{}
		if err := rows.Scan(&instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.Hash, &instance.Size, &instance.AE, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return SourceMySQLv1InstanceType{}, err
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

func (i *InstanceStruct) GetInstanceBySerieUuid(uuid string) ([]SourceMySQLv1InstanceType, error) {
	q := `SELECT 
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
left join(SELECT * FROM Resources where resourceType = 2) resourseSerie on resourseSerie.internalId = resourseInstance.parentId
left join(SELECT * FROM Resources where resourceType = 1) resourseStudy on resourseStudy.internalId = resourseSerie.parentId
left join(SELECT * FROM MainDicomTags) InstanceTags on InstanceTags.id = resourseInstance.internalId
left join (SELECT * FROM AttachedFiles where fileType=1) InstanceFile on  InstanceFile.id = resourseInstance.internalId
where resourseSerie.publicId=?
ORDER BY resourseInstance.publicId;`
	rows, err := i.client.Query(q, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []SourceMySQLv1InstanceType
	var currentInstance SourceMySQLv1InstanceType = SourceMySQLv1InstanceType{
		StudyUuid:    "",
		SerieUuid:    "",
		InstanceUuid: "",
		Hash:         "",
		Size:         0,
		AE:           "",
		Tags:         map[string]interface{}{},
	}
	for rows.Next() {
		instance := raw{}
		if err := rows.Scan(&instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.Hash, &instance.Size, &instance.AE, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return nil, err
		}
		if currentInstance.InstanceUuid != instance.InstanceUuid {
			result = append(result, SourceMySQLv1InstanceType{
				StudyUuid:    instance.StudyUuid,
				SerieUuid:    instance.SerieUuid,
				InstanceUuid: instance.InstanceUuid,
				Hash:         instance.Hash,
				Size:         instance.Size,
				AE:           instance.AE,
				Tags:         map[string]interface{}{},
			})
			result = append(result, currentInstance)
		}
		tag := utils.Dec2Hex(instance.TagGroup, instance.TagElement)
		result[len(result)-1].Tags[tag] = utils.TagIntConverter(intTags, tag, instance.Value)
	}
	return result, nil
}

func (i *InstanceStruct) GetInstanceByStudyUuid(uuid string) ([]SourceMySQLv1InstanceType, error) {
	q := `SELECT 
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
left join(SELECT * FROM Resources where resourceType = 2) resourseSerie on resourseSerie.internalId = resourseInstance.parentId
left join(SELECT * FROM Resources where resourceType = 1) resourseStudy on resourseStudy.internalId = resourseSerie.parentId
left join(SELECT * FROM MainDicomTags) InstanceTags on InstanceTags.id = resourseInstance.internalId
left join (SELECT * FROM AttachedFiles where fileType=1) InstanceFile on  InstanceFile.id = resourseInstance.internalId
where resourseStudy.publicId=?
ORDER BY resourseInstance.publicId;`
	rows, err := i.client.Query(q, uuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []SourceMySQLv1InstanceType
	var currentInstance SourceMySQLv1InstanceType = SourceMySQLv1InstanceType{
		StudyUuid:    "",
		SerieUuid:    "",
		InstanceUuid: "",
		Hash:         "",
		Size:         0,
		AE:           "",
		Tags:         map[string]interface{}{},
	}
	for rows.Next() {
		instance := raw{}
		if err := rows.Scan(&instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.Hash, &instance.Size, &instance.AE, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return nil, err
		}
		if instance.InstanceUuid != currentInstance.InstanceUuid {
			currentInstance = SourceMySQLv1InstanceType{
				StudyUuid:    instance.StudyUuid,
				SerieUuid:    instance.SerieUuid,
				InstanceUuid: instance.InstanceUuid,
				Hash:         instance.Hash,
				Size:         instance.Size,
				AE:           instance.AE,
				Tags:         map[string]interface{}{},
			}
			result = append(result, currentInstance)
		}
		tag := utils.Dec2Hex(instance.TagGroup, instance.TagElement)
		result[len(result)-1].Tags[tag] = utils.TagIntConverter(intTags, tag, instance.Value)
	}
	return result, nil
}

func (i *InstanceStruct) GetInstanceByChangeRange(from int64, to int64) ([]SourceMySQLv1InstanceType, error) {
	q := `SELECT 
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
order by StudyResourse.publicId,SeriesResourse.publicId;`
	rows, err := i.client.Query(q, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []SourceMySQLv1InstanceType
	var currentInstance SourceMySQLv1InstanceType = SourceMySQLv1InstanceType{
		StudyUuid:    "",
		SerieUuid:    "",
		InstanceUuid: "",
		Hash:         "",
		Size:         0,
		AE:           "",
		Tags:         map[string]interface{}{},
	}
	for rows.Next() {
		instance := raw{}
		if err := rows.Scan(&instance.StudyUuid, &instance.SerieUuid, &instance.InstanceUuid, &instance.Hash, &instance.Size, &instance.AE, &instance.TagGroup, &instance.TagElement, &instance.Value); err != nil {
			return nil, err
		}
		if instance.InstanceUuid != currentInstance.InstanceUuid {
			currentInstance = SourceMySQLv1InstanceType{
				StudyUuid:    instance.StudyUuid,
				SerieUuid:    instance.SerieUuid,
				InstanceUuid: instance.InstanceUuid,
				Hash:         instance.Hash,
				Size:         instance.Size,
				AE:           instance.AE,
				Tags:         map[string]interface{}{},
			}
			result = append(result, currentInstance)
		}
		tag := utils.Dec2Hex(instance.TagGroup, instance.TagElement)
		result[len(result)-1].Tags[tag] = utils.TagIntConverter(intTags, tag, instance.Value)
	}
	return result, nil
}
