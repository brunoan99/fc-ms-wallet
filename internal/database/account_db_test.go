package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com/brunoan99/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type AccountDbTestSuite struct {
	suite.Suite
	db        *sql.DB
	AccountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDbTestSuite) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE, updated_at DATE)")
	db.Exec("CREATE TABLE accounts (id VARCHAR(255), client_id VARCHAR(255), balance FLOAT64, created_at DATE)")
	s.AccountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John Doe", "john@doe.com")
}

func (s *AccountDbTestSuite) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
	s.db.Exec("DROP TABLE accounts")
}

func TestAccountDbTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDbTestSuite))
}

func (s *AccountDbTestSuite) TestSave() {
	account, _ := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDbTestSuite) TestFindByID() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at, updated_at) VALUES (?,?,?,?,?)",
		s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt, s.client.UpdatedAt,
	)
	account, _ := entity.NewAccount(s.client)
	err := s.AccountDB.Save(account)
	s.Nil(err)
	accountDB, err := s.AccountDB.FindByID(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountDB.ID)
	s.Equal(account.Client.ID, accountDB.Client.ID)
	s.Equal(account.Balance, accountDB.Balance)
	s.WithinDuration(account.CreatedAt, accountDB.CreatedAt, 1*time.Millisecond)
	s.Equal(account.Client.Name, accountDB.Client.Name)
	s.Equal(account.Client.Email, accountDB.Client.Email)
	s.WithinDuration(account.Client.CreatedAt, accountDB.Client.CreatedAt, 1*time.Millisecond)
	s.WithinDuration(account.Client.UpdatedAt, accountDB.Client.UpdatedAt, 1*time.Millisecond)

}
