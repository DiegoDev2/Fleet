package main

// CargoDeps - Cargo subcommand to building dependency graphs of Rust projects
// Homepage: https://crates.io/crates/cargo-deps

import (
	"fmt"
	
	"os/exec"
)

func installCargoDeps() {
	// Método 1: Descargar y extraer .tar.gz
	cargodeps_tar_url := "https://static.crates.io/crates/cargo-deps/cargo-deps-1.5.1.crate"
	cargodeps_cmd_tar := exec.Command("curl", "-L", cargodeps_tar_url, "-o", "package.tar.gz")
	err := cargodeps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargodeps_zip_url := "https://static.crates.io/crates/cargo-deps/cargo-deps-1.5.1.crate"
	cargodeps_cmd_zip := exec.Command("curl", "-L", cargodeps_zip_url, "-o", "package.zip")
	err = cargodeps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargodeps_bin_url := "https://static.crates.io/crates/cargo-deps/cargo-deps-1.5.1.crate"
	cargodeps_cmd_bin := exec.Command("curl", "-L", cargodeps_bin_url, "-o", "binary.bin")
	err = cargodeps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargodeps_src_url := "https://static.crates.io/crates/cargo-deps/cargo-deps-1.5.1.crate"
	cargodeps_cmd_src := exec.Command("curl", "-L", cargodeps_src_url, "-o", "source.tar.gz")
	err = cargodeps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargodeps_cmd_direct := exec.Command("./binary")
	err = cargodeps_cmd_direct.Run()
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
