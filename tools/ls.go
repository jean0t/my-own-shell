package tools

import (
	"os"
	"strings"
	"fmt"
)

func Ls(moreInfo bool) string {
	var results []string
	var entries, err = os.ReadDir(".")
	if err != nil {
		return "Error reading directory"
	}

	if !moreInfo {

		for _, entry := range entries {
			info, err := entry.Info()
			if err != nil {
				return "Error getting file info"
			}

			if entry.IsDir() {
				results = append(results, fmt.Sprintf("%s/", info.Name()))
			} else {
				results = append(results, info.Name())
			}

		}

		return strings.Join(results, " \t ") + "\n"

	} else {

		for _, entry := range entries {
				var IsDirectory string = "ND"
				info, err := entry.Info()
				if err != nil {
					return "Error getting file info"
				}
				if info.IsDir() {
					IsDirectory = "D"
				}
				results = append(results, fmt.Sprintf("%s\t%s\t%d bytes", IsDirectory, info.Name(), info.Size()))
			}

		return strings.Join(results, "\n") + "\n"

	}

}
