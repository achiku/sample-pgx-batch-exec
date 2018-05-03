package main

import (
	"database/sql"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
)

// Queryer database/sql compatible query interface
type Queryer interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

// Txer database/sql transaction interface
type Txer interface {
	Queryer
	Commit() error
	Rollback() error
}

// DBer database/sql
type DBer interface {
	Queryer
	Begin() (*sql.Tx, error)
	Close() error
	Ping() error
}

// DB rdbms connection
type DB struct {
	*sql.DB
}

// NewDB create DB
func NewDB(cfg *Config) (*DB, error) {
	dbCfg := &stdlib.DriverConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     cfg.DB.Host,
			User:     cfg.DB.User,
			Password: cfg.DB.Password,
			Database: cfg.DB.Name,
			Port:     cfg.DB.Port,
		},
	}
	stdlib.RegisterDriverConfig(dbCfg)
	db, err := sql.Open("pgx", dbCfg.ConnectionString(""))
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	// db.SetConnMaxLifetime(time.Second * 10)
	return &DB{db}, nil
}
