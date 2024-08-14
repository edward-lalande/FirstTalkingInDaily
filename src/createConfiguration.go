package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func getDomainMailName() error {
	configurationFile, confErr := os.OpenFile(FILEPATH_OF_CONFIGURATION, os.O_CREATE|os.O_WRONLY, 0600)
	defer configurationFile.Close()

	if confErr != nil {
		return errors.New("Error on creating configuration file")
	}

	fmt.Println("Configuration have not been done yet so let's do it!")
	fmt.Print("Enter the domain mail of your company like this example: [@company.com]: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return errors.New("Error on reading input")
	}

	var domainNameMail string = scanner.Text()
	configurationFile.WriteString(fmt.Sprintf("Domain Name:%s\n", domainNameMail))

	return nil
}

func getEmployeeMailName() error {
	scanner := bufio.NewScanner(os.Stdin)
	employesNameFile, nameErr := os.OpenFile(FILEPATH_OF_EMPLOYES_NAME, os.O_CREATE|os.O_WRONLY, 0600)
	defer employesNameFile.Close()

	if nameErr != nil {
		return errors.New("Error on creating configuration file")
	}

	fmt.Println("Write the name of the people on your team like this example: [Name LastName]")
	fmt.Println("When you have finish just press CTRL + D or CTRL + Z")
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		employesNameFile.WriteString(fmt.Sprintf("%s\n", text))
	}

	return nil
}

func CreateConfiguration() (bool, error) {
	domainMailError := getDomainMailName()

	if domainMailError != nil {
		return false, domainMailError
	}

	employesNameError := getEmployeeMailName()

	if employesNameError != nil {
		return false, employesNameError
	}

	return true, nil
}
