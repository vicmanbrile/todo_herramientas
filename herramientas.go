package main

import "fmt"

type Herramienta struct{}

var Herramientas_Path string = "data/herramienta.csv"

func (h *Herramienta) NewFile() File {
	file := File{
		path: Herramientas_Path,
	}

	QuerySelector := []*struct {
		Descripcion string `csv:"descripcion"`
	}{}

	file.ReadFile(&QuerySelector)

	var data []string
	for _, v := range QuerySelector {
		value := fmt.Sprintf("%v", *v)
		data = append(data, value)
	}

	file.data = data

	return file

}
