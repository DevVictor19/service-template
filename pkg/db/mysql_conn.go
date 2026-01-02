package db

import (
	"database/sql"
	"fmt"
	"log"
	"servicetemplate/pkg/env"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySqlDB(c *env.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4&loc=Local",
		c.MySQL.User,
		c.MySQL.Password,
		c.MySQL.Host,
		c.MySQL.Port,
		c.MySQL.DBname,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to open mysql connection: %v", err)
	}

	db.SetMaxOpenConns(c.MySQL.MaxOpenConns)
	db.SetMaxIdleConns(c.MySQL.MaxIdleConns)
	db.SetConnMaxLifetime(c.MySQL.ConnMaxLifetime)
	db.SetConnMaxIdleTime(c.MySQL.ConnMaxIdleTime)

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping mysql: %v", err)
	}

	return db
}
