package main

// RuffLsp - Language Server Protocol implementation for Ruff
// Homepage: https://github.com/astral-sh/ruff-lsp

import (
	"fmt"
	
	"os/exec"
)

func installRuffLsp() {
	// Método 1: Descargar y extraer .tar.gz
	rufflsp_tar_url := "https://files.pythonhosted.org/packages/77/56/a1836adc11c516f75fb7f468b238cdd5d4a248fe9113176b002d09f02ecf/ruff_lsp-0.0.56.tar.gz"
	rufflsp_cmd_tar := exec.Command("curl", "-L", rufflsp_tar_url, "-o", "package.tar.gz")
	err := rufflsp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rufflsp_zip_url := "https://files.pythonhosted.org/packages/77/56/a1836adc11c516f75fb7f468b238cdd5d4a248fe9113176b002d09f02ecf/ruff_lsp-0.0.56.zip"
	rufflsp_cmd_zip := exec.Command("curl", "-L", rufflsp_zip_url, "-o", "package.zip")
	err = rufflsp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rufflsp_bin_url := "https://files.pythonhosted.org/packages/77/56/a1836adc11c516f75fb7f468b238cdd5d4a248fe9113176b002d09f02ecf/ruff_lsp-0.0.56.bin"
	rufflsp_cmd_bin := exec.Command("curl", "-L", rufflsp_bin_url, "-o", "binary.bin")
	err = rufflsp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rufflsp_src_url := "https://files.pythonhosted.org/packages/77/56/a1836adc11c516f75fb7f468b238cdd5d4a248fe9113176b002d09f02ecf/ruff_lsp-0.0.56.src.tar.gz"
	rufflsp_cmd_src := exec.Command("curl", "-L", rufflsp_src_url, "-o", "source.tar.gz")
	err = rufflsp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rufflsp_cmd_direct := exec.Command("./binary")
	err = rufflsp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: ruff")
exec.Command("latte", "install", "ruff")
}
