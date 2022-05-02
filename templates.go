package main

import (
	"html/template"
	"os"
)

type ITypingsData struct {
	Name string
}

type IDataTypes interface {
    ITypingsData
}

type ITemplateInfo struct {
	ModuleName string
	Path string
	FileName string
	TempateString string
}

type ITemplate[T ITypingsData] struct {
	Info ITemplateInfo
	Data T
}

const typingsTemplate = `
export interface I{{ .Name}} {}

export interface I{{ .Name}}Reducer {}
`

func createTemplate[T IDataTypes](data ITemplate[T]) {
	// Indlæser vores templateString som en ny template
	t, err := template.New(data.Info.ModuleName).Parse(data.Info.TempateString)
	if err != nil {
		panic(err)
	}

	// Opretter vores template som en ny fil 
	file, err := os.Create(data.Info.Path + "/" + data.Info.FileName)
	if err != nil {
	  panic(err)
  }

  t.Execute(file, data.Data)
}

func createTypingsTemplate(path string) {
	data := ITemplate {
		Info {

		},
		Data: 
	}

	createTemplate(data)




	t, err := template.New("typingsTemplate").Parse(typingsTemplate)
	  if err != nil {
		  panic(err)
	  }

	  file, err := os.Create(path + "/typingsTemplate.ts")
	  if err != nil {
		panic(err)
	}

	t.Execute(file, ITypingsTemplate{*moduleName})


}