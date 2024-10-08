package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func install(packageName string) {
	fmt.Println("Installing " + packageName + ".zip")

	n, err := http.Get("http://localhost:8080/zip/" + packageName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer n.Body.Close()

	out, err := os.Create(packageName + ".zip")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer out.Close()

	_, err = io.Copy(out, n.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	cmd := exec.Command("unzip", packageName+".zip")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Installation complete.")
}

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <package-name>")
		return
	}

	packageName := os.Args[1]
	install(packageName)
}
