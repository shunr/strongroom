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
			password, _ := reader.ReadString('\n')
			acc_id, err := uuid.Parse(strings.TrimSpace(id))
			if err != nil {
				fmt.Println("Incorrect id or password")
				break
			}

			password = strings.TrimSpace(password)
			account := client.Accounts()[acc_id]
			sess, err = client.NewSession(account, password)

			if err != nil {
				fmt.Println("Incorrect id or password")
				break
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
			client.AddVault(sess, strings.TrimSpace(vault))
			break
		case "open_vault":
			if sess == nil {
				fmt.Println("Must login before checking vaults")
				break
			}

			id, _ := reader.ReadString('\n')
			uuid, _ := uuid.Parse(strings.TrimSpace(id))
			vault, err := client.OpenVault(sess, uuid)

			if err != nil {
				fmt.Println("Cannot open vault: ", err.Error())
				break
			}

			// Id       uuid.UUID
			// Name     string
			// Items    map[uuid.UUID]VaultItem
			// Metadata map[uuid.UUID]VaultItemMetadata

			// type VaultItemMetadata struct {
			// 	Name        string
			// 	Description string
			// }

			// type VaultItem struct {
			// 	Id            uuid.UUID
			// 	EncryptedData []byte
			// 	Nonce         []byte
			// }

			fmt.Println("Vault Id: ", vault.Id)
			fmt.Println("Vault Name: ", vault.Name)

			for k, v := range vault.Items {
				fmt.Println(vault.Metadata[k].Name)
				fmt.Println(vault.Metadata[k].Description)
				fmt.Println(string(v.EncryptedData))
			}
			break
		default:
			fmt.Println("Help")
			fmt.Println("command: create, usage: ", "create an account")
			fmt.Println("command: login, usage: ", "to the account")
			fmt.Println("command: list_accounts, usage: ", "list all accounts")
			fmt.Println("command: list_vaults, usage: ", "list all vaults")
			fmt.Println("command: open_vault, usage: ", "open a particular vaults")
			fmt.Println("command: add_vault, usage: ", "add vault to your account")
			fmt.Println("command: quit, usage: ", "exit the program")
		}
	}
}
