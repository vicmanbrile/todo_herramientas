package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}

/*
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
*/


/* Insert data to a document

		csvContent, err := gocsv.MarshalString(&outs)
		if err != nil {
			panic(err)
		}

		ioutil.WriteFile("data.csv", []byte(csvContent), 0644)

*/

