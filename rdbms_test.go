package main

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/stdlib"
)

func TestNewDB(t *testing.T) {
	c := testNewConfig(t)
	db, err := NewDB(c)
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
}

func TestNewPgx(t *testing.T) {
	db := testNewDB(t)
	conn, err := stdlib.AcquireConn(db)
	if err != nil {
		t.Fatal(err)
	}
	defer stdlib.ReleaseConn(db, conn)

	b := conn.BeginBatch()
	b.Queue("insert into bulk_insert (val, created_at) values ($1, $2)",
		[]interface{}{"v", time.Now()}, []pgtype.OID{pgtype.VarcharOID, pgtype.TimestamptzOID}, nil)
	b.Queue("insert into bulk_insert (val, created_at) values ($1, $2)",
		[]interface{}{"v", time.Now()}, []pgtype.OID{pgtype.VarcharOID, pgtype.TimestamptzOID}, nil)
	ctx := context.Background()
	b.Send(ctx, nil)
}
