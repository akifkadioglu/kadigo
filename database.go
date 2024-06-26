package kadigo

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

var db *gorm.DB

type PostgreSQL struct {
	DSN string
}
type MySQL struct {
	DSN string
}
type SQLite struct {
	Path string
}
type DBType interface {
	PostgreSQL | MySQL | SQLite
	connect() error
}

// StartDatabase starts database connection
func StartDatabase[T DBType](database T, entities ...any) error {
	err := database.connect()
	if err != nil {
		return err
	}
	err = db.AutoMigrate(entities...)
	return err
}

// DBManager returns database connection
func DBManager() *gorm.DB {
	return db
}

func (d MySQL) connect() error {
	var err error
	db, err = gorm.Open(mysql.Open(d.DSN), &gorm.Config{})
	return err
}
func (d PostgreSQL) connect() error {
	var err error
	db, err = gorm.Open(postgres.Open(d.DSN), &gorm.Config{})
	return err
}

func (d SQLite) connect() error {
	var err error
	db, err = gorm.Open(sqlite.Open(d.Path), &gorm.Config{})
	return err
}
