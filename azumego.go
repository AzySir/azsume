package main

import (
	"azsumego/cli"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(realMain(args))
}

func realMain(args []string) string {
	if len(args) <= 1 {
		cliOptions()
	}
	var result = ""
	switch args[1] {
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
	return result
}

func cliOptions() string {
	var result = `Please input a command.\n
  Options:"
	list	This will list the available service principals
	set		This will set the service principal credentials as ENV variables")
	login	This will login to the service principal")
`
	return result
}