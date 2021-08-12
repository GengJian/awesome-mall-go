package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// InitMysql ...
func InitMysql(fp string) error {
	db, err := sql.Open("mysql", fp)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
