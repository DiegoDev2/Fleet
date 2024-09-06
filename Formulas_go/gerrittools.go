package main

// GerritTools - Tools to ease Gerrit code review
// Homepage: https://github.com/indirect/gerrit-tools

import (
	"fmt"
	
	"os/exec"
)

func installGerritTools() {
	// Método 1: Descargar y extraer .tar.gz
	gerrittools_tar_url := "https://github.com/indirect/gerrit-tools/archive/refs/tags/v1.0.0.tar.gz"
	gerrittools_cmd_tar := exec.Command("curl", "-L", gerrittools_tar_url, "-o", "package.tar.gz")
	err := gerrittools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gerrittools_zip_url := "https://github.com/indirect/gerrit-tools/archive/refs/tags/v1.0.0.zip"
	gerrittools_cmd_zip := exec.Command("curl", "-L", gerrittools_zip_url, "-o", "package.zip")
	err = gerrittools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gerrittools_bin_url := "https://github.com/indirect/gerrit-tools/archive/refs/tags/v1.0.0.bin"
	gerrittools_cmd_bin := exec.Command("curl", "-L", gerrittools_bin_url, "-o", "binary.bin")
	err = gerrittools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gerrittools_src_url := "https://github.com/indirect/gerrit-tools/archive/refs/tags/v1.0.0.src.tar.gz"
	gerrittools_cmd_src := exec.Command("curl", "-L", gerrittools_src_url, "-o", "source.tar.gz")
	err = gerrittools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gerrittools_cmd_direct := exec.Command("./binary")
	err = gerrittools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
