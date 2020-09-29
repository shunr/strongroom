package main

import (
	"bufio"
	"fmt"
	"github.com/shunr/strongroom_core/api"
	"github.com/shunr/strongroom_core/crypto"
	"github.com/shunr/strongroom_core/util"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')
	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	account := api.CreateAccount(username, password)

	muk := util.DeriveKeyFromMasterPasswordAndSecretKey(username, password, account.SecretKey, account.MasterUnlockSalt)
	private := crypto.AESGCMDecrypt(account.EncryptedPrivateKeyJson, muk, account.PrivateKeyNonce)

	fmt.Println(string(private))
}
