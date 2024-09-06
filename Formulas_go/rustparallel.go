package main

// RustParallel - Run commands in parallel with Rust's Tokio framework
// Homepage: https://github.com/aaronriekenberg/rust-parallel

import (
	"fmt"
	
	"os/exec"
)

func installRustParallel() {
	// Método 1: Descargar y extraer .tar.gz
	rustparallel_tar_url := "https://github.com/aaronriekenberg/rust-parallel/archive/refs/tags/v1.18.1.tar.gz"
	rustparallel_cmd_tar := exec.Command("curl", "-L", rustparallel_tar_url, "-o", "package.tar.gz")
	err := rustparallel_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rustparallel_zip_url := "https://github.com/aaronriekenberg/rust-parallel/archive/refs/tags/v1.18.1.zip"
	rustparallel_cmd_zip := exec.Command("curl", "-L", rustparallel_zip_url, "-o", "package.zip")
	err = rustparallel_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rustparallel_bin_url := "https://github.com/aaronriekenberg/rust-parallel/archive/refs/tags/v1.18.1.bin"
	rustparallel_cmd_bin := exec.Command("curl", "-L", rustparallel_bin_url, "-o", "binary.bin")
	err = rustparallel_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rustparallel_src_url := "https://github.com/aaronriekenberg/rust-parallel/archive/refs/tags/v1.18.1.src.tar.gz"
	rustparallel_cmd_src := exec.Command("curl", "-L", rustparallel_src_url, "-o", "source.tar.gz")
	err = rustparallel_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rustparallel_cmd_direct := exec.Command("./binary")
	err = rustparallel_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
