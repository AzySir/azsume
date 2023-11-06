package cli

import (
	"azsume/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func Init() {
	// Get Current Directory
	currentDir, err := os.Executable()
	shell := os.Getenv("SHELL")
	log.Println("Getting Profile...")
	profile := utils.GetShellProfile(shell)
	if err != nil {
		log.Fatal(("Error"))
	}
	log.Println("Writing Alias...")
	writeAlias("azsume()", currentDir, profile)
}

func writeAlias(alias string, currentDir string, profile string) {
	exist := checkAliasExists(profile, alias)
	log.Printf("Exist: %s", exist)
	if exist {
		shellAlias := fmt.Sprintf(`
function %s {
	clicommands=("init" "list" "get" "sync" "login")
	if [ "$1" = "set" ]; then
		eval $("%[2]s" "$@")
	elif [[ ! ("${clicommands[@]}" =~ "${1}") ]]; then
		eval $("%[2]s" "$@")
	else
		%[2]s "$@"
	fi
}
`, alias, currentDir)
		log.Printf(shellAlias)
		f, _ := os.OpenFile(profile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		_, err := f.WriteString(shellAlias)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func checkAliasExists(profile string, alias string) bool {
	body, _ := os.ReadFile(profile)
	contains := strings.Contains(string(body), alias)
	return contains
}
