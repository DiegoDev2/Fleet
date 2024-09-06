package main

// Vultr - Command-line tool for Vultr services
// Homepage: https://github.com/vultr/vultr-cli

import (
	"fmt"
	
	"os/exec"
)

func installVultr() {
	// Método 1: Descargar y extraer .tar.gz
	vultr_tar_url := "https://github.com/vultr/vultr-cli/archive/refs/tags/v3.3.1.tar.gz"
	vultr_cmd_tar := exec.Command("curl", "-L", vultr_tar_url, "-o", "package.tar.gz")
	err := vultr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vultr_zip_url := "https://github.com/vultr/vultr-cli/archive/refs/tags/v3.3.1.zip"
	vultr_cmd_zip := exec.Command("curl", "-L", vultr_zip_url, "-o", "package.zip")
	err = vultr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vultr_bin_url := "https://github.com/vultr/vultr-cli/archive/refs/tags/v3.3.1.bin"
	vultr_cmd_bin := exec.Command("curl", "-L", vultr_bin_url, "-o", "binary.bin")
	err = vultr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vultr_src_url := "https://github.com/vultr/vultr-cli/archive/refs/tags/v3.3.1.src.tar.gz"
	vultr_cmd_src := exec.Command("curl", "-L", vultr_src_url, "-o", "source.tar.gz")
	err = vultr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vultr_cmd_direct := exec.Command("./binary")
	err = vultr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
