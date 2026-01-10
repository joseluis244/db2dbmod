package mysqlv1

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLv1 struct {
	client *sql.DB
}

func New() *MySQLv1 {
	return &MySQLv1{
		client: nil,
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
	return nil
}

func (m *MySQLv1) Disconnect() error {
	if m.client == nil {
		return nil
	}
	return m.client.Close()
}
