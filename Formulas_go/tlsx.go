package main

// Tlsx - Fast and configurable TLS grabber focused on TLS based data collection
// Homepage: https://github.com/projectdiscovery/tlsx

import (
	"fmt"
	
	"os/exec"
)

func installTlsx() {
	// Método 1: Descargar y extraer .tar.gz
	tlsx_tar_url := "https://github.com/projectdiscovery/tlsx/archive/refs/tags/v1.1.7.tar.gz"
	tlsx_cmd_tar := exec.Command("curl", "-L", tlsx_tar_url, "-o", "package.tar.gz")
	err := tlsx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tlsx_zip_url := "https://github.com/projectdiscovery/tlsx/archive/refs/tags/v1.1.7.zip"
	tlsx_cmd_zip := exec.Command("curl", "-L", tlsx_zip_url, "-o", "package.zip")
	err = tlsx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tlsx_bin_url := "https://github.com/projectdiscovery/tlsx/archive/refs/tags/v1.1.7.bin"
	tlsx_cmd_bin := exec.Command("curl", "-L", tlsx_bin_url, "-o", "binary.bin")
	err = tlsx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tlsx_src_url := "https://github.com/projectdiscovery/tlsx/archive/refs/tags/v1.1.7.src.tar.gz"
	tlsx_cmd_src := exec.Command("curl", "-L", tlsx_src_url, "-o", "source.tar.gz")
	err = tlsx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tlsx_cmd_direct := exec.Command("./binary")
	err = tlsx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
