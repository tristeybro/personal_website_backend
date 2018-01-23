package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type EmailCredentials struct {
	Email    string
	Password string
}

func importCredentials(pathToCredentials string) EmailCredentials {
	raw, err := ioutil.ReadFile(pathToCredentials)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var emailCredentials EmailCredentials
	json.Unmarshal(raw, &emailCredentials)
	return emailCredentials
}

var Credentials = importCredentials("credentials/credentials.json")
