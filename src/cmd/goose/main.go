package main

import(
	"os"
	"log"
	"bufio"
	"fmt"
	"flag"
)

func main() {
	file, err := os.Open("assets/goose_header.txt")
	var target_path = flag.String("target", "NOT_SET", "The target file or folder")
	flag.Parse()
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

	fmt.Printf("Loading targets from [%v]..\n", *target_path)
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}