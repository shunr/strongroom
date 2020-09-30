package model

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/shunr/strongroom_core/crypto"
)

type VaultKey struct {
	VaultId      uuid.UUID
	EncryptedKey []byte
}

type EncryptedVault struct {
	Id            uuid.UUID
	EncryptedData []byte
	Nonce         []byte
}

type Vault struct {
	Id       uuid.UUID
	Name     string
	Metadata map[uuid.UUID]VaultItemMetadata
	Items    map[uuid.UUID]VaultItem
}

type VaultItemMetadata struct {
	Name        string
	Description string
}

type VaultItem struct {
	Id            uuid.UUID
	EncryptedData []byte
	Nonce         []byte
}

func NewVault(name string) Vault {
	vault_id, _ := uuid.NewRandom()
	vault := Vault{vault_id, name, make(map[uuid.UUID]VaultItemMetadata), make(map[uuid.UUID]VaultItem)}
	return vault
}

func (vault *Vault) AddVaultItem(metadata VaultItemMetadata, data []byte, vault_key []byte) error {
	nonce := crypto.RandNonce()
	enc_data, err := crypto.EncryptAESGCM(data, vault_key, nonce)
	if err != nil {
		return err
	}
	item_id, _ := uuid.NewRandom()
	vault_item := VaultItem{Id: item_id, EncryptedData: enc_data, Nonce: nonce}
	vault.Items[item_id] = vault_item
	vault.Metadata[item_id] = metadata
	return nil
}

func EncryptVault(vault *Vault, vault_key []byte) *EncryptedVault {
	bytes, err := json.Marshal(vault)
	if err != nil {
		panic(err.Error())
	}
	nonce := crypto.RandNonce()
	enc_data, err := crypto.EncryptAESGCM(bytes, vault_key, nonce)
	if err != nil {
		panic(err.Error())
	}
	encrypted_vault := EncryptedVault{
		Id:            vault.Id,
		EncryptedData: enc_data,
		Nonce:         nonce,
	}
	return &encrypted_vault
}

func DecryptVault(enc_vault *EncryptedVault, vault_key []byte) *Vault {
	data, err := crypto.DecryptAESGCM(enc_vault.EncryptedData, vault_key, enc_vault.Nonce)
	if err != nil {
		panic(err.Error())
	}
	var vault Vault
	err = json.Unmarshal(data, &vault)
	if err != nil {
		panic(err.Error())
	}
	return &vault
}
