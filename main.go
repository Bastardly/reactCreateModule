package main

import (
	"os"
	"strings"
)

var moduleName = ""
var parentFolder = "./modules/"

// Console colors
const colorRed = "\033[31m"
const colorGreen = "\033[32m"
const colorReset = "\033[0m"

func main() {
	if len(os.Args) < 2 {
		panic(getErrorText("Modulet skal have et navn: Kør fx. modulemaker MitModul"))
	}

	moduleName = strings.Title(os.Args[1])


	// Vi finder foldername, og ligger den i vores parentfolder
	folderName := strings.ToLower(moduleName)
	path := parentFolder + folderName

	// Vi sikrer os at vi ikke kommer til at overskrive noget ved en fejl
	hasFiles, _ := filepathExists(path)

	// Hvis modulenavnet allerede findes kaster vi en fejl
	if hasFiles {
		println(path)
		panic(getErrorText("Den valgte sti findes allerede. Slet den før du opretter nyt modul"))
	}

	// Vi opretter mapperne, hvori filerne skal være
	createFileFolders(path)

	componentFileName := strings.ToLower(moduleName)
	componentFileNameWithExtension := componentFileName + ".tsx"

	createTypingsTemplate(path)
	createComponentTemplate(path, componentFileNameWithExtension)
	createContextTemplate(path)
	createReducerTemplate(path)
	createIndexTemplate(path, componentFileName)

	println(string(colorGreen), "\n\n Modulet", moduleName, "blev oprettet her:" + path + "\n\n", string(colorReset))

}