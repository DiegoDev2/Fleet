package main

// Libxlsxwriter - C library for creating Excel XLSX files
// Homepage: https://libxlsxwriter.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installLibxlsxwriter() {
	// Método 1: Descargar y extraer .tar.gz
	libxlsxwriter_tar_url := "https://github.com/jmcnamara/libxlsxwriter/archive/refs/tags/v1.1.8.tar.gz"
	libxlsxwriter_cmd_tar := exec.Command("curl", "-L", libxlsxwriter_tar_url, "-o", "package.tar.gz")
	err := libxlsxwriter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxlsxwriter_zip_url := "https://github.com/jmcnamara/libxlsxwriter/archive/refs/tags/v1.1.8.zip"
	libxlsxwriter_cmd_zip := exec.Command("curl", "-L", libxlsxwriter_zip_url, "-o", "package.zip")
	err = libxlsxwriter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxlsxwriter_bin_url := "https://github.com/jmcnamara/libxlsxwriter/archive/refs/tags/v1.1.8.bin"
	libxlsxwriter_cmd_bin := exec.Command("curl", "-L", libxlsxwriter_bin_url, "-o", "binary.bin")
	err = libxlsxwriter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxlsxwriter_src_url := "https://github.com/jmcnamara/libxlsxwriter/archive/refs/tags/v1.1.8.src.tar.gz"
	libxlsxwriter_cmd_src := exec.Command("curl", "-L", libxlsxwriter_src_url, "-o", "source.tar.gz")
	err = libxlsxwriter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxlsxwriter_cmd_direct := exec.Command("./binary")
	err = libxlsxwriter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
