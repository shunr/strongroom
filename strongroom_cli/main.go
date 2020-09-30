package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/shunr/strongroom_core/client"
	. "github.com/shunr/strongroom_core/client"
)

const LOCAL_STORE_FILE string = "/tmp/strongroom_store.json"

func main() {

	reader := bufio.NewReader(os.Stdin)
	store, err := client.NewLocalStore(LOCAL_STORE_FILE)
	if err != nil {
		panic(err.Error())
	}
	client, err := NewClient(store)
	if err != nil {
		panic(err.Error())
	}

	var sess *Session
	for true {
		fmt.Print("Enter Command: ")
		command, _ := reader.ReadString('\n')
		switch strings.TrimSpace(command) {
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
		case "list_accounts":
			accounts := client.Accounts()
			for k, v := range accounts {
				fmt.Println(v.Username, k)
			}
			break
		case "quit":
			return
		case "login":
			fmt.Print("Id: ")
			id, _ := reader.ReadString('\n')
			fmt.Print("Password: ")
			password, err := reader.ReadString('\n')
			if err != nil {
				panic(err.Error())
			}
			acc_id, _ := uuid.Parse(strings.TrimSpace(id))
			password = strings.TrimSpace(password)

			account := client.Accounts()[acc_id]
			sess, err = client.NewSession(account, password)

			if err != nil {
				panic(err.Error())
			}
			break
		case "list_vaults":
			if sess == nil {
				fmt.Println("Must login before checking vaults")
				break
			}
			for k, _ := range sess.CurrentAccount.VaultKeys {
				fmt.Println("Vault id: ", k)
			}
			break
		case "add_vault":
			if sess == nil {
				fmt.Println("Must login before checking vaults")
				break
			}

			fmt.Print("Vault Name: ")
			vault, _ := reader.ReadString('\n')
			client.AddVault(sess, vault)
			break
		default:
			fmt.Println("Help")
			fmt.Println("command: create, usage: ", "create an account")
			fmt.Println("command: login, usage: ", "to the account")
			fmt.Println("command: list_accounts, usage: ", "list all accounts")
			fmt.Println("command: list_vaults, usage: ", "list all vaults")
			fmt.Println("command: add_vault, usage: ", "add vault to your account")
			fmt.Println("command: quit, usage: ", "exit the program")
		}
	}
}
