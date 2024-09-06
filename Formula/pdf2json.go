package main

// Pdf2json - PDF to JSON and XML converter
// Homepage: https://github.com/flexpaper/pdf2json

import (
	"fmt"
	
	"os/exec"
)

func installPdf2json() {
	// Método 1: Descargar y extraer .tar.gz
	pdf2json_tar_url := "https://github.com/flexpaper/pdf2json/archive/refs/tags/0.71.tar.gz"
	pdf2json_cmd_tar := exec.Command("curl", "-L", pdf2json_tar_url, "-o", "package.tar.gz")
	err := pdf2json_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pdf2json_zip_url := "https://github.com/flexpaper/pdf2json/archive/refs/tags/0.71.zip"
	pdf2json_cmd_zip := exec.Command("curl", "-L", pdf2json_zip_url, "-o", "package.zip")
	err = pdf2json_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pdf2json_bin_url := "https://github.com/flexpaper/pdf2json/archive/refs/tags/0.71.bin"
	pdf2json_cmd_bin := exec.Command("curl", "-L", pdf2json_bin_url, "-o", "binary.bin")
	err = pdf2json_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pdf2json_src_url := "https://github.com/flexpaper/pdf2json/archive/refs/tags/0.71.src.tar.gz"
	pdf2json_cmd_src := exec.Command("curl", "-L", pdf2json_src_url, "-o", "source.tar.gz")
	err = pdf2json_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pdf2json_cmd_direct := exec.Command("./binary")
	err = pdf2json_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
