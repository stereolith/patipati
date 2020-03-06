package main

import (
	"os"
	webview "webviewcookies"
)

/*
CLI:
	nwebview URL title
*/
func main() {
	// Open wikipedia in a 800x600 resizable window
	w := webview.New(webview.Settings{
		URL:       os.Args[1],
		Debug:     true,
		Width:     1200,
		Height:    750,
		Resizable: true,
		Title:     os.Args[2],
	})

	w.Exit()
	w.Run()
}
