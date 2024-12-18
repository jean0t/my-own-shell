package tools

import (
	"os"
	"fmt"
)

func Rename(oldname, newname string) string {
	var err error = os.Rename(oldname, newname)
	if err != nil {
		return "An error has occurred when renaming the file.\n"
	}
	return fmt.Sprintf("Renamed Successfully: %s -> %s\n", oldname, newname)
}
