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

	var targetPath = flag.String("target", "", "[REQUIRED] The target file or folder to scan.  If the target is a valid git repository, Goose will enumerate its commits.")
	var decisiontree = flag.Bool("decisiontree", false, "[DEFAULT:FALSE] Runs goose in decision tree mode.")
	var configpath = flag.String("config", "", "[DEFAULT:EMPTY] Provides a path to configuration file.")
	var help = flag.Bool("help", false, "Print the help screen with command line arguments for Goose.")
	var interactive = flag.Bool("interactive", false, "[DEFAULT:FALSE] Runs the application in interactive mode")
	var commitDepth = flag.Int("commitDepth", 0, "[DEFAULT:0] Specifies the maximum commit depth to scan.")
	var outputmode = flag.String("outputmode", "", "[DEFAULT: EMPTY] Specifies an output mode to use for integration mode.  Goose serialization is the default.")
	var filterPaths = flag.String("ignore", "test", "[DEFAULT: EMPTY] List of path fragments to ignore")
	flag.Parse()
	if *help == true || *targetPath == "" {
		flag.CommandLine.PrintDefaults()
		os.Exit(0)
	}
	if *interactive {
		printHeader()
		if *commitDepth > 0 {
			log.Printf("Scanning git commits to a depth of [%v]..", *commitDepth)
		}
		if *decisiontree {
			log.Printf("Initialized in decision tree mode.  Running..\n")
		} else {
			log.Printf("Initialized with %v rules.  Running..\n", app.RuleCount())
		}
	}
	app.Init(*configpath, *targetPath, *interactive, *commitDepth)
	app.Run(*interactive, *decisiontree, *outputmode, *filterPaths)
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
