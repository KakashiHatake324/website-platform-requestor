package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"website-platform-requestor/parser"
)

func main() {
	fmt.Println("Welcome to the Website Platform Requestor.")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Please enter a web address (including 'https://'): ")

		scanner.Scan()
		webAddress := scanner.Text()

		if err := scanner.Err(); err != nil {
			log.Fatal("There was an error with your input:", err)
		}

		if !strings.Contains(webAddress, "http") {
			log.Println("There was an error with your input: does not contain protocol")
			continue
		}

		parsed, err := parser.ParseUrlToPlatform(webAddress)
		if err != nil {
			log.Println("Error Parsing:", err)
			continue
		}

		fmt.Println("The platform is:", parsed)
		fmt.Println("Feel free to post another web address!")
	}
}
