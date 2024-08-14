package main

import (
	"FirstTalkingInDaily.com/m/src/Utils"
	"log"
	"os"
	"strings"
)

const FILEPATH_OF_EMPLOYES_NAME string = ".employes_name.txt"
const FILEPATH_OF_CONFIGURATION string = ".config.txt"
const INDEX_OF_COMPANY_NAME int = 1

func isAleradyConfigure() bool {
	_, err := os.Open(FILEPATH_OF_CONFIGURATION)

	if err != nil {
		return false
	}
	return true
}

func getDomainMail() string {
	domainMail, err := Utils.OpenFile(FILEPATH_OF_CONFIGURATION, ":")

	if err != nil {
		return ""
	}

	return domainMail[INDEX_OF_COMPANY_NAME]
}

func main() {

	if !isAleradyConfigure() {
		_, err := CreateConfiguration()

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	domainMail := getDomainMail()

	employesName, err := Utils.OpenFile(FILEPATH_OF_EMPLOYES_NAME, "\n")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	var firstTalker string = strings.ReplaceAll(strings.ToLower(GetRandomName(employesName)), " ", ".") + domainMail

	err = SendEmail(firstTalker)
	if err != nil {
		log.Println(err)
	}
	os.Exit(0)
}
