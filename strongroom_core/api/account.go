package api

import (
	"crypto/rsa"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shunr/strongroom_core/crypto"
	"github.com/shunr/strongroom_core/util"
	"io/ioutil"
)

const SECRET_KEY_LEN int = 26

type StrongroomAccount struct {
	Id                  uuid.UUID
	Username            string
	SecretKey           string
	MasterUnlockSalt    []byte
	AuthenticationSalt  []byte
	PublicKey           rsa.PublicKey
	EncryptedPrivateKey []byte
	PrivateKeyNonce     []byte
	VaultKeys           map[uuid.UUID]VaultKey
}

func CreateAccount(username string, password string) StrongroomAccount {

	secret_key := util.GenerateSecretKey(username, SECRET_KEY_LEN)

	// Salt for unlock key
	master_unlock_salt := util.GenerateSalt(32)

	// Salt for authentication key
	authentication_salt := util.GenerateSalt(32)

	// Generate key to decrypt
	master_unlock_key := util.DeriveKeyFromMasterPasswordAndSecretKey(
		username,
		password,
		secret_key,
		master_unlock_salt)
	/*
		// Generate x for SRP
		srp_x := util.DeriveKeyFromMasterPasswordAndSecretKey(
			username,
			password,
			secret_key,
			authentication_salt,
			32)*/

	private_key, public_key := util.GenerateAsymmetricKeyPair()

	// Encrypt private key with MUK
	nonce := crypto.RandBytes(12)
	private_key_json, _ := json.Marshal(private_key)
	encrypted_private_key, _ := crypto.AESGCMEncrypt(private_key_json, master_unlock_key, nonce)

	account_id, _ := uuid.NewRandom()
	account := StrongroomAccount{
		Id:                  account_id,
		Username:            username,
		SecretKey:           secret_key,
		MasterUnlockSalt:    master_unlock_salt,
		AuthenticationSalt:  authentication_salt,
		PublicKey:           *public_key,
		EncryptedPrivateKey: encrypted_private_key,
		PrivateKeyNonce:     nonce,
		VaultKeys:           map[uuid.UUID]VaultKey{},
	}

	return account
}

func ImportAccountFromFile(file_path string) (StrongroomAccount, error) {
	file_data, err := ioutil.ReadFile(file_path)
	if err != nil {
		return StrongroomAccount{}, err
	}
	var account StrongroomAccount
	err = json.Unmarshal(file_data, &account)
	if err != nil {
		return StrongroomAccount{}, err
	}
	return account, nil
}

func (account *StrongroomAccount) ExportToFile(file_path string) error {
	file_data, _ := json.MarshalIndent(account, "", "  ")
	return ioutil.WriteFile(file_path, file_data, 0644)
}

func (account *StrongroomAccount) GetPrivateKey(password string) (*rsa.PrivateKey, error) {
	muk := util.DeriveKeyFromMasterPasswordAndSecretKey(account.Username, password, account.SecretKey, account.MasterUnlockSalt)
	private_json, err := crypto.AESGCMDecrypt(account.EncryptedPrivateKey, muk, account.PrivateKeyNonce)
	if err != nil {
		return nil, err
	}
	var private_key rsa.PrivateKey
	json.Unmarshal(private_json, &private_key)
	return &private_key, nil
}

func (account *StrongroomAccount) AddVault(name string) error {
	vault := CreateVault(name)
	key := util.GenerateSymmetricKey()
	enc_key, err := crypto.EncryptRSAOAEP(key, []byte("vault_key"), &account.PublicKey)
	vault_key := VaultKey{VaultId: vault.Id, EncryptedKey: enc_key}
	// TODO: Consider checking to not overwrite existing vault key, throw error instead
	account.VaultKeys[vault.Id] = vault_key
	return err
}
