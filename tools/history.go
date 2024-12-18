package tools

import "fmt"

type History struct {
	History []string
}

func (h *History) Append(s string) {
	if len(h.History) < 20 {
		h.History = append(h.History, s)
	} else {
		h.History = append(h.History[1:], s)
	}
}

func (h *History) Show() {
	for index, element := range h.History {
		fmt.Printf("%d\t%s\n", index+1, element)
	}
}
