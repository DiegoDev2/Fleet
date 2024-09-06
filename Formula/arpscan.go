package main

// ArpScan - ARP scanning and fingerprinting tool
// Homepage: https://github.com/royhills/arp-scan

import (
	"fmt"
	
	"os/exec"
)

func installArpScan() {
	// Método 1: Descargar y extraer .tar.gz
	arpscan_tar_url := "https://github.com/royhills/arp-scan/archive/refs/tags/1.10.0.tar.gz"
	arpscan_cmd_tar := exec.Command("curl", "-L", arpscan_tar_url, "-o", "package.tar.gz")
	err := arpscan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	arpscan_zip_url := "https://github.com/royhills/arp-scan/archive/refs/tags/1.10.0.zip"
	arpscan_cmd_zip := exec.Command("curl", "-L", arpscan_zip_url, "-o", "package.zip")
	err = arpscan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	arpscan_bin_url := "https://github.com/royhills/arp-scan/archive/refs/tags/1.10.0.bin"
	arpscan_cmd_bin := exec.Command("curl", "-L", arpscan_bin_url, "-o", "binary.bin")
	err = arpscan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	arpscan_src_url := "https://github.com/royhills/arp-scan/archive/refs/tags/1.10.0.src.tar.gz"
	arpscan_cmd_src := exec.Command("curl", "-L", arpscan_src_url, "-o", "source.tar.gz")
	err = arpscan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	arpscan_cmd_direct := exec.Command("./binary")
	err = arpscan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
