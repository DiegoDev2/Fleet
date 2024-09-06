package main

// SsllabsScan - This tool is a command-line client for the SSL Labs APIs
// Homepage: https://github.com/ssllabs/ssllabs-scan/

import (
	"fmt"
	
	"os/exec"
)

func installSsllabsScan() {
	// Método 1: Descargar y extraer .tar.gz
	ssllabsscan_tar_url := "https://github.com/ssllabs/ssllabs-scan/archive/refs/tags/v1.5.0.tar.gz"
	ssllabsscan_cmd_tar := exec.Command("curl", "-L", ssllabsscan_tar_url, "-o", "package.tar.gz")
	err := ssllabsscan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ssllabsscan_zip_url := "https://github.com/ssllabs/ssllabs-scan/archive/refs/tags/v1.5.0.zip"
	ssllabsscan_cmd_zip := exec.Command("curl", "-L", ssllabsscan_zip_url, "-o", "package.zip")
	err = ssllabsscan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ssllabsscan_bin_url := "https://github.com/ssllabs/ssllabs-scan/archive/refs/tags/v1.5.0.bin"
	ssllabsscan_cmd_bin := exec.Command("curl", "-L", ssllabsscan_bin_url, "-o", "binary.bin")
	err = ssllabsscan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ssllabsscan_src_url := "https://github.com/ssllabs/ssllabs-scan/archive/refs/tags/v1.5.0.src.tar.gz"
	ssllabsscan_cmd_src := exec.Command("curl", "-L", ssllabsscan_src_url, "-o", "source.tar.gz")
	err = ssllabsscan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ssllabsscan_cmd_direct := exec.Command("./binary")
	err = ssllabsscan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
