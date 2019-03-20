package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"syscall"

	"gitlab.com/tovijaeschke/goEncrypt/Encryption"
	"golang.org/x/crypto/ssh/terminal"
)

func GetPassword() string {
	fmt.Print("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Unable to get password")
	}
	fmt.Print("\n")
	return string(password)
}

func main() {
	encFlag := flag.Bool("e", false, "Encrypt a file")
	decFlag := flag.Bool("d", false, "Decrypt a file")
	helpFlag := flag.Bool("h", false, "Shows help message")

	flag.Parse()

	if (!*encFlag && !*decFlag) || *helpFlag {
		var CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
		fmt.Fprintf(CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	if len(flag.Args()) != 1 {
		fmt.Println(errors.New(fmt.Sprintf("Error: did not provide file to encrypt/decrypt")))
		os.Exit(1)
	}

	for i := 0; i < len(flag.Args()); i++ {
		if _, err := os.Stat(flag.Args()[i]); os.IsNotExist(err) {
			fmt.Println(errors.New(fmt.Sprintf("Error: %s does not exist", flag.Args()[i])))
			os.Exit(1)
		}
	}

	if *encFlag {
		pass := GetPassword()
		for i := 0; i < len(flag.Args()); i++ {
			err := Encryption.EncryptFile(pass, flag.Args()[i])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Printf("Sucessfully encrypted %s\n", flag.Args()[i])
			}
		}
	} else if *decFlag {
		pass := GetPassword()
		for i := 0; i < len(flag.Args()); i++ {
			err := Encryption.DecryptFile(pass, flag.Args()[i])
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				fmt.Printf("Sucessfully decrypted %s\n", flag.Args()[i])
			}
		}
	}
}
