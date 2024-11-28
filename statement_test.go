//go:build go1.6
// +build go1.6

package sqlmock

import (
	"errors"
	"testing"

	"github.com/google/go-safeweb/safesql"
)

func TestExpectedPreparedStatementCloseError(t *testing.T) {
	conn, mock, err := New()
	if err != nil {
		t.Fatal("failed to open sqlmock database:", err)
	}

	mock.ExpectBegin()
	want := errors.New("STMT ERROR")
	mock.ExpectPrepare("SELECT").WillReturnCloseError(want)

	txn, err := conn.Begin()
	if err != nil {
		t.Fatal("unexpected error while opening transaction:", err)
	}

	stmt, err := txn.Prepare(safesql.New("SELECT"))
	if err != nil {
		t.Fatal("unexpected error while preparing a statement:", err)
	}

	if err := stmt.Close(); err != want {
		t.Fatalf("got = %v, want = %v", err, want)
	}
}
