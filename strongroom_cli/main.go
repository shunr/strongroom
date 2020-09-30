package main

import (
	"log"
	"os"
	"sort"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang, l",
				Value: "english",
				Usage: "Language for the greeting",
			},
			&cli.StringFlag{
				Name:  "config, c",
				Usage: "Load configuration from `FILE`",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add a task to the list",
				Action: func(c *cli.Context) error {
					return nil
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
	// cmd := os.Args[1]
	// switch cmd {
	// case "create":
	// 	reader := bufio.NewReader(os.Stdin)
	// 	fmt.Print("Username: ")
	// 	username, _ := reader.ReadString('\n')
	// 	fmt.Print("Password: ")
	// 	password, _ := reader.ReadString('\n')
	// 	username = strings.TrimSpace(username)
	// 	password = strings.TrimSpace(password)
	// 	account := api.CreateAccount(username, password)
	// 	err := account.ExportToFile(os.Args[2])
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	break
	// case "load":
	// 	account, err := api.ImportAccountFromFile(os.Args[2])
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	fmt.Print("Password: ")
	// 	reader := bufio.NewReader(os.Stdin)
	// 	password, _ := reader.ReadString('\n')
	// 	muk := util.DeriveKeyFromMasterPasswordAndSecretKey(account.Username, password, account.SecretKey, account.MasterUnlockSalt)
	// 	private, err := crypto.AESGCMDecrypt(account.EncryptedPrivateKey, muk, account.PrivateKeyNonce)
	// 	if err != nil {
	// 		fmt.Println("Wrong master password")
	// 		os.Exit(-1)
	// 	} else {
	// 		fmt.Println(string(private))
	// 	}

	// 	break
	// }

}
