package repository_test

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SuiteUser struct {
	suite.Suite
	GormDB     *gorm.DB
	SqlDB      *sql.DB
	mock       sqlmock.Sqlmock
	repository repository.UserRepository
}

func SetupSuiteUser() *SuiteUser {
	s := &SuiteUser{}
	var err error

	s.SqlDB, s.mock, err = sqlmock.New()
	if err != nil {
		log.Panicf("Failed to open mock sql db, got error: %v", err)
	}

	if s.SqlDB == nil {
		log.Panic("mock db is null")
	}

	if s.mock == nil {
		log.Panic("sqlmock is null")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 s.SqlDB,
		PreferSimpleProtocol: true,
	})

	s.GormDB, err = gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Panicf("Failed to open gorm v2 db, got error: %v", err)
	}
	if s.GormDB == nil {
		log.Panic("gorm db is null")
	}

	s.repository = repository.NewUserRepository(s.GormDB)
	return s
}

func TestCreateNewUser_NoError(t *testing.T) {
	s := SetupSuiteUser()

	rows := sqlmock.NewRows([]string{"name", "email", "password"}).
		AddRow("test1", "test1@shopee.com", "pass1")

	testEnt := entity.User{
		Password: "pass1",
		Email:    "test1@shopee.com",
		Name:     "test1",
	}

	q := `
		INSERT INTO "users"
	`
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(q)).
		WithArgs(testEnt.Password, testEnt.Email, testEnt.Name).
		WillReturnRows(rows)
	s.mock.ExpectCommit()

	res, rowsAffected, err := s.repository.CreateUser(&testEnt)

	assert.NoError(t, err)
	assert.Equal(t, 1, rowsAffected)
	assert.NotNil(t, res)

}

func TestCreateNewUser_Error(t *testing.T) {
	s := SetupSuiteUser()

	rows := sqlmock.NewRows([]string{"name", "email", "password"}).
		AddRow("test1", "test1@shopee.com", "pass1")

	testEnt := entity.User{
		Password: "",
		Email:    "test1@shopee.com",
	}

	q := `
		INSERT INTO "users"
	`
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(q)).
		WithArgs(testEnt.Email, testEnt.Password).
		WillReturnRows(rows)
	s.mock.ExpectCommit()

	_, rowsAffected, err := s.repository.CreateUser(&testEnt)

	assert.Error(t, err)
	assert.Equal(t, 0, rowsAffected)

}

func TestGetUserDetails_NoError(t *testing.T) {
	s := SetupSuiteUser()

	row := sqlmock.NewRows([]string{"name", "email", "password"}).
		AddRow("test1", "test1@shopee.com", "pass1")

	s.mock.ExpectQuery("SELECT (.+)").WithArgs(1).WillReturnRows(row)

	user, err := s.repository.GetUserDetails(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestGetUserByEmail_NoError(t *testing.T) {
	s := SetupSuiteUser()

	row := sqlmock.NewRows([]string{"name", "email", "password"}).
		AddRow("test1", "test1@shopee.com", "pass1")

	s.mock.ExpectQuery("SELECT (.+)").WithArgs("test1@shopee.com").WillReturnRows(row)

	user, rowsAffected, err := s.repository.GetUserByEmail("test1@shopee.com")

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, rowsAffected)
}

func TestGetUserById_NoError(t *testing.T) {
	s := SetupSuiteUser()

	row := sqlmock.NewRows([]string{"wallet_id", "name", "email", "password"}).
		AddRow(100001, "test1", "test1@shopee.com", "pass1")

	s.mock.ExpectQuery("SELECT (.+)").WithArgs(100001).WillReturnRows(row)

	user, rowsAffected, err := s.repository.GetUserByID(100001)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, 1, rowsAffected)
}

func TestUpdateBalanceById_Error(t *testing.T) {
	s := SetupSuiteUser()

	s.mock.ExpectExec("UPDATE (.+)").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 0))

	_, err := s.repository.UpdateBalanceByWalletID(2, 2)

	assert.NotNil(t, err)

}
