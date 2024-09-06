package main

// Sslscan - Test SSL/TLS enabled services to discover supported cipher suites
// Homepage: https://github.com/rbsec/sslscan

import (
	"fmt"
	
	"os/exec"
)

func installSslscan() {
	// Método 1: Descargar y extraer .tar.gz
	sslscan_tar_url := "https://github.com/rbsec/sslscan/archive/refs/tags/2.1.4.tar.gz"
	sslscan_cmd_tar := exec.Command("curl", "-L", sslscan_tar_url, "-o", "package.tar.gz")
	err := sslscan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sslscan_zip_url := "https://github.com/rbsec/sslscan/archive/refs/tags/2.1.4.zip"
	sslscan_cmd_zip := exec.Command("curl", "-L", sslscan_zip_url, "-o", "package.zip")
	err = sslscan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sslscan_bin_url := "https://github.com/rbsec/sslscan/archive/refs/tags/2.1.4.bin"
	sslscan_cmd_bin := exec.Command("curl", "-L", sslscan_bin_url, "-o", "binary.bin")
	err = sslscan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sslscan_src_url := "https://github.com/rbsec/sslscan/archive/refs/tags/2.1.4.src.tar.gz"
	sslscan_cmd_src := exec.Command("curl", "-L", sslscan_src_url, "-o", "source.tar.gz")
	err = sslscan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sslscan_cmd_direct := exec.Command("./binary")
	err = sslscan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
