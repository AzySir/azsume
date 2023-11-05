package utils

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func LoadIni() (*ini.File, error) {
	inidata, err := ini.Load("%s/.azure/credentials", os.Getenv("HOME"))
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

func GetShellProfile(shell string) string {
	result := ""
	switch shell {
	case "/bin/zsh":
		result = fmt.Sprintf("%s/.zshrc", os.Getenv("HOME"))
	case "/bin/bash":
		result = fmt.Sprintf("%s/.bashrc", os.Getenv("HOME"))
	}
	return result
}
