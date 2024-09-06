package main

// Subfinder - Subdomain discovery tool
// Homepage: https://github.com/projectdiscovery/subfinder

import (
	"fmt"
	
	"os/exec"
)

func installSubfinder() {
	// Método 1: Descargar y extraer .tar.gz
	subfinder_tar_url := "https://github.com/projectdiscovery/subfinder/archive/refs/tags/v2.6.6.tar.gz"
	subfinder_cmd_tar := exec.Command("curl", "-L", subfinder_tar_url, "-o", "package.tar.gz")
	err := subfinder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	subfinder_zip_url := "https://github.com/projectdiscovery/subfinder/archive/refs/tags/v2.6.6.zip"
	subfinder_cmd_zip := exec.Command("curl", "-L", subfinder_zip_url, "-o", "package.zip")
	err = subfinder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	subfinder_bin_url := "https://github.com/projectdiscovery/subfinder/archive/refs/tags/v2.6.6.bin"
	subfinder_cmd_bin := exec.Command("curl", "-L", subfinder_bin_url, "-o", "binary.bin")
	err = subfinder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	subfinder_src_url := "https://github.com/projectdiscovery/subfinder/archive/refs/tags/v2.6.6.src.tar.gz"
	subfinder_cmd_src := exec.Command("curl", "-L", subfinder_src_url, "-o", "source.tar.gz")
	err = subfinder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	subfinder_cmd_direct := exec.Command("./binary")
	err = subfinder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
