package main

// Hex - Futuristic take on hexdump
// Homepage: https://github.com/sitkevij/hex

import (
	"fmt"
	
	"os/exec"
)

func installHex() {
	// Método 1: Descargar y extraer .tar.gz
	hex_tar_url := "https://github.com/sitkevij/hex/archive/refs/tags/v0.6.0.tar.gz"
	hex_cmd_tar := exec.Command("curl", "-L", hex_tar_url, "-o", "package.tar.gz")
	err := hex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hex_zip_url := "https://github.com/sitkevij/hex/archive/refs/tags/v0.6.0.zip"
	hex_cmd_zip := exec.Command("curl", "-L", hex_zip_url, "-o", "package.zip")
	err = hex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hex_bin_url := "https://github.com/sitkevij/hex/archive/refs/tags/v0.6.0.bin"
	hex_cmd_bin := exec.Command("curl", "-L", hex_bin_url, "-o", "binary.bin")
	err = hex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hex_src_url := "https://github.com/sitkevij/hex/archive/refs/tags/v0.6.0.src.tar.gz"
	hex_cmd_src := exec.Command("curl", "-L", hex_src_url, "-o", "source.tar.gz")
	err = hex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hex_cmd_direct := exec.Command("./binary")
	err = hex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
