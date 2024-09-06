package main

// ClojureLsp - Language Server (LSP) for Clojure
// Homepage: https://github.com/clojure-lsp/clojure-lsp

import (
	"fmt"
	
	"os/exec"
)

func installClojureLsp() {
	// Método 1: Descargar y extraer .tar.gz
	clojurelsp_tar_url := "https://github.com/clojure-lsp/clojure-lsp/releases/download/2024.04.22-11.50.26/clojure-lsp-standalone.jar"
	clojurelsp_cmd_tar := exec.Command("curl", "-L", clojurelsp_tar_url, "-o", "package.tar.gz")
	err := clojurelsp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clojurelsp_zip_url := "https://github.com/clojure-lsp/clojure-lsp/releases/download/2024.04.22-11.50.26/clojure-lsp-standalone.jar"
	clojurelsp_cmd_zip := exec.Command("curl", "-L", clojurelsp_zip_url, "-o", "package.zip")
	err = clojurelsp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clojurelsp_bin_url := "https://github.com/clojure-lsp/clojure-lsp/releases/download/2024.04.22-11.50.26/clojure-lsp-standalone.jar"
	clojurelsp_cmd_bin := exec.Command("curl", "-L", clojurelsp_bin_url, "-o", "binary.bin")
	err = clojurelsp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clojurelsp_src_url := "https://github.com/clojure-lsp/clojure-lsp/releases/download/2024.04.22-11.50.26/clojure-lsp-standalone.jar"
	clojurelsp_cmd_src := exec.Command("curl", "-L", clojurelsp_src_url, "-o", "source.tar.gz")
	err = clojurelsp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clojurelsp_cmd_direct := exec.Command("./binary")
	err = clojurelsp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
