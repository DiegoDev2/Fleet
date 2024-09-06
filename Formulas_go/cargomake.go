package main

// CargoMake - Rust task runner and build tool
// Homepage: https://github.com/sagiegurari/cargo-make

import (
	"fmt"
	
	"os/exec"
)

func installCargoMake() {
	// Método 1: Descargar y extraer .tar.gz
	cargomake_tar_url := "https://github.com/sagiegurari/cargo-make/archive/refs/tags/0.37.16.tar.gz"
	cargomake_cmd_tar := exec.Command("curl", "-L", cargomake_tar_url, "-o", "package.tar.gz")
	err := cargomake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargomake_zip_url := "https://github.com/sagiegurari/cargo-make/archive/refs/tags/0.37.16.zip"
	cargomake_cmd_zip := exec.Command("curl", "-L", cargomake_zip_url, "-o", "package.zip")
	err = cargomake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargomake_bin_url := "https://github.com/sagiegurari/cargo-make/archive/refs/tags/0.37.16.bin"
	cargomake_cmd_bin := exec.Command("curl", "-L", cargomake_bin_url, "-o", "binary.bin")
	err = cargomake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargomake_src_url := "https://github.com/sagiegurari/cargo-make/archive/refs/tags/0.37.16.src.tar.gz"
	cargomake_cmd_src := exec.Command("curl", "-L", cargomake_src_url, "-o", "source.tar.gz")
	err = cargomake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargomake_cmd_direct := exec.Command("./binary")
	err = cargomake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: rustup")
exec.Command("latte", "install", "rustup")
}
