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
	file, err := os.Open("assets/goose_header.txt")
	var targetPath = flag.String("target", "NOT_SET", "The target file or folder")
	flag.Parse()
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
	app.Init("", *targetPath)
	fmt.Printf("Initialized with %v rules.  Running..", app.RuleCount())
	app.Run()
}
