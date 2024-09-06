package main

// Terrascan - Detect compliance and security violations across Infrastructure as Code
// Homepage: https://runterrascan.io/

import (
	"fmt"
	
	"os/exec"
)

func installTerrascan() {
	// Método 1: Descargar y extraer .tar.gz
	terrascan_tar_url := "https://github.com/tenable/terrascan/archive/refs/tags/v1.19.3.tar.gz"
	terrascan_cmd_tar := exec.Command("curl", "-L", terrascan_tar_url, "-o", "package.tar.gz")
	err := terrascan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	terrascan_zip_url := "https://github.com/tenable/terrascan/archive/refs/tags/v1.19.3.zip"
	terrascan_cmd_zip := exec.Command("curl", "-L", terrascan_zip_url, "-o", "package.zip")
	err = terrascan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	terrascan_bin_url := "https://github.com/tenable/terrascan/archive/refs/tags/v1.19.3.bin"
	terrascan_cmd_bin := exec.Command("curl", "-L", terrascan_bin_url, "-o", "binary.bin")
	err = terrascan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	terrascan_src_url := "https://github.com/tenable/terrascan/archive/refs/tags/v1.19.3.src.tar.gz"
	terrascan_cmd_src := exec.Command("curl", "-L", terrascan_src_url, "-o", "source.tar.gz")
	err = terrascan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	terrascan_cmd_direct := exec.Command("./binary")
	err = terrascan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
