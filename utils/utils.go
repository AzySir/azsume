package utils

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func LoadIni() (*ini.File, error) {
	inidata, err := ini.Load("/Users/azysir/.azure/credentials")
	if err != nil {
		log.Printf("[ERROR] Fail to read file: %v", err)
		os.Exit(1)
	}

	return inidata, err
}

func GetSection(section string) *ini.Section {
	inidata, _ := LoadIni()
	_section := inidata.Section(section)
	return _section
}
