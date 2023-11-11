package cli

import (
	"azsume/utils"
	"log"

	"github.com/TwiN/go-color"
)

func Set(args []string) string {
	result, profile := Get(args)
	log.SetFlags(0)
	log.Printf(color.Green + profile + " successfully assumed at " + utils.GetCurrentTime() + color.Reset)
	return result
}
