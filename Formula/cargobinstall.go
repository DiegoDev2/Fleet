package main

// CargoBinstall - Binary installation for rust projects
// Homepage: https://github.com/cargo-bins/cargo-binstall

import (
	"fmt"
	
	"os/exec"
)

func installCargoBinstall() {
	// Método 1: Descargar y extraer .tar.gz
	cargobinstall_tar_url := "https://github.com/cargo-bins/cargo-binstall/archive/refs/tags/v1.10.3.tar.gz"
	cargobinstall_cmd_tar := exec.Command("curl", "-L", cargobinstall_tar_url, "-o", "package.tar.gz")
	err := cargobinstall_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargobinstall_zip_url := "https://github.com/cargo-bins/cargo-binstall/archive/refs/tags/v1.10.3.zip"
	cargobinstall_cmd_zip := exec.Command("curl", "-L", cargobinstall_zip_url, "-o", "package.zip")
	err = cargobinstall_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargobinstall_bin_url := "https://github.com/cargo-bins/cargo-binstall/archive/refs/tags/v1.10.3.bin"
	cargobinstall_cmd_bin := exec.Command("curl", "-L", cargobinstall_bin_url, "-o", "binary.bin")
	err = cargobinstall_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargobinstall_src_url := "https://github.com/cargo-bins/cargo-binstall/archive/refs/tags/v1.10.3.src.tar.gz"
	cargobinstall_cmd_src := exec.Command("curl", "-L", cargobinstall_src_url, "-o", "source.tar.gz")
	err = cargobinstall_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargobinstall_cmd_direct := exec.Command("./binary")
	err = cargobinstall_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
