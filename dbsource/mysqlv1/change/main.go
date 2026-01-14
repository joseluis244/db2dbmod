package change

import (
	"database/sql"
)

type SourceMySQLv1ChangesType struct {
	Seq          int64  `json:"seq"`
	ChangeType   int    `json:"changeType"`
	InternalId   int64  `json:"internalId"`
	ResourceType string `json:"resourceType"`
	Date         string `json:"date"`
}

type ChangeStruct struct {
	client *sql.DB
}

func New(client *sql.DB) *ChangeStruct {
	return &ChangeStruct{
		client: client,
	}
}

func (m *ChangeStruct) LastChange() (int64, error) {
	q := `SELECT value FROM GlobalIntegers;`
	var value int64
	if err := m.client.QueryRow(q).Scan(&value); err != nil {
		return 0, err
	}
	return value, nil
}

func (m *ChangeStruct) ChangesRange(from int64, to int64) ([]SourceMySQLv1ChangesType, error) {
	q := `SELECT * FROM Changes where (changeType<10 and changeType!=3) and seq>=? and seq<=?;`
	rows, err := m.client.Query(q, from, to)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	changes := []SourceMySQLv1ChangesType{}
	for rows.Next() {
		change := SourceMySQLv1ChangesType{}
		if err := rows.Scan(&change.Seq, &change.ChangeType, &change.InternalId, &change.ResourceType, &change.Date); err != nil {
			return nil, err
		}
		changes = append(changes, change)
	}
	return changes, nil
}
