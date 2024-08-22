package main

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

var _ SQLGenerator = &SQLiteGenerator{}
var _ FakeDataGenerator = &GoFakeitGenerator{}
var _ Tabler = &User{}

type Tabler interface {
	TableName() string
}

// Интерфейс для генерации SQL-запросов
type SQLGenerator interface {
	CreateTableSQL(table Tabler) string
	CreateInsertSQL(model Tabler) string
}

// Интерфейс для генерации фейковых данных
type FakeDataGenerator interface {
	GenerateFakeUser() User
}

// Определение структуры пользователя
type User struct {
	ID        int    `db_field:"id" db_type:"SERIAL PRIMARY KEY"`
	FirstName string `db_field:"first_name" db_type:"VARCHAR(100)"`
	LastName  string `db_field:"last_name" db_type:"VARCHAR(100)"`
	Email     string `db_field:"email" db_type:"VARCHAR(100) UNIQUE"`
}

// TableName implements Tabler.
func (u *User) TableName() string {
	return "users"
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

type GoFakeitGenerator struct{}

// GenerateFakeUser implements FakeDataGenerator.
func (g *GoFakeitGenerator) GenerateFakeUser() User {
	gofakeit.Seed(0)
	return User{
		ID:        int(gofakeit.Uint64()),
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
	}
}

// как проверить функции генерирующие случайные значения???
func GenerateUserInserts(count int) []string {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	queries := make([]string, 0, count)
	for i := 0; i < count; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		queries = append(queries, query)
		time.Sleep(time.Nanosecond)
	}
	return queries
}

func main() {
	sqlGenerator := &SQLiteGenerator{}
	fakeDataGenerator := &GoFakeitGenerator{}
	user := User{}
	sql := sqlGenerator.CreateTableSQL(&user)
	fmt.Println(sql)
	for i := 0; i < 34; i++ {
		fakeUser := fakeDataGenerator.GenerateFakeUser()
		query := sqlGenerator.CreateInsertSQL(&fakeUser)
		fmt.Println(query)
		time.Sleep(time.Nanosecond)
	}
	queries := GenerateUserInserts(34)
	for _, query := range queries {
		fmt.Println(query)
	}
}
