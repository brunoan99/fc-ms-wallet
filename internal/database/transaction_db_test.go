package database

import (
	"database/sql"
	"testing"

	"github.com/brunoan99/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type TransactionDbTestSuite struct {
	suite.Suite
	db            *sql.DB
	client1       *entity.Client
	client2       *entity.Client
	accountFrom   *entity.Account
	accountTo     *entity.Account
	TransactionDB *TransactionDB
}

func (s *TransactionDbTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE, updated_at DATE)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance FLOAT64, created_at DATE)")
	db.Exec("CREATE TABLE transactions (id VARCHAR(255), from_id VARCHAR(255), to_id VARCHAR(255), amount FLOAT64, created_at DATE)")
	s.TransactionDB = NewTransactionDB(db)
	s.client1, _ = entity.NewClient("John Doe", "john@doe.com")
	s.client2, _ = entity.NewClient("Janie Doe", "janie@doe.com")
	s.accountFrom, _ = entity.NewAccount(s.client1)
	s.accountFrom.Balance = 1000
	s.accountTo, _ = entity.NewAccount(s.client2)
	s.accountTo.Balance = 1000
}

func (s *TransactionDbTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE transactions")
}

func TestTransactionDbTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionDbTestSuite))
}

func (s *TransactionDbTestSuite) TestCreate() {
	transaction, err := entity.NewTransaction(s.accountFrom, s.accountTo, 100)
	s.Nil(err)
	err = s.TransactionDB.Create(transaction)
	s.Nil(err)
}
