package main

import (
	"aliasync/cmd"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/pterm/pterm"
	"gopkg.in/validator.v2"
)

func aliasInput() cmd.KeyVal {
	cmd.CLI()

	aliasKeyInput := pterm.InteractiveTextInputPrinter{
		DefaultText: "Alias Key: ",
		TextStyle:   &pterm.ThemeDefault.PrimaryStyle,
	}

	aliasValInput := pterm.InteractiveTextInputPrinter{
		DefaultText: "Alias Value: ",
		TextStyle:   &pterm.ThemeDefault.PrimaryStyle,
	}
	aliasKeyStr, _ := aliasKeyInput.WithMultiLine(false).Show()
	aliasValStr, _ := aliasValInput.WithMultiLine(false).Show()
	pterm.Println()

	aliasModel := cmd.KeyVal{
		Key:   aliasKeyStr,
		Value: aliasValStr,
	}

	if errs := validator.Validate(aliasModel); errs != nil {
		log.Fatal(errs)
	}

	return aliasModel
}

func WriteAlias(keyValStr string) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	aliasData := []byte(keyValStr)
	aliasyncPath := filepath.Join(dirname, ".aliasync")
	existingFile, err := os.OpenFile(aliasyncPath, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Println("File doesn't exist. Creating...")
		err = os.WriteFile(aliasyncPath, aliasData, 0644)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if _, err = existingFile.Write(aliasData); err != nil {
			log.Fatal(err)
		}
	}
	defer existingFile.Close()
}

func main() {

	aliasObj := aliasInput()
	keyValStr := fmt.Sprintf("%s=%s\n", aliasObj.Key, aliasObj.Value)
	pterm.Info.Printfln("Adding following alias: %s %s", aliasObj.Key, aliasObj.Value)

	WriteAlias(keyValStr)

}
