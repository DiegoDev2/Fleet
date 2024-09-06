package main

// Ncrack - Network authentication cracking tool
// Homepage: https://nmap.org/ncrack/

import (
	"fmt"
	
	"os/exec"
)

func installNcrack() {
	// Método 1: Descargar y extraer .tar.gz
	ncrack_tar_url := "https://github.com/nmap/ncrack/archive/refs/tags/0.7.tar.gz"
	ncrack_cmd_tar := exec.Command("curl", "-L", ncrack_tar_url, "-o", "package.tar.gz")
	err := ncrack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ncrack_zip_url := "https://github.com/nmap/ncrack/archive/refs/tags/0.7.zip"
	ncrack_cmd_zip := exec.Command("curl", "-L", ncrack_zip_url, "-o", "package.zip")
	err = ncrack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ncrack_bin_url := "https://github.com/nmap/ncrack/archive/refs/tags/0.7.bin"
	ncrack_cmd_bin := exec.Command("curl", "-L", ncrack_bin_url, "-o", "binary.bin")
	err = ncrack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ncrack_src_url := "https://github.com/nmap/ncrack/archive/refs/tags/0.7.src.tar.gz"
	ncrack_cmd_src := exec.Command("curl", "-L", ncrack_src_url, "-o", "source.tar.gz")
	err = ncrack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ncrack_cmd_direct := exec.Command("./binary")
	err = ncrack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
