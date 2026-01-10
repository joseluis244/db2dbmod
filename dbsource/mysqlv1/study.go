package mysqlv1

import "github.com/joseluis244/db2dbmod/models"

func (m *MySQLv1) GetStudyById(id string) ([]models.SourceMySQLv1StudyType, error) {
	q := `SELECT 
* 
FROM symphony.Resources resourse
left join (select * from symphony.MainDicomTags) study on study.id = resourse.internalId
where internalId = ?;`
	rows, err := m.client.Query(q, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	studies := []models.SourceMySQLv1StudyType{}
	for rows.Next() {
		study := models.SourceMySQLv1StudyType{}
		if err := rows.Scan(&study.StudyUuid, &study.Tags, &study.CreatedAt, &study.UpdatedAt, &study.BuildTime); err != nil {
			return nil, err
		}
		studies = append(studies, study)
	}
	return studies, nil
}
