package credentials

import (
	"azsumego/utils"
	"fmt"
	"log"
)

func GetCredentials(args []string) string {
	section := utils.GetSection(args[2])
	result := ""
	if len(section.KeysHash()) == 0 {
		log.Fatal("[ERROR] This section does not exist")
	}
	for key, val := range section.KeysHash() {
		result += fmt.Sprintf("export %s=%s\n", key, val)
	}

	return result
}
