package main

import (
	"flag"
	"strings"
)

var moduleName = flag.String("name", "",  "Navnet på modulet")
var parentFolder = "./modules/"

func main() {
	flag.Parse()

	// Vi finder foldername, og ligger den i vores parentfolder
	folderName := strings.ToLower(*moduleName)
	path := parentFolder + folderName

	// Vi sikrer os at vi ikke kommer til at overskrive noget ved en fejl
	// hasFiles, _ := filepathExists(path)

	// if hasFiles {
	// 	println(path)
	// 	panic("Den valgte sti findes allerede. Slet den før du opretter nyt modul")
	// }

	// Vi opretter mapperne, hvori filerne skal være
	createFileFolders(path)

	componentFileName := strings.ToLower(*moduleName) + ".tsx"

	createTypingsTemplate(path)
	createComponentTemplate(path, componentFileName)
	createContextTemplate(path)
	createReducerTemplate(path)
}