package cli

import (
	"azsume/credentials"
	"log"
	"os"
)

// Get Credentials from .azure/credentials
func Get(args []string) (string, string) {
	if len(args) < 2 {
		log.Fatal("[ERROR] Please provide a name of the Credentials to Get")
		os.Exit(1)
	}
	result, profile := credentials.GetCredentials(args)
	return result, profile
}
