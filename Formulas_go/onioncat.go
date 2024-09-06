package main

// Onioncat - VPN-adapter that provides location privacy using Tor or I2P
// Homepage: https://github.com/rahra/onioncat

import (
	"fmt"
	
	"os/exec"
)

func installOnioncat() {
	// Método 1: Descargar y extraer .tar.gz
	onioncat_tar_url := "https://github.com/rahra/onioncat/archive/refs/tags/v4.11.0.tar.gz"
	onioncat_cmd_tar := exec.Command("curl", "-L", onioncat_tar_url, "-o", "package.tar.gz")
	err := onioncat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	onioncat_zip_url := "https://github.com/rahra/onioncat/archive/refs/tags/v4.11.0.zip"
	onioncat_cmd_zip := exec.Command("curl", "-L", onioncat_zip_url, "-o", "package.zip")
	err = onioncat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	onioncat_bin_url := "https://github.com/rahra/onioncat/archive/refs/tags/v4.11.0.bin"
	onioncat_cmd_bin := exec.Command("curl", "-L", onioncat_bin_url, "-o", "binary.bin")
	err = onioncat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	onioncat_src_url := "https://github.com/rahra/onioncat/archive/refs/tags/v4.11.0.src.tar.gz"
	onioncat_cmd_src := exec.Command("curl", "-L", onioncat_src_url, "-o", "source.tar.gz")
	err = onioncat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	onioncat_cmd_direct := exec.Command("./binary")
	err = onioncat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: tor")
exec.Command("latte", "install", "tor")
}
