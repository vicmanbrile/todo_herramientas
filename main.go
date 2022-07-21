package main

import "net/http"

func main() {

	http.HandleFunc("/", Home)
	http.ListenAndServe(":8080", nil)
}

/* Insert data to a document

csvContent, err := gocsv.MarshalString(&outs)
if err != nil {
	panic(err)
}

ioutil.WriteFile("data.csv", []byte(csvContent), 0644)

*/
