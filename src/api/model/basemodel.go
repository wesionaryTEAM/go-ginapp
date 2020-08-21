package model

import (
	"fmt"
	"go-ginapp/auth"

	//github.com is ...
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

//Server is ...
type Server struct {
	DB *gorm.DB
}

var (
	//Server now implements the modelInterface, so can define its methods

	//Model is ...
	Model modelInterface = &Server{}
)

type modelInterface interface {
	//db initialization
	Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error)

	//user methods
	ValidateEmail(string) error
	CreateUser(*User) (*User, error)
	GetUserByEmail(string) (*User, error)

	//todo methods:
	CreateTodo(*Todo) (*Todo, error)

	//auth methods:
	FetchAuth(*auth.AuthDetails) (*Auth, error)
	DeleteAuth(*auth.AuthDetails) error
	CreateAuth(uint64) (*Auth, error)
}

//Initialize is ...
func (s *Server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	var err error
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s/ sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	s.DB, err = gorm.Open(Dbdriver, DBURL)
	if err != nil {
		return nil, err
	}
	s.DB.Debug().AutoMigrate(
		&User{},
		&Auth{},
		&Todo{},
	)
	return s.DB, nil
}
