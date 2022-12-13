package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, err := NewAccount(client)
	assert.Nil(t, err)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account, err := NewAccount(nil)
	assert.Nil(t, account)
	assert.NotNil(t, err)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	assert.Equal(t, 0.0, account.Balance)
	account.Credit(100.0)
	assert.Equal(t, 100.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account, _ := NewAccount(client)
	account.Credit(100.0)
	assert.Equal(t, 100.0, account.Balance)
	account.Debit(50.0)
	assert.Equal(t, 50.0, account.Balance)
}
