package main

// Epubcheck - Validate EPUB files, version 2.0 and later
// Homepage: https://github.com/w3c/epubcheck

import (
	"fmt"
	
	"os/exec"
)

func installEpubcheck() {
	// Método 1: Descargar y extraer .tar.gz
	epubcheck_tar_url := "https://github.com/w3c/epubcheck/releases/download/v5.1.0/epubcheck-5.1.0.zip"
	epubcheck_cmd_tar := exec.Command("curl", "-L", epubcheck_tar_url, "-o", "package.tar.gz")
	err := epubcheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	epubcheck_zip_url := "https://github.com/w3c/epubcheck/releases/download/v5.1.0/epubcheck-5.1.0.zip"
	epubcheck_cmd_zip := exec.Command("curl", "-L", epubcheck_zip_url, "-o", "package.zip")
	err = epubcheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	epubcheck_bin_url := "https://github.com/w3c/epubcheck/releases/download/v5.1.0/epubcheck-5.1.0.zip"
	epubcheck_cmd_bin := exec.Command("curl", "-L", epubcheck_bin_url, "-o", "binary.bin")
	err = epubcheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	epubcheck_src_url := "https://github.com/w3c/epubcheck/releases/download/v5.1.0/epubcheck-5.1.0.zip"
	epubcheck_cmd_src := exec.Command("curl", "-L", epubcheck_src_url, "-o", "source.tar.gz")
	err = epubcheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	epubcheck_cmd_direct := exec.Command("./binary")
	err = epubcheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
