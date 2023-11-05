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
	profile := utils.GetShellProfile(shell)
	if err != nil {
		log.Fatal(("Error"))
	}

	writeAlias("azsume()", currentDir, profile)
}

func writeAlias(alias string, currentDir string, profile string) {
	exist := checkAliasExists(profile, alias)
	if !exist {
		shellAlias := fmt.Sprintf(`
%s {
	if [ "$1" = "set" ]; then
		eval "$(%s "$@")"
	else
		%[2]s "$@"
	fi
}
`, alias, currentDir)
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
