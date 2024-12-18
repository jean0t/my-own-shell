package tools

import (
	"os"
	"fmt"
)
func Cd(path string) string {
	_ = os.Chdir(path)
	return fmt.Sprintf("You are now in: %s\n", Pwd())
}
