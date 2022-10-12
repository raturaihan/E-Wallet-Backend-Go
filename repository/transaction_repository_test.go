package repository_test

import (
	"assignment-golang-backend/entity"
	"assignment-golang-backend/repository"
	"database/sql"
	"errors"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SuiteTrans struct {
	suite.Suite
	GormDB     *gorm.DB
	SqlDB      *sql.DB
	mock       sqlmock.Sqlmock
	repository repository.TransactionRepository
}

func SetupSuiteTrans() *SuiteTrans {
	s := &SuiteTrans{}
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

	s.repository = repository.NewTransactionRepository(s.GormDB)
	return s
}

func TestDoTransaction_Error(t *testing.T) {
	s := SetupSuiteTrans()

	rows := sqlmock.NewRows([]string{"wallet_id", "trans_type", "amount", "fund_id"}).
		AddRow(100001, "TOPUP", "100000", 1)

	testEnt := entity.Transaction{
		WalletID:  100001,
		TransType: "TOPUP",
		Amount:    100000,
		FundID:    1,
	}

	q := `
		INSERT INTO "transactions"
	`
	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(q)).
		WithArgs(testEnt.WalletID, testEnt.TransType, testEnt.Amount, testEnt.FundID).
		WillReturnRows(rows)
	s.mock.ExpectCommit()

	err := s.repository.DoTransaction(&testEnt)

	assert.Error(t, err)

}

func TestGetFundNameById_NoError(t *testing.T) {
	s := SetupSuiteTrans()

	row := sqlmock.NewRows([]string{"fund_id", "fund_name"}).
		AddRow(1, "bank transfer")

	s.mock.ExpectQuery("SELECT (.+)").WithArgs(1).WillReturnRows(row)

	fund, err := s.repository.GetFundNameById(1)

	assert.Nil(t, err)
	assert.NotNil(t, fund)
}

func TestGetFundNameById_Error(t *testing.T) {
	s := SetupSuiteTrans()

	s.mock.ExpectQuery("SELECT (.+)").WithArgs(1).WillReturnError(errors.New("id not found"))

	_, err := s.repository.GetFundNameById(1)

	assert.NotNil(t, err)
}

func TestGetAllTransaction_NoError(t *testing.T) {
	s := SetupSuiteTrans()

	rows := sqlmock.NewRows([]string{"wallet_id", "trans_type", "amount", "fund_id"}).
		AddRow(100001, "TOPUP", "100000", 1).
		AddRow(100001, "TOPUP", "50000", 3)

	q := `SELECT * FROM "transactions"`

	s.mock.ExpectQuery(regexp.QuoteMeta(q)).WithArgs(100001).WillReturnRows(rows)

	params := make(map[string]string)
	books, err := s.repository.GetAllTransactionById(100001, params)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(books))
}
