package main

// AdrViewer - Generate easy-to-read web pages for your Architecture Decision Records
// Homepage: https://github.com/mrwilson/adr-viewer

import (
	"fmt"
	
	"os/exec"
)

func installAdrViewer() {
	// Método 1: Descargar y extraer .tar.gz
	adrviewer_tar_url := "https://files.pythonhosted.org/packages/1b/72/0f787da38d0f9d69c06b31d8f412735ed4fad383edd7f7d2286f4fc7b5b0/adr_viewer-1.4.0.tar.gz"
	adrviewer_cmd_tar := exec.Command("curl", "-L", adrviewer_tar_url, "-o", "package.tar.gz")
	err := adrviewer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	adrviewer_zip_url := "https://files.pythonhosted.org/packages/1b/72/0f787da38d0f9d69c06b31d8f412735ed4fad383edd7f7d2286f4fc7b5b0/adr_viewer-1.4.0.zip"
	adrviewer_cmd_zip := exec.Command("curl", "-L", adrviewer_zip_url, "-o", "package.zip")
	err = adrviewer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	adrviewer_bin_url := "https://files.pythonhosted.org/packages/1b/72/0f787da38d0f9d69c06b31d8f412735ed4fad383edd7f7d2286f4fc7b5b0/adr_viewer-1.4.0.bin"
	adrviewer_cmd_bin := exec.Command("curl", "-L", adrviewer_bin_url, "-o", "binary.bin")
	err = adrviewer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	adrviewer_src_url := "https://files.pythonhosted.org/packages/1b/72/0f787da38d0f9d69c06b31d8f412735ed4fad383edd7f7d2286f4fc7b5b0/adr_viewer-1.4.0.src.tar.gz"
	adrviewer_cmd_src := exec.Command("curl", "-L", adrviewer_src_url, "-o", "source.tar.gz")
	err = adrviewer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	adrviewer_cmd_direct := exec.Command("./binary")
	err = adrviewer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
