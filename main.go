package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"githun.com/haton14/switchgot/request"
)

func main() {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	helpCmd := flag.NewFlagSet("help", flag.ExitOnError)
	if len(os.Args) < 1 {
		log.Fatalf("please specify list")
	}

	switch os.Args[1] {
	case "list":
		token := listCmd.String("token", "dummy", "SwitchBot OpenAPI Token")
		show := listCmd.Bool("show", false, "show result of devices list")
		listCmd.Parse(os.Args[2:])
		client := request.NewClient(*token)
		client.List(*show)
	case "help":
		helpCmd.Parse(os.Args[2:])
		fmt.Println("please specify second args")
		fmt.Println(" ", listCmd.Name())
	default:
		log.Fatalf("please specify list")
	}
}
