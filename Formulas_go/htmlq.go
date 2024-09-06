package main

// Htmlq - Uses CSS selectors to extract bits content from HTML files
// Homepage: https://github.com/mgdm/htmlq

import (
	"fmt"
	
	"os/exec"
)

func installHtmlq() {
	// Método 1: Descargar y extraer .tar.gz
	htmlq_tar_url := "https://github.com/mgdm/htmlq/archive/refs/tags/v0.4.0.tar.gz"
	htmlq_cmd_tar := exec.Command("curl", "-L", htmlq_tar_url, "-o", "package.tar.gz")
	err := htmlq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	htmlq_zip_url := "https://github.com/mgdm/htmlq/archive/refs/tags/v0.4.0.zip"
	htmlq_cmd_zip := exec.Command("curl", "-L", htmlq_zip_url, "-o", "package.zip")
	err = htmlq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	htmlq_bin_url := "https://github.com/mgdm/htmlq/archive/refs/tags/v0.4.0.bin"
	htmlq_cmd_bin := exec.Command("curl", "-L", htmlq_bin_url, "-o", "binary.bin")
	err = htmlq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	htmlq_src_url := "https://github.com/mgdm/htmlq/archive/refs/tags/v0.4.0.src.tar.gz"
	htmlq_cmd_src := exec.Command("curl", "-L", htmlq_src_url, "-o", "source.tar.gz")
	err = htmlq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	htmlq_cmd_direct := exec.Command("./binary")
	err = htmlq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
