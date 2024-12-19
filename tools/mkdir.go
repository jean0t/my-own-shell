package tools

import (
	"os"
)

func Mkdir(name string) string {
	var err = os.Mkdir(name, 0775)
	if err != nil {
		return "Error creating directory\n"
	}

	return "Directory created\n"
}
