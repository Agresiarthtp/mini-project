package template

import (
	"fmt"
	"os"
)

func Menu() {

	fmt.Println("======================================")
	fmt.Println("\t Welcome to Phincon's Store\t\t")
	fmt.Println("======================================")
	fmt.Println("1. Buy Product")
	fmt.Println("2. Payment Record")
	fmt.Println("3. Exit")
	fmt.Println("======================================")
	var pil int
	fmt.Print("Silahkan input menu : ")
	fmt.Scanln(&pil)

	
	switch pil {
	case 1:
		AddOrder()
	case 2:
		ViewTransaction()
	case 3:
		os.Exit(0)
	}

}

