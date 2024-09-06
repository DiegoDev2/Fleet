package main

// Bingrep - Greps through binaries from various OSs and architectures
// Homepage: https://github.com/m4b/bingrep

import (
	"fmt"
	
	"os/exec"
)

func installBingrep() {
	// Método 1: Descargar y extraer .tar.gz
	bingrep_tar_url := "https://github.com/m4b/bingrep/archive/refs/tags/v0.11.0.tar.gz"
	bingrep_cmd_tar := exec.Command("curl", "-L", bingrep_tar_url, "-o", "package.tar.gz")
	err := bingrep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bingrep_zip_url := "https://github.com/m4b/bingrep/archive/refs/tags/v0.11.0.zip"
	bingrep_cmd_zip := exec.Command("curl", "-L", bingrep_zip_url, "-o", "package.zip")
	err = bingrep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bingrep_bin_url := "https://github.com/m4b/bingrep/archive/refs/tags/v0.11.0.bin"
	bingrep_cmd_bin := exec.Command("curl", "-L", bingrep_bin_url, "-o", "binary.bin")
	err = bingrep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bingrep_src_url := "https://github.com/m4b/bingrep/archive/refs/tags/v0.11.0.src.tar.gz"
	bingrep_cmd_src := exec.Command("curl", "-L", bingrep_src_url, "-o", "source.tar.gz")
	err = bingrep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bingrep_cmd_direct := exec.Command("./binary")
	err = bingrep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
