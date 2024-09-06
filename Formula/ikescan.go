package main

// IkeScan - Discover and fingerprint IKE hosts
// Homepage: https://github.com/royhills/ike-scan

import (
	"fmt"
	
	"os/exec"
)

func installIkeScan() {
	// Método 1: Descargar y extraer .tar.gz
	ikescan_tar_url := "https://github.com/royhills/ike-scan/archive/refs/tags/1.9.5.tar.gz"
	ikescan_cmd_tar := exec.Command("curl", "-L", ikescan_tar_url, "-o", "package.tar.gz")
	err := ikescan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ikescan_zip_url := "https://github.com/royhills/ike-scan/archive/refs/tags/1.9.5.zip"
	ikescan_cmd_zip := exec.Command("curl", "-L", ikescan_zip_url, "-o", "package.zip")
	err = ikescan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ikescan_bin_url := "https://github.com/royhills/ike-scan/archive/refs/tags/1.9.5.bin"
	ikescan_cmd_bin := exec.Command("curl", "-L", ikescan_bin_url, "-o", "binary.bin")
	err = ikescan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ikescan_src_url := "https://github.com/royhills/ike-scan/archive/refs/tags/1.9.5.src.tar.gz"
	ikescan_cmd_src := exec.Command("curl", "-L", ikescan_src_url, "-o", "source.tar.gz")
	err = ikescan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ikescan_cmd_direct := exec.Command("./binary")
	err = ikescan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
