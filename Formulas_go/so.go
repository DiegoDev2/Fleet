package main

// So - Terminal interface for StackOverflow
// Homepage: https://github.com/samtay/so

import (
	"fmt"
	
	"os/exec"
)

func installSo() {
	// Método 1: Descargar y extraer .tar.gz
	so_tar_url := "https://github.com/samtay/so/archive/refs/tags/v0.4.10.tar.gz"
	so_cmd_tar := exec.Command("curl", "-L", so_tar_url, "-o", "package.tar.gz")
	err := so_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	so_zip_url := "https://github.com/samtay/so/archive/refs/tags/v0.4.10.zip"
	so_cmd_zip := exec.Command("curl", "-L", so_zip_url, "-o", "package.zip")
	err = so_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	so_bin_url := "https://github.com/samtay/so/archive/refs/tags/v0.4.10.bin"
	so_cmd_bin := exec.Command("curl", "-L", so_bin_url, "-o", "binary.bin")
	err = so_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	so_src_url := "https://github.com/samtay/so/archive/refs/tags/v0.4.10.src.tar.gz"
	so_cmd_src := exec.Command("curl", "-L", so_src_url, "-o", "source.tar.gz")
	err = so_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	so_cmd_direct := exec.Command("./binary")
	err = so_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
