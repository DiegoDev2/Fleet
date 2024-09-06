package main

// CargoWatch - Watches over your Cargo project's source
// Homepage: https://watchexec.github.io/#cargo-watch

import (
	"fmt"
	
	"os/exec"
)

func installCargoWatch() {
	// Método 1: Descargar y extraer .tar.gz
	cargowatch_tar_url := "https://github.com/watchexec/cargo-watch/archive/refs/tags/v8.5.2.tar.gz"
	cargowatch_cmd_tar := exec.Command("curl", "-L", cargowatch_tar_url, "-o", "package.tar.gz")
	err := cargowatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargowatch_zip_url := "https://github.com/watchexec/cargo-watch/archive/refs/tags/v8.5.2.zip"
	cargowatch_cmd_zip := exec.Command("curl", "-L", cargowatch_zip_url, "-o", "package.zip")
	err = cargowatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargowatch_bin_url := "https://github.com/watchexec/cargo-watch/archive/refs/tags/v8.5.2.bin"
	cargowatch_cmd_bin := exec.Command("curl", "-L", cargowatch_bin_url, "-o", "binary.bin")
	err = cargowatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargowatch_src_url := "https://github.com/watchexec/cargo-watch/archive/refs/tags/v8.5.2.src.tar.gz"
	cargowatch_cmd_src := exec.Command("curl", "-L", cargowatch_src_url, "-o", "source.tar.gz")
	err = cargowatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargowatch_cmd_direct := exec.Command("./binary")
	err = cargowatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
}
