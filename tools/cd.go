package tools

import (
	"os"
  	"strings"
	"fmt"
)

func Cd(path string) string {
	home, _ := os.UserHomeDir() 
	_ = os.Chdir(strings.Replace(path, "~", home, 1))
	return fmt.Sprintf("You are now in: %s", Pwd())
}
