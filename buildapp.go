package main

import (
	"fmt"
	"go-homedir"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func toFilenameStr(s string) string {
	// Make a Regex to say we only want letters and numbers
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(s, "")
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func copy(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

// install webview binary and create desktop shortcut
func install(url string, name string, iconUrl string, dest string) {

	ex, _ := os.Executable()
	appDir := filepath.Dir(filepath.Dir(ex))
	homeDir, _ := homedir.Dir()
	homeAppDir := homeDir + "/.patipati"

	os.Mkdir(homeAppDir, os.ModePerm)

	// install patipati-nwebview
	if _, err := copy(appDir+"/nwebview/nwebview", homeAppDir+"/nwebview-"+toFilenameStr(name)); err != nil {
		panic(err)
	}

	os.Chmod(homeAppDir+"/nwebview-"+toFilenameStr(name), 0755)

	// download icon
	iconPath := homeAppDir + "/" + toFilenameStr(name) + ".png"
	if iconUrl != "" {
		downloadIcon(iconUrl, iconPath)
	} else {
		if _, err := copy(appDir+"/nwebview/nwebview.png", iconPath); err != nil {
			panic(err)
		}
	}

	// write .desktop-file
	desktop := []byte(`[Desktop Entry]
Name=` + name + `
Exec=` + homeAppDir + `/nwebview-` + toFilenameStr(name) + ` ` + url + ` "` + name + `"
Icon=` + iconPath + `
Type=Application
Categories=Utility;
StartupWMClass=nwebview-` + toFilenameStr(name))

	ioutil.WriteFile(homeDir+"/.local/share/applications/patipati-"+toFilenameStr(name)+".desktop", desktop, 0711)
}

func downloadIcon(url string, dest string) {
	if err := DownloadFile(dest, url); err != nil {
		panic(err)
	}
}
