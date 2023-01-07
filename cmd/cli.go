package cmd

import (
	"log"
	"os"

	"github.com/pterm/pterm"
	"github.com/urfave/cli"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/validator.v2"
)

var app = cli.NewApp()

func info() {
	app.Name = "Alias Sync App"
	app.Usage = "This app keeps all you alias in sync!"
	app.Author = "Rishi"
	app.Version = "1.0.0"
}

func loginInput(username string) KeyVal {
	pterm.Print(pterm.Cyan("Password: "))
	password, _ := terminal.ReadPassword(0)
	pterm.Println()

	aliasModel := KeyVal{
		Key:   username,
		Value: string(password),
	}

	if errs := validator.Validate(aliasModel); errs != nil {
		log.Fatal(errs)
	}

	return aliasModel
}

func commands() {
	var username string
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add",
			Action: func(c *cli.Context) {

				log.Println(c.Args())
				createAlias()
			},
		},
		{
			Name:    "sync",
			Aliases: []string{"s"},
			Usage:   "Sync all alias",
			Action: func(c *cli.Context) {
				log.Println("Sync running...")
			},
		},
		{
			Name:    "login",
			Aliases: []string{"l"},
			Usage:   "Login",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "u",
					Usage:       "Enter username",
					Destination: &username,
				},
			},
			Action: func(c *cli.Context) {
				cred := loginInput(username)
				SaveCred(cred)
			},
		},
	}
}
func CLI() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
