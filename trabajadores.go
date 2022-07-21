package main

import "fmt"

type Trabajadores struct{}

var Trabajadores_Path string = "data/trabajadores.csv"

func (h *Trabajadores) NewFile() File {
	file := File{
		path: Trabajadores_Path,
	}

	QuerySelector := []*struct {
		Nombre string `csv:"nombre"`
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
