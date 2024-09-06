package main

// Octobuild - Compiler cache for Unreal Engine
// Homepage: https://github.com/octobuild/octobuild

import (
	"fmt"
	
	"os/exec"
)

func installOctobuild() {
	// Método 1: Descargar y extraer .tar.gz
	octobuild_tar_url := "https://github.com/octobuild/octobuild/archive/refs/tags/1.4.0.tar.gz"
	octobuild_cmd_tar := exec.Command("curl", "-L", octobuild_tar_url, "-o", "package.tar.gz")
	err := octobuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	octobuild_zip_url := "https://github.com/octobuild/octobuild/archive/refs/tags/1.4.0.zip"
	octobuild_cmd_zip := exec.Command("curl", "-L", octobuild_zip_url, "-o", "package.zip")
	err = octobuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	octobuild_bin_url := "https://github.com/octobuild/octobuild/archive/refs/tags/1.4.0.bin"
	octobuild_cmd_bin := exec.Command("curl", "-L", octobuild_bin_url, "-o", "binary.bin")
	err = octobuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	octobuild_src_url := "https://github.com/octobuild/octobuild/archive/refs/tags/1.4.0.src.tar.gz"
	octobuild_cmd_src := exec.Command("curl", "-L", octobuild_src_url, "-o", "source.tar.gz")
	err = octobuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	octobuild_cmd_direct := exec.Command("./binary")
	err = octobuild_cmd_direct.Run()
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
