package main

import (
	"bufio"
	"flag"
	"fmt"
	"internal/app"
	"log"
	"os"
)

func main() {

	var targetPath = flag.String("target", "NOT_SET", "The target file or folder")
	var interactive = flag.Bool("interactive", false, "Runs the application in interactive mode")
	flag.Parse()
	app.Init("", *targetPath, *interactive)
	if *interactive {
		printHeader()
		fmt.Printf("Initialized with %v rules.  Running..", app.RuleCount())
	}

	app.Run(*interactive)
}

func printHeader() {
	file, err := os.Open("assets/goose_header.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
