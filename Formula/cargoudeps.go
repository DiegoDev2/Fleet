package main

// CargoUdeps - Find unused dependencies in Cargo.toml
// Homepage: https://github.com/est31/cargo-udeps

import (
	"fmt"
	
	"os/exec"
)

func installCargoUdeps() {
	// Método 1: Descargar y extraer .tar.gz
	cargoudeps_tar_url := "https://github.com/est31/cargo-udeps/archive/refs/tags/v0.1.50.tar.gz"
	cargoudeps_cmd_tar := exec.Command("curl", "-L", cargoudeps_tar_url, "-o", "package.tar.gz")
	err := cargoudeps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargoudeps_zip_url := "https://github.com/est31/cargo-udeps/archive/refs/tags/v0.1.50.zip"
	cargoudeps_cmd_zip := exec.Command("curl", "-L", cargoudeps_zip_url, "-o", "package.zip")
	err = cargoudeps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargoudeps_bin_url := "https://github.com/est31/cargo-udeps/archive/refs/tags/v0.1.50.bin"
	cargoudeps_cmd_bin := exec.Command("curl", "-L", cargoudeps_bin_url, "-o", "binary.bin")
	err = cargoudeps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargoudeps_src_url := "https://github.com/est31/cargo-udeps/archive/refs/tags/v0.1.50.src.tar.gz"
	cargoudeps_cmd_src := exec.Command("curl", "-L", cargoudeps_src_url, "-o", "source.tar.gz")
	err = cargoudeps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargoudeps_cmd_direct := exec.Command("./binary")
	err = cargoudeps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
	fmt.Println("Instalando dependencia: libgit2@1.7")
	exec.Command("latte", "install", "libgit2@1.7").Run()
	fmt.Println("Instalando dependencia: libssh2")
	exec.Command("latte", "install", "libssh2").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
