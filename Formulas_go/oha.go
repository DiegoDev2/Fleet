package main

// Oha - HTTP load generator, inspired by rakyll/hey with tui animation
// Homepage: https://github.com/hatoo/oha/

import (
	"fmt"
	
	"os/exec"
)

func installOha() {
	// Método 1: Descargar y extraer .tar.gz
	oha_tar_url := "https://github.com/hatoo/oha/archive/refs/tags/v1.4.6.tar.gz"
	oha_cmd_tar := exec.Command("curl", "-L", oha_tar_url, "-o", "package.tar.gz")
	err := oha_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oha_zip_url := "https://github.com/hatoo/oha/archive/refs/tags/v1.4.6.zip"
	oha_cmd_zip := exec.Command("curl", "-L", oha_zip_url, "-o", "package.zip")
	err = oha_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oha_bin_url := "https://github.com/hatoo/oha/archive/refs/tags/v1.4.6.bin"
	oha_cmd_bin := exec.Command("curl", "-L", oha_bin_url, "-o", "binary.bin")
	err = oha_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oha_src_url := "https://github.com/hatoo/oha/archive/refs/tags/v1.4.6.src.tar.gz"
	oha_cmd_src := exec.Command("curl", "-L", oha_src_url, "-o", "source.tar.gz")
	err = oha_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oha_cmd_direct := exec.Command("./binary")
	err = oha_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
