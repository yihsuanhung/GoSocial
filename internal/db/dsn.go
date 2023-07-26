package db

import "fmt"

type DSN struct {
	Dialect   string            `default:"mysql"`
	Username  string            `default:"root"`
	Password  string            `default:"root"`
	Net       string            // Network type
	Host      string            `default:"127.0.0.1"`
	Port      string            `default:"3306"`
	Dbname    string            // Database name
	Params    map[string]string // Connection parameters
	Migration bool              `default:"false"`
}

func (d *DSN) GetMysqlDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", d.Username, d.Password, d.Host, d.Port, d.Dbname)
}
