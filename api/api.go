package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func zipHandler(w http.ResponseWriter, r *http.Request) {

	packageName := filepath.Base(r.URL.Path)
	zipFile := "packages/" + packageName + ".zip"

	fmt.Println("Serving file:", zipFile)

	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	data, err := ioutil.ReadFile(zipFile)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	f, err := writer.Create(packageName + ".zip")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s.zip", packageName))
	w.Write(buf.Bytes())
}

func main() {

	http.HandleFunc("/zip/", zipHandler)
	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
