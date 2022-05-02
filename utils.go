package main

import (
	"log"
	"os"
)

// filepathExists returns bool and error
func filepathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil { return true, nil }
    if os.IsNotExist(err) { return false, nil }
    return false, err
}

func createFileFolders(path string) {

	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Println(err)
	}
}

func getErrorText(text string) string {
	return colorRed + " " + text + " " + colorReset
}