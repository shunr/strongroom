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
	cmd := os.Args[1]
	switch cmd {
	case "create":
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		password = strings.TrimSpace(password)
		account := api.CreateAccount(username, password)
		err := account.ExportToFile(os.Args[2])
		if err != nil {
			panic(err.Error())
		}
		break
	case "load":
		account, err := api.ImportAccountFromFile(os.Args[2])
		if err != nil {
			panic(err.Error())
		}
		fmt.Print("Password: ")
		reader := bufio.NewReader(os.Stdin)
		password, _ := reader.ReadString('\n')
		muk := util.DeriveKeyFromMasterPasswordAndSecretKey(account.Username, password, account.SecretKey, account.MasterUnlockSalt)
		private, err := crypto.AESGCMDecrypt(account.EncryptedPrivateKey, muk, account.PrivateKeyNonce)
		if err != nil {
			fmt.Println("Wrong master password")
			os.Exit(-1)
		} else {
			fmt.Println(string(private))
		}

		break
	}

}
