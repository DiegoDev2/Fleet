package main

// Shairport - Airtunes emulator
// Homepage: https://github.com/abrasive/shairport

import (
	"fmt"
	
	"os/exec"
)

func installShairport() {
	// Método 1: Descargar y extraer .tar.gz
	shairport_tar_url := "https://github.com/abrasive/shairport/archive/refs/tags/1.1.1.tar.gz"
	shairport_cmd_tar := exec.Command("curl", "-L", shairport_tar_url, "-o", "package.tar.gz")
	err := shairport_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shairport_zip_url := "https://github.com/abrasive/shairport/archive/refs/tags/1.1.1.zip"
	shairport_cmd_zip := exec.Command("curl", "-L", shairport_zip_url, "-o", "package.zip")
	err = shairport_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shairport_bin_url := "https://github.com/abrasive/shairport/archive/refs/tags/1.1.1.bin"
	shairport_cmd_bin := exec.Command("curl", "-L", shairport_bin_url, "-o", "binary.bin")
	err = shairport_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shairport_src_url := "https://github.com/abrasive/shairport/archive/refs/tags/1.1.1.src.tar.gz"
	shairport_cmd_src := exec.Command("curl", "-L", shairport_src_url, "-o", "source.tar.gz")
	err = shairport_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shairport_cmd_direct := exec.Command("./binary")
	err = shairport_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
