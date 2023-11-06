package main

import (
	"azsume/cli"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(realMain(args))
}

func realMain(args []string) string {
	var result = ""
	if len(args) <= 1 {
		result = cliOptions()
	} else {
		switch args[1] {
		case "init":
			cli.Init()
		case "list":
			cli.List()
		case "set":
			result = cli.Set(args)
		case "get":
			result = cli.Get(args)
		case "sync":
			fmt.Println("Sync")
		case "login":
			result = "login"
		default:
			result = cli.Set(args)
		}
	}
	return result
}

func cliOptions() string {
	var result = `Please input a command.
  Options:
	init  - Set Alias in SHELL profile to allow for "set" functionality to work
	list  - This will list the available service principals
	set   - This will set the service principal credentials as ENV variables")
	login - This will login to the service principal")
`
	return result
}
