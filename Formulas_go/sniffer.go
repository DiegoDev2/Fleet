package main

// Sniffer - Modern alternative network traffic sniffer
// Homepage: https://github.com/chenjiandongx/sniffer

import (
	"fmt"
	
	"os/exec"
)

func installSniffer() {
	// Método 1: Descargar y extraer .tar.gz
	sniffer_tar_url := "https://github.com/chenjiandongx/sniffer/archive/refs/tags/v0.6.2.tar.gz"
	sniffer_cmd_tar := exec.Command("curl", "-L", sniffer_tar_url, "-o", "package.tar.gz")
	err := sniffer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sniffer_zip_url := "https://github.com/chenjiandongx/sniffer/archive/refs/tags/v0.6.2.zip"
	sniffer_cmd_zip := exec.Command("curl", "-L", sniffer_zip_url, "-o", "package.zip")
	err = sniffer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sniffer_bin_url := "https://github.com/chenjiandongx/sniffer/archive/refs/tags/v0.6.2.bin"
	sniffer_cmd_bin := exec.Command("curl", "-L", sniffer_bin_url, "-o", "binary.bin")
	err = sniffer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sniffer_src_url := "https://github.com/chenjiandongx/sniffer/archive/refs/tags/v0.6.2.src.tar.gz"
	sniffer_cmd_src := exec.Command("curl", "-L", sniffer_src_url, "-o", "source.tar.gz")
	err = sniffer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sniffer_cmd_direct := exec.Command("./binary")
	err = sniffer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
