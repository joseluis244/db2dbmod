package ortmysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joseluis244/db2dbmod/databases/ortmysql/change"
	"github.com/joseluis244/db2dbmod/databases/ortmysql/instance"
	"github.com/joseluis244/db2dbmod/databases/ortmysql/serie"
	"github.com/joseluis244/db2dbmod/databases/ortmysql/study"
)

type OrtMySQL struct {
	client   *sql.DB
	Change   *change.ChangeStruct
	Study    *study.StudyStruct
	Serie    *serie.SerieStruct
	Instance *instance.InstanceStruct
}

func New() *OrtMySQL {
	return &OrtMySQL{
		client:   nil,
		Change:   nil,
		Study:    nil,
		Serie:    nil,
		Instance: nil,
	}
}

func (m *OrtMySQL) Connect(dsn string) error {
	fmt.Println("Connecting to MySQL", dsn)
	client, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	if err := client.Ping(); err != nil {
		return err
	}
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	m.client = client
	m.Study = study.New(client)
	m.Change = change.New(client)
	m.Serie = serie.New(client)
	m.Instance = instance.New(client)
	fmt.Println("Connected to MySQL")
	return nil
}

func (m *OrtMySQL) Disconnect() error {
	if m.client == nil {
		return nil
	}
	return m.client.Close()
}
