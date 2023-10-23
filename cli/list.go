package cli

import (
	"azsumego/tablepkg"
	"azsumego/utils"

	"github.com/jedib0t/go-pretty/table"
)

func listCredentials() table.Writer {
	inidata, _ := utils.LoadIni()
	sections := inidata.SectionStrings()
	creds := tablepkg.FormatList(sections[1:])
	return creds
}

func List() table.Writer {
	return listCredentials()
}
