package main

// Html2text - Advanced HTML-to-text converter
// Homepage: https://github.com/grobian/html2text

import (
	"fmt"
	
	"os/exec"
)

func installHtml2text() {
	// Método 1: Descargar y extraer .tar.gz
	html2text_tar_url := "https://github.com/grobian/html2text/releases/download/v2.2.3/html2text-2.2.3.tar.gz"
	html2text_cmd_tar := exec.Command("curl", "-L", html2text_tar_url, "-o", "package.tar.gz")
	err := html2text_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	html2text_zip_url := "https://github.com/grobian/html2text/releases/download/v2.2.3/html2text-2.2.3.zip"
	html2text_cmd_zip := exec.Command("curl", "-L", html2text_zip_url, "-o", "package.zip")
	err = html2text_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	html2text_bin_url := "https://github.com/grobian/html2text/releases/download/v2.2.3/html2text-2.2.3.bin"
	html2text_cmd_bin := exec.Command("curl", "-L", html2text_bin_url, "-o", "binary.bin")
	err = html2text_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	html2text_src_url := "https://github.com/grobian/html2text/releases/download/v2.2.3/html2text-2.2.3.src.tar.gz"
	html2text_cmd_src := exec.Command("curl", "-L", html2text_src_url, "-o", "source.tar.gz")
	err = html2text_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	html2text_cmd_direct := exec.Command("./binary")
	err = html2text_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
