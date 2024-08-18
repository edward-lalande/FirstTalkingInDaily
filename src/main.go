package main

import (
	"FirstTalkingInDaily.com/m/src/Utils"
	"FirstTalkingInDaily.com/m/src/header"
	"fmt"
	"log"
	"os"
	"strings"
)

func isAleradyConfigure() bool {
	file, err := Utils.OpenFile(header.FILEPATH_OF_CONFIGURATION, "\n")

	if err != nil {
		return false
	}

	fmt.Println(file)
	fmt.Println(len(file))
	if len(file) != header.SIZE_OF_CONFIGURATION {
		return false
	}

	return true
}

func getMailInformation() (string, string, string) {
	file, err := Utils.OpenFile(header.FILEPATH_OF_CONFIGURATION, "\n")

	if err != nil {
		return "", "", ""
	}

	return strings.Split(file[header.INDEX_OF_DOMAIN_MAIL], ":")[1],
		strings.Split(file[header.INDEX_OF_SENDER_EMAIL], ":")[1],
		strings.Split(file[header.INDEX_OF_SENDER_PASSWORD], ":")[1]
}

func main() {

	value := isAleradyConfigure()

	if !value {
		fmt.Println("Configuration have not been done yet so let's do it!")
		_, err := CreateConfiguration()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	domainMail, email, password := getMailInformation()

	employesName, err := Utils.OpenFile(header.FILEPATH_OF_EMPLOYES_NAME, "\n")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var firstTalker string = strings.ReplaceAll(strings.ToLower(GetRandomName(employesName)), " ", ".") + domainMail

	err = SendEmail(email, password, firstTalker)
	if err != nil {
		log.Println(err)
	}
	os.Exit(0)
}
