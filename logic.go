package main

import (
	"github.com/gocarina/gocsv"
)

type Document interface {
	NewFile(path string) File
}

type File struct {
	path string
	out  []interface{}
}


func (f *File) ReadFile(path string) {

	clientsFile, err := os.Open(f.path)
	if err != nil {
		panic(err)
	}

	defer clientsFile.Close()

	gocsv.SetCSVReader(Patter)

	if err := gocsv.UnmarshalFile(clientsFile, &f.out); err != nil {
		panic(err)
	}

	for _, i := range f.out {
		fmt.Println(i.Nombre)
	}


	if _, err := clientsFile.Seek(0, 0); err != nil {
		fmt.Println(err)
	}

}

func Patter(in io.Reader) gocsv.CSVReader {
	r := csv.NewReader(in)
	r.Comma = ';'
	return r
}
