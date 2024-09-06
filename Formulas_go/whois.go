package main

// Whois - Lookup tool for domain names and other internet resources
// Homepage: https://github.com/rfc1036/whois

import (
	"fmt"
	
	"os/exec"
)

func installWhois() {
	// Método 1: Descargar y extraer .tar.gz
	whois_tar_url := "https://github.com/rfc1036/whois/archive/refs/tags/v5.5.23.tar.gz"
	whois_cmd_tar := exec.Command("curl", "-L", whois_tar_url, "-o", "package.tar.gz")
	err := whois_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	whois_zip_url := "https://github.com/rfc1036/whois/archive/refs/tags/v5.5.23.zip"
	whois_cmd_zip := exec.Command("curl", "-L", whois_zip_url, "-o", "package.zip")
	err = whois_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	whois_bin_url := "https://github.com/rfc1036/whois/archive/refs/tags/v5.5.23.bin"
	whois_cmd_bin := exec.Command("curl", "-L", whois_bin_url, "-o", "binary.bin")
	err = whois_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	whois_src_url := "https://github.com/rfc1036/whois/archive/refs/tags/v5.5.23.src.tar.gz"
	whois_cmd_src := exec.Command("curl", "-L", whois_src_url, "-o", "source.tar.gz")
	err = whois_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	whois_cmd_direct := exec.Command("./binary")
	err = whois_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libidn2")
exec.Command("latte", "install", "libidn2")
}
