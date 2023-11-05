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
			result = cliOptions()
		}
	}
	return result
}

func cliOptions() string {
	var result = `Please input a command.\n
  Options:
	list\tThis will list the available service principals
	set\tThis will set the service principal credentials as ENV variables")
	login\tThis will login to the service principal")
`
	return result
}
