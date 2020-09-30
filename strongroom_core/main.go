package main

import (
	"bufio"
	"fmt"
	"github.com/google/uuid"
	"github.com/shunr/strongroom_core/client"
	"os"
	"strings"
)

const LOCAL_STORE_FILE string = "/tmp/strongroom_store.json"

func main() {
	reader := bufio.NewReader(os.Stdin)
	cmd := os.Args[1]
	store, err := client.NewLocalStore(LOCAL_STORE_FILE)
	if err != nil {
		panic(err.Error())
	}
	client, err := client.NewClient(store)
	if err != nil {
		panic(err.Error())
	}
	switch cmd {
	case "create":
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		password = strings.TrimSpace(password)
		err := client.CreateAccount(username, password)
		if err != nil {
			panic(err.Error())
		}
		break
	case "load":
		accounts := client.Accounts()
		fmt.Println(accounts)
		fmt.Print("Id: ")
		id, _ := reader.ReadString('\n')
		fmt.Print("Password: ")
		password, err := reader.ReadString('\n')
		if err != nil {
			panic(err.Error())
		}
		acc_id, _ := uuid.Parse(strings.TrimSpace(id))
		password = strings.TrimSpace(password)
		account := accounts[acc_id]
		sess, err := client.NewSession(account, password)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(sess.CurrentAccount.Username)
		vid := client.AddVault(sess, "vault1")
		vault, err := client.OpenVault(sess, vid)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(vault)
		//vault.AddVaultItem("password_1", "My password", []byte("bruh123"))
		break
	default:
		fmt.Println(client.Accounts())
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		fmt.Print("Password: ")
		password, _ := reader.ReadString('\n')
		username = strings.TrimSpace(username)
		password = strings.TrimSpace(password)

	}
}
