package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// Get arguments
	filePath := flag.String("path", "myapp.log", "The path to the log that should be analyzed")
	logType := flag.String("type", "ERROR", "Log type to search for. Options are: DEBUG, INFO, ERROR and CRITICAL")

	// Fill variables above
	flag.Parse()

	file, err := os.Open(*filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		string, err := reader.ReadString('\n')

		if err != nil {
			break
		}

		if strings.Contains(string, *logType) {
			fmt.Println(string)
		}
	}
}
