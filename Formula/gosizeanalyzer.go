package main

// GoSizeAnalyzer - Analyzing the dependencies in compiled Golang binaries
// Homepage: https://github.com/Zxilly/go-size-analyzer

import (
	"fmt"
	
	"os/exec"
)

func installGoSizeAnalyzer() {
	// Método 1: Descargar y extraer .tar.gz
	gosizeanalyzer_tar_url := "https://github.com/Zxilly/go-size-analyzer/archive/refs/tags/v1.7.2.tar.gz"
	gosizeanalyzer_cmd_tar := exec.Command("curl", "-L", gosizeanalyzer_tar_url, "-o", "package.tar.gz")
	err := gosizeanalyzer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gosizeanalyzer_zip_url := "https://github.com/Zxilly/go-size-analyzer/archive/refs/tags/v1.7.2.zip"
	gosizeanalyzer_cmd_zip := exec.Command("curl", "-L", gosizeanalyzer_zip_url, "-o", "package.zip")
	err = gosizeanalyzer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gosizeanalyzer_bin_url := "https://github.com/Zxilly/go-size-analyzer/archive/refs/tags/v1.7.2.bin"
	gosizeanalyzer_cmd_bin := exec.Command("curl", "-L", gosizeanalyzer_bin_url, "-o", "binary.bin")
	err = gosizeanalyzer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gosizeanalyzer_src_url := "https://github.com/Zxilly/go-size-analyzer/archive/refs/tags/v1.7.2.src.tar.gz"
	gosizeanalyzer_cmd_src := exec.Command("curl", "-L", gosizeanalyzer_src_url, "-o", "source.tar.gz")
	err = gosizeanalyzer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gosizeanalyzer_cmd_direct := exec.Command("./binary")
	err = gosizeanalyzer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: pnpm")
	exec.Command("latte", "install", "pnpm").Run()
}
