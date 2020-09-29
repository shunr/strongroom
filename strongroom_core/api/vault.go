package api

import "github.com/google/uuid"

type VaultKey struct {
	VaultId      uuid.UUID
	EncryptedKey []byte
}

type Vault struct {
	Id       uuid.UUID
	Name     string
	Items    []VaultItem
	Metadata map[uuid.UUID]VaultItemMetadata
}

type VaultItemMetadata struct {
	Name        string
	Description string
}

type VaultItem struct {
	Id            uuid.UUID
	EncryptedData []byte
}

func CreateVault(name string) Vault {
	vault_id, _ := uuid.NewRandom()
	return Vault{vault_id, name, nil, nil}
}
