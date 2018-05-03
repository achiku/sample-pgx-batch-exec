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

func testNewDB(t *testing.T) *sql.DB {
	c, err := NewConfig("./config.toml")
	if err != nil {
		t.Fatal(err)
	}
	db, err := NewDB(c)
	if err != nil {
		t.Fatal(err)
	}
	return db.DB
}
