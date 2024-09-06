package main

// Pdfcpu - PDF processor written in Go
// Homepage: https://pdfcpu.io

import (
	"fmt"
	
	"os/exec"
)

func installPdfcpu() {
	// Método 1: Descargar y extraer .tar.gz
	pdfcpu_tar_url := "https://github.com/pdfcpu/pdfcpu/archive/refs/tags/v0.8.1.tar.gz"
	pdfcpu_cmd_tar := exec.Command("curl", "-L", pdfcpu_tar_url, "-o", "package.tar.gz")
	err := pdfcpu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdfcpu_zip_url := "https://github.com/pdfcpu/pdfcpu/archive/refs/tags/v0.8.1.zip"
	pdfcpu_cmd_zip := exec.Command("curl", "-L", pdfcpu_zip_url, "-o", "package.zip")
	err = pdfcpu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdfcpu_bin_url := "https://github.com/pdfcpu/pdfcpu/archive/refs/tags/v0.8.1.bin"
	pdfcpu_cmd_bin := exec.Command("curl", "-L", pdfcpu_bin_url, "-o", "binary.bin")
	err = pdfcpu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdfcpu_src_url := "https://github.com/pdfcpu/pdfcpu/archive/refs/tags/v0.8.1.src.tar.gz"
	pdfcpu_cmd_src := exec.Command("curl", "-L", pdfcpu_src_url, "-o", "source.tar.gz")
	err = pdfcpu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdfcpu_cmd_direct := exec.Command("./binary")
	err = pdfcpu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
