package main

// Rustscan - Modern Day Portscanner
// Homepage: https://github.com/rustscan/rustscan

import (
	"fmt"
	
	"os/exec"
)

func installRustscan() {
	// Método 1: Descargar y extraer .tar.gz
	rustscan_tar_url := "https://github.com/RustScan/RustScan/archive/refs/tags/2.3.0.tar.gz"
	rustscan_cmd_tar := exec.Command("curl", "-L", rustscan_tar_url, "-o", "package.tar.gz")
	err := rustscan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rustscan_zip_url := "https://github.com/RustScan/RustScan/archive/refs/tags/2.3.0.zip"
	rustscan_cmd_zip := exec.Command("curl", "-L", rustscan_zip_url, "-o", "package.zip")
	err = rustscan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rustscan_bin_url := "https://github.com/RustScan/RustScan/archive/refs/tags/2.3.0.bin"
	rustscan_cmd_bin := exec.Command("curl", "-L", rustscan_bin_url, "-o", "binary.bin")
	err = rustscan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rustscan_src_url := "https://github.com/RustScan/RustScan/archive/refs/tags/2.3.0.src.tar.gz"
	rustscan_cmd_src := exec.Command("curl", "-L", rustscan_src_url, "-o", "source.tar.gz")
	err = rustscan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rustscan_cmd_direct := exec.Command("./binary")
	err = rustscan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: nmap")
exec.Command("latte", "install", "nmap")
}
