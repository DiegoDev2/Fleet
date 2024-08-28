package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	githubAPIURL = "https://api.github.com/repos/Homebrew/homebrew-core/contents/Formula/c"
	rawURLBase   = "https://raw.githubusercontent.com/Homebrew/homebrew-core/master/Formula/c/"
	downloadDir  = "formulas"
)

// Formula representa una fórmula de Homebrew
type Formula struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// getFormulas obtiene la lista de archivos de fórmulas desde GitHub
func getFormulas() ([]Formula, error) {
	resp, err := http.Get(githubAPIURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la solicitud: %s", resp.Status)
	}

	var formulas []Formula
	if err := json.NewDecoder(resp.Body).Decode(&formulas); err != nil {
		return nil, err
	}

	return formulas, nil
}

// downloadFile descarga un archivo desde una URL y lo guarda en el destino
func downloadFile(url, dest string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	// Crea el directorio de destino si no existe
	if err := os.MkdirAll(downloadDir, os.ModePerm); err != nil {
		fmt.Printf("Error al crear el directorio %s: %v\n", downloadDir, err)
		return
	}

	// Obtiene la lista de fórmulas
	formulas, err := getFormulas()
	if err != nil {
		fmt.Println("Error al obtener la lista de fórmulas:", err)
		return
	}

	// Descarga cada archivo de fórmula
	for _, formula := range formulas {
		if formula.Path != "" {
			fileURL := rawURLBase + formula.Name
			destPath := downloadDir + "/" + formula.Name

			fmt.Printf("Descargando %s...\n", formula.Name)
			if err := downloadFile(fileURL, destPath); err != nil {
				fmt.Printf("Error al descargar %s: %v\n", formula.Name, err)
				continue
			}

			fmt.Printf("Archivo %s descargado exitosamente.\n", formula.Name)
		}
	}

	fmt.Println("Descarga completada.")
}
