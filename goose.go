package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lapt0r/goose/internal/app"
)

func main() {

	var targetPath = flag.String("target", "", "[REQUIRED] The target file or folder")
	var interactive = flag.Bool("interactive", false, "[DEFAULT:FALSE] Runs the application in interactive mode")
	flag.Parse()
	if *targetPath == "" {
		flag.CommandLine.PrintDefaults()
		os.Exit(0)
	}

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
		fmt.Printf("%v\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
