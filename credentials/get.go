package credentials

import (
	"azsume/utils"
	"fmt"
	"log"

	"github.com/TwiN/go-color"
	"gopkg.in/ini.v1"
)

func GetCredentials(args []string) string {
	var result string
	section, profile := getSection(args)
	if len(section.KeysHash()) == 0 {
		if profile == "set" {
			log.Fatal(color.Red + "[ERROR] Please select an Azure profile located in ~/.azure/credentials")
		} else {
			log.Fatalf(color.Red+"[ERROR] %s profile does not exist", profile)
		}
	}
	for key, val := range section.KeysHash() {
		result += fmt.Sprintf("export %s=%s\n", key, val)
	}
	result += fmt.Sprintf("export ARM_PROFILE=%s\n", profile)

	return result
}

func getSection(args []string) (*ini.Section, string) {
	var section *ini.Section
	var profile string
	switch len(args) {
	case 3:
		section = utils.GetSection(args[2])
		profile = args[2]
	case 2:
		section = utils.GetSection(args[1])
		profile = args[1]
	default:
		log.Fatal(color.Red + "[ERROR] This section does not exist")
	}
	return section, profile
}
