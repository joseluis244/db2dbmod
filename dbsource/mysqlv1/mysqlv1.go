package mysqlv1

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1/change"
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1/instance"
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1/serie"
	"github.com/joseluis244/db2dbmod/dbsource/mysqlv1/study"
)

type MySQLv1 struct {
	client   *sql.DB
	Change   *change.ChangeStruct
	Study    *study.StudyStruct
	Serie    *serie.SerieStruct
	Instance *instance.InstanceStruct
}

func New() *MySQLv1 {
	return &MySQLv1{
		client:   nil,
		Change:   nil,
		Study:    nil,
		Serie:    nil,
		Instance: nil,
	}
}

func (m *MySQLv1) Connect(dsn string) error {
	fmt.Println("Connecting to MySQL", dsn)
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
	m.Instance = instance.New(client)
	fmt.Println("Connected to MySQL")
	return nil
}

func (m *MySQLv1) Disconnect() error {
	if m.client == nil {
		return nil
	}
	return m.client.Close()
}
