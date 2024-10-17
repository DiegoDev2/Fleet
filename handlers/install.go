package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Install(pkg string) {
	fmt.Println("Installing " + pkg)

	// Verifica si el archivo existe en la carpeta Formulas
	filePath := filepath.Join("Formulas", pkg+".go")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("Error: El archivo de instalación para %s no se encontró en Formulas/\n", pkg)
		return
	}

	// Ejecuta el archivo Go para instalar el paquete
	cmd := exec.Command("go", "run", filePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error ejecutando la instalación de %s: %v\n", pkg, err)
		return
	}

	// Suponiendo que el binario se llama igual que el paquete
	binaryName := pkg
	binaryPath := filepath.Join("Formulas", binaryName)

	// Mover el binario al PATH (por ejemplo, /usr/local/bin)
	destinationPath := "/usr/local/bin/" + binaryName
	if err := os.Rename(binaryPath, destinationPath); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Error: El binario para %s no se encontró en %s\n", pkg, binaryPath)
		} else {
			fmt.Printf("Error moviendo el binario a %s: %v\n", destinationPath, err)
		}
		return
	}

	fmt.Println("Binario instalado correctamente en:", destinationPath)
	fmt.Println("Done")
}
