package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var vaultName string

type PasswordVault struct {
	Name     string
	Password string
}

func main() {
	printInstructions()
}

func printInstructions() {

	var isTrue bool = true

	for isTrue {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Create a new password vault")
		fmt.Println("2. Sign in to a password vault")
		fmt.Println("3. Add a password to a vault")
		fmt.Println("4. Fetch a password from a vault")
		fmt.Println("Quit (enter q or quit)")

		var passwordObtained string
		var input int
		fmt.Scan(&input)

		switch input {
		case 1:
			createVault()
		case 2:
			vaultName = signInToVault()
			fmt.Println("Vault name obtained is ", vaultName)
		case 3:
			addPaswordToVault(vaultName)
		case 4:
			passwordObtained = fetchPasswordFromVault(vaultName)
			fmt.Println("Password obtained from ", vaultName+" is ", passwordObtained)

		default:
			isTrue = false
		}
	}

}

func createVault() {

	pw := PasswordVault{}

	fmt.Println("Creating a new vault")
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please provide a name for the vault: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}
	pw.Name = strings.TrimSpace(name)

	fmt.Print("Please enter a master password: ")
	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}
	pw.Password = strings.TrimSpace(password)

	fmt.Print("Please confirm the master password: ")
	mPassword, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}
	if mPassword != password {
		fmt.Println("Password not matched")
		return
	}

	fmt.Println()
	saveVaultTotext(pw.Name, pw)
}

func signInToVault() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the valult name: ")
	vault, err := reader.ReadString('\n')
	vault = strings.TrimSpace(vault)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		os.Exit(0)
	}

	data, err := os.ReadFile(vault)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		os.Exit(0)
	}

	var pw PasswordVault

	err = json.Unmarshal([]byte(data), &pw)
	if err != nil {
		fmt.Println("Error reading data: ", err)
		os.Exit(0)
	}

	fmt.Print("Enter the vault password: ")
	password, err := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	if pw.Name == vault && pw.Password == password {
		fmt.Println("Signed into the vault successfully..!")
		return pw.Name
	} else {
		fmt.Println("Invalid vault name or password.")
		os.Exit(1)
	}
	fmt.Println()
	return ""
}

func addPaswordToVault(vaultName string) {

	var pw PasswordVault

	data, err := os.ReadFile(vaultName)
	err = json.Unmarshal([]byte(data), &pw)

	fmt.Print("Enter the password to add in the vault : ")
	reader := bufio.NewReader(os.Stdin)

	password, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error in reading the file", err)
		os.Exit(0)
	}

	password = strings.TrimSpace(password)
	pw.Password = password

	updatedPw, err := json.Marshal(pw)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}
	//write back the file with updated password
	err = os.WriteFile(vaultName, updatedPw, 0644)

	fmt.Println("Password updated successfully")
	fmt.Println()

}

func fetchPasswordFromVault(vaultName string) string {

	data, err := os.ReadFile(vaultName)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		os.Exit(0)
	}

	var pw PasswordVault

	err = json.Unmarshal([]byte(data), &pw)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		os.Exit(0)
	}
	fmt.Println()
	return pw.Password

}

func saveVaultTotext(filename string, p PasswordVault) error {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file: ", err)
		return err
	}
	defer file.Close()

	//convert to json using marshall
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return err
	}
	_, err = file.WriteString(string(data))

	fmt.Println("Vault succesffully created..")
	fmt.Println()
	return nil

}
