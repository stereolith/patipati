package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic("missing argument: url")
	}

	url, _ := url.Parse(os.Args[1])

	icon := getIcon(url.String())

	title := ""
	if len(os.Args) > 2 {
		title = os.Args[2]
	} else {
		title = getTitle(os.Args[1])
	}

	fmt.Println("Installed '" + title + "'")

	install(os.Args[1], title, icon, ".")
}
