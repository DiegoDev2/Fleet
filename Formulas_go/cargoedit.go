package main

// CargoEdit - Utility for managing cargo dependencies from the command-line
// Homepage: https://killercup.github.io/cargo-edit/

import (
	"fmt"
	
	"os/exec"
)

func installCargoEdit() {
	// Método 1: Descargar y extraer .tar.gz
	cargoedit_tar_url := "https://github.com/killercup/cargo-edit/archive/refs/tags/v0.12.2.tar.gz"
	cargoedit_cmd_tar := exec.Command("curl", "-L", cargoedit_tar_url, "-o", "package.tar.gz")
	err := cargoedit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargoedit_zip_url := "https://github.com/killercup/cargo-edit/archive/refs/tags/v0.12.2.zip"
	cargoedit_cmd_zip := exec.Command("curl", "-L", cargoedit_zip_url, "-o", "package.zip")
	err = cargoedit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargoedit_bin_url := "https://github.com/killercup/cargo-edit/archive/refs/tags/v0.12.2.bin"
	cargoedit_cmd_bin := exec.Command("curl", "-L", cargoedit_bin_url, "-o", "binary.bin")
	err = cargoedit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargoedit_src_url := "https://github.com/killercup/cargo-edit/archive/refs/tags/v0.12.2.src.tar.gz"
	cargoedit_cmd_src := exec.Command("curl", "-L", cargoedit_src_url, "-o", "source.tar.gz")
	err = cargoedit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargoedit_cmd_direct := exec.Command("./binary")
	err = cargoedit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: rustup")
exec.Command("latte", "install", "rustup")
	fmt.Println("Instalando dependencia: libgit2@1.6")
exec.Command("latte", "install", "libgit2@1.6")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
