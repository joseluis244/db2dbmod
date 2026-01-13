package mysqlv1

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1/change"
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1/serie"
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1/study"
)

type MySQLv1 struct {
	client *sql.DB
	Change *change.ChangeStruct
	Study  *study.StudyStruct
	Serie  *serie.SerieStruct
}

func New() *MySQLv1 {
	return &MySQLv1{
		client: nil,
		Change: nil,
		Study:  nil,
		Serie:  nil,
	}
}

func (m *MySQLv1) Connect(dsn string) error {
	client, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err := client.Ping(); err != nil {
		return err
	}
	m.client = client
	m.Study = study.New(client)
	m.Change = change.New(client)
	m.Serie = serie.New(client)
	return nil
}

func (m *MySQLv1) Disconnect() error {
	if m.client == nil {
		return nil
	}
	return m.client.Close()
}
