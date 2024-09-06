package main

// Texlab - Implementation of the Language Server Protocol for LaTeX
// Homepage: https://github.com/latex-lsp/texlab/

import (
	"fmt"
	
	"os/exec"
)

func installTexlab() {
	// Método 1: Descargar y extraer .tar.gz
	texlab_tar_url := "https://github.com/latex-lsp/texlab/archive/refs/tags/v5.19.0.tar.gz"
	texlab_cmd_tar := exec.Command("curl", "-L", texlab_tar_url, "-o", "package.tar.gz")
	err := texlab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	texlab_zip_url := "https://github.com/latex-lsp/texlab/archive/refs/tags/v5.19.0.zip"
	texlab_cmd_zip := exec.Command("curl", "-L", texlab_zip_url, "-o", "package.zip")
	err = texlab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	texlab_bin_url := "https://github.com/latex-lsp/texlab/archive/refs/tags/v5.19.0.bin"
	texlab_cmd_bin := exec.Command("curl", "-L", texlab_bin_url, "-o", "binary.bin")
	err = texlab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	texlab_src_url := "https://github.com/latex-lsp/texlab/archive/refs/tags/v5.19.0.src.tar.gz"
	texlab_cmd_src := exec.Command("curl", "-L", texlab_src_url, "-o", "source.tar.gz")
	err = texlab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	texlab_cmd_direct := exec.Command("./binary")
	err = texlab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
