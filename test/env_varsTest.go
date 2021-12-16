// Golang program to show the usage of
// Setenv(), Getenv and Unsetenv method

package main

import (
	"fmt"
	"os"
)

// Main function
func main() {

	DB_IP := os.Getenv("DB_IPLOCAL")
	fmt.Println("DB_IP:", DB_IP)

	if len(DB_IP) <= 0 {
		DB_IP = "104.197.144.126"
	}

	DB_USER := os.Getenv("DB_USERLOCAL")
	DB_PASS := os.Getenv("DB_PASSLOCAL")
	DB_NAME := os.Getenv("DB_NAMELOCAL")
	ACCESS_SECRET_TOKEN := os.Getenv("ACCESS_SECRET_TOKEN")

	fmt.Println("DB_IP:", DB_IP)
	fmt.Println("DB_USER:", DB_USER)
	fmt.Println("DB_PASS:", DB_PASS)
	fmt.Println("DB_NAME:", DB_NAME)
	fmt.Println("ACCESS_SECRET_TOKEN: ", ACCESS_SECRET_TOKEN)

}
