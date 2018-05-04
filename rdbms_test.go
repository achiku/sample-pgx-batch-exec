package main

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx"
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

func TestBatchExec(t *testing.T) {
	db, cleanup := testNewDB(t)
	defer cleanup()

	conn, err := stdlib.AcquireConn(db)
	if err != nil {
		t.Fatal(err)
	}
	defer stdlib.ReleaseConn(db, conn)

	_, err = conn.Prepare("q1", "insert into bulk_insert (val, num, created_at) values ($1, $2, $3)")
	if err != nil {
		t.Fatal(err)
	}

	b := conn.BeginBatch()
	defer b.Close()

	b.Queue("q1",
		[]interface{}{"v", 10, time.Now()},
		[]pgtype.OID{pgtype.VarcharOID, pgtype.NumericOID, pgtype.TimestamptzOID}, nil,
	)
	b.Queue("q1",
		[]interface{}{"v", 10.11, time.Now()},
		[]pgtype.OID{pgtype.VarcharOID, pgtype.NumericOID, pgtype.TimestamptzOID}, nil,
	)

	if err := b.Send(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	res, err := b.ExecResults()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", res)

	var cnt int
	if err := db.QueryRow(`select count(*) from bulk_insert`).Scan(&cnt); err != nil {
		t.Fatal(err)
	}
	t.Logf("%d", cnt)
}

func TestCopyExec(t *testing.T) {
	db, cleanup := testNewDB(t)
	defer cleanup()

	conn, err := stdlib.AcquireConn(db)
	if err != nil {
		t.Fatal(err)
	}
	defer stdlib.ReleaseConn(db, conn)

	cols := []string{"val", "num", "created_at"}
	rows := [][]interface{}{
		{"val1", 10.12, time.Now()},
		{"val2", 10.12, time.Now()},
		{"val2", 10.12, time.Now()},
		{"val2", 10.12, time.Now()},
		{"val2", 10.12, time.Now()},
		{"val2", 10.12, time.Now()},
		{"val2", 10.12, time.Now()},
		{"val2", 10.12, time.Now()},
	}
	cnt, err := conn.CopyFrom(pgx.Identifier{"bulk_insert"}, cols, pgx.CopyFromRows(rows))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%d", cnt)
}
