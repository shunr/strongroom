package api

import (
	"encoding/json"
	"github.com/shunr/strongroom_core/crypto"
	"github.com/shunr/strongroom_core/util"
)

const SECRET_KEY_LEN int = 26

type StrongroomAccount struct {
	Id                      string
	Username                string
	SecretKey               string
	MasterUnlockSalt        []byte
	AuthenticationSalt      []byte
	PublicKeyJson           []byte
	EncryptedPrivateKeyJson []byte
	PrivateKeyNonce         []byte
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
	public_key_json, _ := json.Marshal(public_key)
	encrypted_private_key := crypto.AESGCMEncrypt(private_key_json, master_unlock_key, nonce)

	account := StrongroomAccount{
		Id:                      "123",
		Username:                username,
		SecretKey:               secret_key,
		MasterUnlockSalt:        master_unlock_salt,
		AuthenticationSalt:      authentication_salt,
		PublicKeyJson:           public_key_json,
		EncryptedPrivateKeyJson: encrypted_private_key,
		PrivateKeyNonce:         nonce,
	}

	return account
}
