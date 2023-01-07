package cmd

import (
	"fmt"
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

func loginInput() KeyVal {
	aliasKeyInput := pterm.InteractiveTextInputPrinter{
		DefaultText: "Uername: ",
		TextStyle:   &pterm.ThemeDefault.PrimaryStyle,
	}

	aliasKeyStr, _ := aliasKeyInput.WithMultiLine(false).Show()
	fmt.Println("Password: ")
	password, _ := terminal.ReadPassword(0)
	pterm.Println()

	aliasModel := KeyVal{
		Key:   aliasKeyStr,
		Value: string(password),
	}

	if errs := validator.Validate(aliasModel); errs != nil {
		log.Fatal(errs)
	}

	return aliasModel
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "sync",
			Aliases: []string{"p"},
			Usage:   "Sync all alias",
			Action: func(c *cli.Context) {
				cred := loginInput()
				SaveCred(cred)
			},
		},
		{
			Name:    "login",
			Aliases: []string{"pa"},
			Usage:   "Login",
			Action: func(c *cli.Context) {
				cred := loginInput()
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
