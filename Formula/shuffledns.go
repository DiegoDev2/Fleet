package main

// Shuffledns - Enumerate subdomains using active bruteforce & resolve subdomains with wildcards
// Homepage: https://github.com/projectdiscovery/shuffledns

import (
	"fmt"
	
	"os/exec"
)

func installShuffledns() {
	// Método 1: Descargar y extraer .tar.gz
	shuffledns_tar_url := "https://github.com/projectdiscovery/shuffledns/archive/refs/tags/v1.1.0.tar.gz"
	shuffledns_cmd_tar := exec.Command("curl", "-L", shuffledns_tar_url, "-o", "package.tar.gz")
	err := shuffledns_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shuffledns_zip_url := "https://github.com/projectdiscovery/shuffledns/archive/refs/tags/v1.1.0.zip"
	shuffledns_cmd_zip := exec.Command("curl", "-L", shuffledns_zip_url, "-o", "package.zip")
	err = shuffledns_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shuffledns_bin_url := "https://github.com/projectdiscovery/shuffledns/archive/refs/tags/v1.1.0.bin"
	shuffledns_cmd_bin := exec.Command("curl", "-L", shuffledns_bin_url, "-o", "binary.bin")
	err = shuffledns_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shuffledns_src_url := "https://github.com/projectdiscovery/shuffledns/archive/refs/tags/v1.1.0.src.tar.gz"
	shuffledns_cmd_src := exec.Command("curl", "-L", shuffledns_src_url, "-o", "source.tar.gz")
	err = shuffledns_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shuffledns_cmd_direct := exec.Command("./binary")
	err = shuffledns_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
