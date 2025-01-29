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
        			if len(info.Name()) > 20 {
					results = append(results, fmt.Sprintf("%s/\n", info.Name()))
        			} else {
          				results = append(results, fmt.Sprintf("%20s/", info.Name()))
        			}
			} else {
        			if len(info.Name()) > 20{
					results = append(results, fmt.Sprintf("%s\n", info.Name()))
        			} else {
        				results = append(results, fmt.Sprintf("%20s", info.Name()))
        			}
			}
		}
		return strings.Join(results, "\n") + "\n"

	} else {

		for _, entry := range entries {
				var IsDirectory string = "d"
				info, err := entry.Info()
				if err != nil {
					return "Error getting file info\n"
				}
				if info.IsDir() {
					IsDirectory = "-"
				}
				results = append(results, fmt.Sprintf("%s%20s%15d bytes", IsDirectory, info.Name(), info.Size()))
			}

		return strings.Join(results, "\n") + "\n"

	}

}
