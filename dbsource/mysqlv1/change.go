package mysqlv1

import "github.com/joseluis244/db2dbmod/models"

func (m *MySQLv1) LastChange() (int64, error) {
	q := `SELECT value FROM GlobalIntegers;`
	var value int64
	if err := m.client.QueryRow(q).Scan(&value); err != nil {
		return 0, err
	}
	return value, nil
}

func (m *MySQLv1) ChangesRange(from int64, to int64) ([]models.SourceMySQLv1ChangesType, error) {
	q := `SELECT * FROM Changes where (changeType<10 and changeType!=3) and seq>=? and seq<=?;`
	rows, err := m.client.Query(q, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	changes := []models.SourceMySQLv1ChangesType{}
	for rows.Next() {
		change := models.SourceMySQLv1ChangesType{}
		if err := rows.Scan(&change.Seq, &change.ChangeType, &change.InternalId, &change.ResourceType, &change.Date); err != nil {
			return nil, err
		}
		changes = append(changes, change)
	}
	return changes, nil
}
