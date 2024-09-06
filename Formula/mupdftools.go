package main

// MupdfTools - Lightweight PDF and XPS viewer
// Homepage: https://mupdf.com/

import (
	"fmt"
	
	"os/exec"
)

func installMupdfTools() {
	// Método 1: Descargar y extraer .tar.gz
	mupdftools_tar_url := "https://mupdf.com/downloads/archive/mupdf-1.23.11-source.tar.gz"
	mupdftools_cmd_tar := exec.Command("curl", "-L", mupdftools_tar_url, "-o", "package.tar.gz")
	err := mupdftools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mupdftools_zip_url := "https://mupdf.com/downloads/archive/mupdf-1.23.11-source.zip"
	mupdftools_cmd_zip := exec.Command("curl", "-L", mupdftools_zip_url, "-o", "package.zip")
	err = mupdftools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mupdftools_bin_url := "https://mupdf.com/downloads/archive/mupdf-1.23.11-source.bin"
	mupdftools_cmd_bin := exec.Command("curl", "-L", mupdftools_bin_url, "-o", "binary.bin")
	err = mupdftools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mupdftools_src_url := "https://mupdf.com/downloads/archive/mupdf-1.23.11-source.src.tar.gz"
	mupdftools_cmd_src := exec.Command("curl", "-L", mupdftools_src_url, "-o", "source.tar.gz")
	err = mupdftools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mupdftools_cmd_direct := exec.Command("./binary")
	err = mupdftools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
