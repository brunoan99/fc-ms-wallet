package database

import (
	"database/sql"
	"testing"
	"time"

	"github.com/brunoan99/fc-ms-wallet/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDbTestSuit struct {
	suite.Suite
	db       *sql.DB
	ClientDB *ClientDB
}

func (s *ClientDbTestSuit) SetupTest() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db
	db.Exec("CREATE TABLE clients (id VARCHAR(255), name VARCHAR(255), email VARCHAR(255), created_at DATE, updated_at DATE)")
	s.ClientDB = NewClientDB(s.db)
}

func (s *ClientDbTestSuit) TearDownTest() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE clients")
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDbTestSuit))
}

func (s *ClientDbTestSuit) TestSave() {
	client := &entity.Client{
		ID:        "1",
		Name:      "John Doe",
		Email:     "john@doe.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err := s.ClientDB.Save(client)
	s.Nil(err)
}

func (s *ClientDbTestSuit) TestGet() {
	client, _ := entity.NewClient("John Doe", "j@j.com")
	s.ClientDB.Save(client)

	clientDB, err := s.ClientDB.Get(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
	s.WithinDuration(client.CreatedAt, clientDB.CreatedAt, 1*time.Millisecond)
	s.WithinDuration(client.UpdatedAt, clientDB.UpdatedAt, 1*time.Millisecond)
}
