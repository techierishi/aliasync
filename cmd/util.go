package cmd

import (
	"log"

	"github.com/zalando/go-keyring"
)

func SaveCred(keyVal KeyVal) {
	err := keyring.Set("aliasync", keyVal.Key, keyVal.Value)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCred(user string) string {
	str, err := keyring.Get("aliasync", user)
	if err != nil {
		log.Fatal(err)
	}

	return str
}
