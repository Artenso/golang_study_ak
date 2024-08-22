package main

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

func (u *User) TableName() string {
	return "users"
}

type Tabler interface {
	TableName() string
}

type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{
		db:           db,
		sqlGenerator: sqlGenerator,
	}
}

func (m *Migrator) Migrate(models ...Tabler) error {
	tx, err := m.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to init transaction: %s", err.Error())
	}
	defer tx.Commit()

	for _, model := range models {
		query := m.sqlGenerator.CreateTableSQL(model)
		_, err := tx.Exec(query)
		if err != nil {
			return fmt.Errorf("failed to add query to transaction: %s", err.Error())
		}
	}
	return nil
}

type SQLiteGenerator struct{}

// CreateInsertSQL implements SQLGenerator.
func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	user := reflect.ValueOf(model).Elem()

	columns := make([]string, 0, user.NumField())
	values := make([]string, 0, user.NumField())

	for i := 0; i < user.NumField(); i++ {
		column := user.Type().Field(i).Tag.Get(`db_field`)
		columns = append(columns, column)

		data := user.Field(i).Interface()
		value := fmt.Sprintf(`%v`, data)
		if _, ok := data.(string); ok {
			value = fmt.Sprintf(`'%v'`, value)
		}
		values = append(values, value)
	}

	return fmt.Sprintf(
		`INSERT INTO %s (%s) VALUES (%s);`,
		model.TableName(),
		strings.Join(columns, ", "),
		strings.Join(values, ", "),
	)

}

// CreateTableSQL implements SQLGenerator.
func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	user := reflect.ValueOf(table).Elem()

	columns := make([]string, 0, user.NumField())

	for i := 0; i < user.NumField(); i++ {
		column := fmt.Sprintf(`%s %s`,
			user.Type().Field(i).Tag.Get(`db_field`),
			user.Type().Field(i).Tag.Get(`db_type`),
		)
		columns = append(columns, column)
	}

	return fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (%s);`,
		table.TableName(),
		strings.Join(columns, ", "),
	)
}

// Основная функция
func main() {
	// Подключение к SQLite БД
	db, err := sql.Open("sqlite3", "course2/2.oop/3.oop_abstraction/task2.2.3.4/mydb.db")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}

	// Создание мигратора с использованием вашего SQLGenerator
	YourSQLGeneratorInstance := &SQLiteGenerator{}
	migrator := NewMigrator(db, YourSQLGeneratorInstance)

	// Миграция таблицы User
	if err := migrator.Migrate(&User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
