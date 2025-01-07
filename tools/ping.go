package tools

import (
	"net/http"
	"fmt"
	"strings"
)

func Ping(url string) string {

	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

	content, err := http.Get(url)
	if err != nil {
		return fmt.Sprintf("Please check your internet connection: %s\n", err)
	}

	defer content.Body.Close()
	return fmt.Sprintf("Status: %s\n", content.Status)
}
