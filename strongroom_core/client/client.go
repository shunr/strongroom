package client

import (
	"errors"
	"github.com/google/uuid"
	"github.com/shunr/strongroom_core/crypto"
	. "github.com/shunr/strongroom_core/model"
	"github.com/shunr/strongroom_core/util"
)

const LOCAL_STORE_FILE string = "/tmp/strongroom_store.json"

type Session struct {
	CurrentAccount *StrongroomAccount
	PrivateKey     *crypto.PrivateKey
}

type StrongroomClient struct {
	local_store *LocalStore
}

func NewClient(store *LocalStore) (*StrongroomClient, error) {
	client := StrongroomClient{
		local_store: store,
	}
	return &client, nil
}

func (client *StrongroomClient) NewSession(account *StrongroomAccount, password string) (*Session, error) {
	private, err := account.GetPrivateKey(password)
	if err != nil {
		return nil, errors.New("Incorrect master password")
	}
	session := Session{CurrentAccount: account, PrivateKey: private}
	return &session, nil
}

func (client *StrongroomClient) Accounts() map[uuid.UUID]*StrongroomAccount {
	return client.local_store.Accounts
}

func (client *StrongroomClient) CreateAccount(username string, password string) error {
	account := NewAccount(username, password)
	err := client.local_store.AddAccount(account)
	if err != nil {
		return err
	}
	defer client.local_store.Save()
	return nil
}

func (client *StrongroomClient) Vaults() map[uuid.UUID]*EncryptedVault {
	return client.local_store.EncryptedVaults
}

func (client *StrongroomClient) AddVault(session *Session, name string) {
	vault := NewVault(name)
	raw_vault_key := util.GenerateSymmetricKey()
	public_key := &session.CurrentAccount.PublicKey

	// Add encrypted vault to local store
	client.local_store.EncryptedVaults[vault.Id] = EncryptVault(&vault, raw_vault_key)

	// Encrypt and store encrypted vault key with account
	enc_vault_key, err := crypto.EncryptRSAOAEP(raw_vault_key, []byte("vault_key_"+name), public_key)
	if err != nil {
		panic(err.Error())
	}
	vault_key := VaultKey{VaultId: vault.Id, EncryptedKey: enc_vault_key}
	session.CurrentAccount.VaultKeys[vault.Id] = vault_key
	defer client.local_store.Save()
}
