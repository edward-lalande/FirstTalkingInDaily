package main

import (
	"FirstTalkingInDaily.com/m/src/header"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func getDomainMailName() error {
	configurationFile, confErr := os.OpenFile(header.FILEPATH_OF_CONFIGURATION, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	defer configurationFile.Close()

	if confErr != nil {
		return errors.New("Error on creating configuration file")
	}

	scanner := bufio.NewScanner(os.Stdin)

	var getDomainMailScan func() (string, error)

	getDomainMailScan = func() (string, error) {
		fmt.Print("Enter the domain mail of your company like this example: [@company.com]: ")
		if !scanner.Scan() {
			return "", errors.New("Error on reading input")
		}
		var domainNameMail string = scanner.Text()
		if strings.Index(domainNameMail, "@") == -1 {
			fmt.Println("Domain name is not valid!")
			return getDomainMailScan()
		}
		return domainNameMail, nil
	}

	domainNameMail, err := getDomainMailScan()
	if err != nil {
		return err
	}

	configurationFile.WriteString(fmt.Sprintf("Domain Name:%s\n", domainNameMail))

	return nil
}

func getEmailAndPassword() error {

	configurationFile, confErr := os.OpenFile(header.FILEPATH_OF_CONFIGURATION, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	scanner := bufio.NewScanner(os.Stdin)

	var getMail func() (string, error)
	getMail = func() (string, error) {

		if confErr != nil {
			return "", confErr
		}

		fmt.Print("Enter your email address: ")
		if !scanner.Scan() {
			return "", errors.New("Error on reading input")
		}
		email := scanner.Text()
		if strings.Index(email, "@") == -1 {
			fmt.Println("Invalid email address")
			value, err := getMail()
			if err == nil {
				return value, nil
			}
		}
		return email, nil

	}
	email, err := getMail()

	if err != nil {
		return errors.New("Error on reading input")
	}
	configurationFile.WriteString(fmt.Sprintf("mail:%s\n", email))
	fmt.Print("Enter your password: ")
	if !scanner.Scan() {
		return errors.New("Error on reading input")
	}
	configurationFile.WriteString(fmt.Sprintf("password:%s\n", scanner.Text()))

	return nil
}

func getEmployeeMailName() error {
	scanner := bufio.NewScanner(os.Stdin)
	employesNameFile, nameErr := os.OpenFile(header.FILEPATH_OF_EMPLOYES_NAME, os.O_CREATE|os.O_WRONLY, 0600)
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

	emailError := getEmailAndPassword()

	if emailError != nil {
		return false, emailError
	}

	employesNameError := getEmployeeMailName()

	if employesNameError != nil {
		return false, employesNameError
	}

	return true, nil
}
