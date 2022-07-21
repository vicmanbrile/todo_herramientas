package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/gocarina/gocsv"
)

type DateTime struct {
	time.Time
}

func (date *DateTime) MarshalCSV() (string, error) {
	return date.Time.Format("2006-01-02 15:04:05"), nil
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", csv)
	return err
}

type Document interface {
	NewFile() File
}

func NewFile(d Document) File {
	return d.NewFile()
}

type File struct {
	path string
	data []string
}

func (f *File) ReadFile(out interface{}) {

	clientsFile, err := os.Open(f.path)
	if err != nil {
		panic(err)
	}

	defer clientsFile.Close()

	gocsv.SetCSVReader(Patter)

	if err := gocsv.UnmarshalFile(clientsFile, out); err != nil {
		panic(err)
	}

	if _, err := clientsFile.Seek(0, 0); err != nil {
		fmt.Println(err)
	}

}

func (f *File) PrintData() {
	if len(f.data) == 0 {
		fmt.Println("No hay Data")
	} else {
		for _, v := range f.data {
			fmt.Println(v)
		}
	}

}

func Patter(in io.Reader) gocsv.CSVReader {
	r := csv.NewReader(in)
	r.Comma = ';'
	return r
}
