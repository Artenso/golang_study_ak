package main

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit"
	_ "github.com/mattn/go-sqlite3"
)

func TestUser_TableName(t *testing.T) {
	tests := []struct {
		name string
		u    *User
		want string
	}{
		{
			name: "#1",
			u:    &User{},
			want: "users",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.TableName(); got != tt.want {
				t.Errorf("User.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMigrator(t *testing.T) {
	type args struct {
		db           *sql.DB
		sqlGenerator SQLGenerator
	}
	tests := []struct {
		name string
		args args
		want *Migrator
	}{
		{
			name: "#1",
			args: args{
				db:           &sql.DB{},
				sqlGenerator: &SQLiteGenerator{},
			},
			want: &Migrator{
				db:           &sql.DB{},
				sqlGenerator: &SQLiteGenerator{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMigrator(tt.args.db, tt.args.sqlGenerator); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMigrator() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	type args struct {
		model Tabler
	}

	testUser := &User{
		ID:        int(gofakeit.Uint64()),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}

	tests := []struct {
		name string
		s    *SQLiteGenerator
		args args
		want string
	}{
		{
			name: "#1",
			s:    &SQLiteGenerator{},
			args: args{
				model: testUser,
			},
			want: fmt.Sprintf(
				`INSERT INTO users (id, first_name, last_name, email) VALUES (%d, '%s', '%s', '%s');`,
				testUser.ID,
				testUser.FirstName,
				testUser.LastName,
				testUser.Email,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CreateInsertSQL(tt.args.model); got != tt.want {
				t.Errorf("SQLiteGenerator.CreateInsertSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	type args struct {
		table Tabler
	}
	tests := []struct {
		name string
		s    *SQLiteGenerator
		args args
		want string
	}{
		{
			name: "#1",
			s:    &SQLiteGenerator{},
			args: args{
				table: &User{},
			},
			want: `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, first_name VARCHAR(100), last_name VARCHAR(100), email VARCHAR(100) UNIQUE);`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.CreateTableSQL(tt.args.table); got != tt.want {
				t.Errorf("SQLiteGenerator.CreateTableSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}
