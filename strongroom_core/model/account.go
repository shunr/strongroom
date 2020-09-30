package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shunr/strongroom_core/crypto"
	"github.com/shunr/strongroom_core/util"
)

const SECRET_KEY_LEN int = 26

type StrongroomAccount struct {
	Id                  uuid.UUID
	Username            string
	SecretKey           string
	MasterUnlockSalt    []byte
	AuthenticationSalt  []byte
	PublicKey           crypto.PublicKey
	EncryptedPrivateKey []byte
	PrivateKeyNonce     []byte
	VaultKeys           map[uuid.UUID]VaultKey
}

func NewAccount(username string, password string) *StrongroomAccount {

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
	nonce := crypto.RandNonce()
	private_key_json, _ := json.Marshal(private_key)
	encrypted_private_key, _ := crypto.EncryptAESGCM(private_key_json, master_unlock_key, nonce)

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

	return &account
}

func (account *StrongroomAccount) GetPrivateKey(password string) (*crypto.PrivateKey, error) {
	muk := util.DeriveKeyFromMasterPasswordAndSecretKey(account.Username, password, account.SecretKey, account.MasterUnlockSalt)
	private_json, err := crypto.DecryptAESGCM(account.EncryptedPrivateKey, muk, account.PrivateKeyNonce)
	if err != nil {
		return nil, err
	}
	var private_key crypto.PrivateKey
	json.Unmarshal(private_json, &private_key)
	return &private_key, nil
}
