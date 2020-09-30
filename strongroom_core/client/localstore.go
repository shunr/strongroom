package client

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	. "github.com/shunr/strongroom_core/model"
	"io/ioutil"
	"os"
)

type LocalStore struct {
	Accounts        map[uuid.UUID]*StrongroomAccount
	EncryptedVaults map[uuid.UUID]*EncryptedVault
	file_path       string
}

func fileExists(file_path string) bool {
	info, err := os.Stat(file_path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func NewLocalStore(file_path string) (*LocalStore, error) {
	var store LocalStore
	store.file_path = file_path
	if fileExists(file_path) {
		file_data, err := ioutil.ReadFile(file_path)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(file_data, &store)
		if err != nil {
			return nil, err
		}
	} else {
		store.Accounts = make(map[uuid.UUID]*StrongroomAccount)
		store.EncryptedVaults = make(map[uuid.UUID]*EncryptedVault)
		err := store.Save()
		if err != nil {
			return nil, err
		}
	}
	return &store, nil
}

func (store *LocalStore) AddAccount(account *StrongroomAccount) error {
	key := account.Id
	if _, exists := store.Accounts[key]; exists {
		return errors.New("Account with key " + key.String() + " already exists in store")
	}
	store.Accounts[key] = account
	return nil
}

func (store *LocalStore) Save() error {
	file_data, _ := json.MarshalIndent(store, "", "  ")
	return ioutil.WriteFile(store.file_path, file_data, 0644)
}
