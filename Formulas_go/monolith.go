package main

// Monolith - CLI tool for saving complete web pages as a single HTML file
// Homepage: https://github.com/Y2Z/monolith

import (
	"fmt"
	
	"os/exec"
)

func installMonolith() {
	// Método 1: Descargar y extraer .tar.gz
	monolith_tar_url := "https://github.com/Y2Z/monolith/archive/refs/tags/v2.8.2.tar.gz"
	monolith_cmd_tar := exec.Command("curl", "-L", monolith_tar_url, "-o", "package.tar.gz")
	err := monolith_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	monolith_zip_url := "https://github.com/Y2Z/monolith/archive/refs/tags/v2.8.2.zip"
	monolith_cmd_zip := exec.Command("curl", "-L", monolith_zip_url, "-o", "package.zip")
	err = monolith_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	monolith_bin_url := "https://github.com/Y2Z/monolith/archive/refs/tags/v2.8.2.bin"
	monolith_cmd_bin := exec.Command("curl", "-L", monolith_bin_url, "-o", "binary.bin")
	err = monolith_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	monolith_src_url := "https://github.com/Y2Z/monolith/archive/refs/tags/v2.8.2.src.tar.gz"
	monolith_cmd_src := exec.Command("curl", "-L", monolith_src_url, "-o", "source.tar.gz")
	err = monolith_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	monolith_cmd_direct := exec.Command("./binary")
	err = monolith_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
