package main

// Modsurfer - Validate, audit and investigate WebAssembly binaries
// Homepage: https://dev.dylibso.com/docs/modsurfer/

import (
	"fmt"
	
	"os/exec"
)

func installModsurfer() {
	// Método 1: Descargar y extraer .tar.gz
	modsurfer_tar_url := "https://github.com/dylibso/modsurfer/archive/refs/tags/v0.0.10.tar.gz"
	modsurfer_cmd_tar := exec.Command("curl", "-L", modsurfer_tar_url, "-o", "package.tar.gz")
	err := modsurfer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	modsurfer_zip_url := "https://github.com/dylibso/modsurfer/archive/refs/tags/v0.0.10.zip"
	modsurfer_cmd_zip := exec.Command("curl", "-L", modsurfer_zip_url, "-o", "package.zip")
	err = modsurfer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	modsurfer_bin_url := "https://github.com/dylibso/modsurfer/archive/refs/tags/v0.0.10.bin"
	modsurfer_cmd_bin := exec.Command("curl", "-L", modsurfer_bin_url, "-o", "binary.bin")
	err = modsurfer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	modsurfer_src_url := "https://github.com/dylibso/modsurfer/archive/refs/tags/v0.0.10.src.tar.gz"
	modsurfer_cmd_src := exec.Command("curl", "-L", modsurfer_src_url, "-o", "source.tar.gz")
	err = modsurfer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	modsurfer_cmd_direct := exec.Command("./binary")
	err = modsurfer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
