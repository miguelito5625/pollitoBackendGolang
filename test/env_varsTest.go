// Golang program to show the usage of
// Setenv(), Getenv and Unsetenv method

package main

import (
	"fmt"
	"os"
)

// Main function
func main() {

	DB_IP := os.Getenv("DB_IP")
	DB_USER := os.Getenv("DB_USER")
	DB_PASS := os.Getenv("DB_PASS")
	DB_NAME := os.Getenv("DB_NAME")
	ACCESS_SECRET_TOKEN := os.Getenv("ACCESS_SECRET_TOKEN")

	fmt.Println("DB_IP:", DB_IP)
	fmt.Println("DB_USER:", DB_USER)
	fmt.Println("DB_PASS:", DB_PASS)
	fmt.Println("DB_NAME:", DB_NAME)
	fmt.Println("ACCESS_SECRET_TOKEN: ", ACCESS_SECRET_TOKEN)

}
