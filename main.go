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
		panic(getErrorText("The module needs a name as an argument: For instance ./modulemaker MyModule"))
	}

	moduleName = strings.Title(os.Args[1])
	folderName := strings.ToLower(moduleName)
	path := parentFolder + folderName

	// We make use that we don't delete any files by mistake.
	hasFiles, _ := filepathExists(path)

	if hasFiles {
		println(path)
		panic(getErrorText("The path already exist. Please delete the path folder, or call your module something else."))
	}

	createFileFolders(path)

	componentFileName := strings.ToLower(moduleName)
	componentFileNameWithExtension := componentFileName + ".tsx"

	// Create files from templates
	createTypingsTemplate(path)
	createComponentTemplate(path, componentFileNameWithExtension)
	createContextTemplate(path)
	createReducerTemplate(path)
	createIndexTemplate(path, componentFileName)
	createActionsTemplate(path)

	println(string(colorGreen), "\n\n The module", moduleName, "was created here:" + path + "\n\n", string(colorReset))

}