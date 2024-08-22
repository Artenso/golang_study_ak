package main

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
	panic("unimplemented")
}

type SQLiteGenerator struct{}

// CreateInsertSQL implements SQLGenerator.
func (s *SQLiteGenerator) CreateInsertSQL(model Tabler) string {
	panic("unimplemented")
}

// CreateTableSQL implements SQLGenerator.
func (s *SQLiteGenerator) CreateTableSQL(table Tabler) string {
	panic("unimplemented")
}

type GoFakeitGenerator struct{}

// GenerateFakeUser implements FakeDataGenerator.
func (g *GoFakeitGenerator) GenerateFakeUser() User {
	panic("unimplemented")
}

// func main() {
// 	sqlGenerator := &SQLiteGenerator{}
// 	fakeDataGenerator := &GoFakeitGenerator{}
// 	user := User{}
// 	sql := sqlGenerator.CreateTableSQL(&user)
// 	fmt.Println(sql)
// 	for i := 0; i < 34; i++ {
// 		fakeUser := fakeDataGenerator.GenerateFakeUser()
// 		query := sqlGenerator.CreateInsertSQL(&fakeUser)
// 		fmt.Println(query)
// 	}
// }
