package main

// CargoDocset - Cargo subcommand to generate a Dash/Zeal docset for your Rust packages
// Homepage: https://github.com/Robzz/cargo-docset

import (
	"fmt"
	
	"os/exec"
)

func installCargoDocset() {
	// Método 1: Descargar y extraer .tar.gz
	cargodocset_tar_url := "https://github.com/Robzz/cargo-docset/archive/refs/tags/v0.3.1.tar.gz"
	cargodocset_cmd_tar := exec.Command("curl", "-L", cargodocset_tar_url, "-o", "package.tar.gz")
	err := cargodocset_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargodocset_zip_url := "https://github.com/Robzz/cargo-docset/archive/refs/tags/v0.3.1.zip"
	cargodocset_cmd_zip := exec.Command("curl", "-L", cargodocset_zip_url, "-o", "package.zip")
	err = cargodocset_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargodocset_bin_url := "https://github.com/Robzz/cargo-docset/archive/refs/tags/v0.3.1.bin"
	cargodocset_cmd_bin := exec.Command("curl", "-L", cargodocset_bin_url, "-o", "binary.bin")
	err = cargodocset_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargodocset_src_url := "https://github.com/Robzz/cargo-docset/archive/refs/tags/v0.3.1.src.tar.gz"
	cargodocset_cmd_src := exec.Command("curl", "-L", cargodocset_src_url, "-o", "source.tar.gz")
	err = cargodocset_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargodocset_cmd_direct := exec.Command("./binary")
	err = cargodocset_cmd_direct.Run()
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
