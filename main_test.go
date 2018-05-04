package main

import (
	"database/sql"
	"testing"
)

func testNewConfig(t *testing.T) *Config {
	c, err := NewConfig("./config.toml")
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func testNewDB(t *testing.T) (*sql.DB, func()) {
	c, err := NewConfig("./config.toml")
	if err != nil {
		t.Fatal(err)
	}
	db, err := NewDB(c)
	if err != nil {
		t.Fatal(err)
	}
	return db.DB, func() {
		db.Exec(`truncate bulk_insert`)
		db.Exec(`truncate bulk_update`)
	}
}
