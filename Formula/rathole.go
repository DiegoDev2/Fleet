package main

// Rathole - Reverse proxy for NAT traversal
// Homepage: https://github.com/rapiz1/rathole

import (
	"fmt"
	
	"os/exec"
)

func installRathole() {
	// Método 1: Descargar y extraer .tar.gz
	rathole_tar_url := "https://github.com/rapiz1/rathole/archive/refs/tags/v0.5.0.tar.gz"
	rathole_cmd_tar := exec.Command("curl", "-L", rathole_tar_url, "-o", "package.tar.gz")
	err := rathole_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rathole_zip_url := "https://github.com/rapiz1/rathole/archive/refs/tags/v0.5.0.zip"
	rathole_cmd_zip := exec.Command("curl", "-L", rathole_zip_url, "-o", "package.zip")
	err = rathole_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rathole_bin_url := "https://github.com/rapiz1/rathole/archive/refs/tags/v0.5.0.bin"
	rathole_cmd_bin := exec.Command("curl", "-L", rathole_bin_url, "-o", "binary.bin")
	err = rathole_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rathole_src_url := "https://github.com/rapiz1/rathole/archive/refs/tags/v0.5.0.src.tar.gz"
	rathole_cmd_src := exec.Command("curl", "-L", rathole_src_url, "-o", "source.tar.gz")
	err = rathole_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rathole_cmd_direct := exec.Command("./binary")
	err = rathole_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
