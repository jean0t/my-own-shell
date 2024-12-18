package tools

import (
	"os"
	"fmt"
)

func Env() string {
	var environ = os.Environ()
	var environment = ""
	for _, v := range environ {
		environment += fmt.Sprintf("%s\n", v)
	}

	return environment
}
