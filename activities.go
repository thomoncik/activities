package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// listCommand := flag.NewFlagSet("list", flag.ExitOnError)
	startCommand := flag.NewFlagSet("start", flag.ExitOnError)
	stopCommand := flag.NewFlagSet("stop", flag.ExitOnError)
	// clearCommand := flag.NewFlagSet("clear", flag.ExitOnError)

	var startName string
	var stopName string

	startCommand.StringVar(&startName, "name", "", "Name of activity")
	stopCommand.StringVar(&stopName, "name", "", "Name of activity")

	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("Usage: avtivities <list | start | stop | clear>")
		os.Exit(2)
	}

	switch os.Args[1] {
	case "list":
		fmt.Println("Listing .... done")
	case "start":
		startCommand.Parse(os.Args[2:])
		if startName == "" {
			startCommand.PrintDefaults()
			os.Exit(2)
		}
		fmt.Printf("Starting %v\n", startName)
	case "stop":
		stopCommand.Parse(os.Args[2:])
		if stopName == "" {
			stopCommand.PrintDefaults()
			os.Exit(2)
		}
		fmt.Printf("Stopping %v\n", stopName)
	case "clear":
		fmt.Println("Clearing done")
	default:
		flag.PrintDefaults()
		os.Exit(2)
	}
}
