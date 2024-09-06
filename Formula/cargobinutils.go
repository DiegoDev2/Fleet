package main

// CargoBinutils - Cargo subcommands to invoke the LLVM tools shipped with the Rust toolchain
// Homepage: https://github.com/rust-embedded/cargo-binutils

import (
	"fmt"
	
	"os/exec"
)

func installCargoBinutils() {
	// Método 1: Descargar y extraer .tar.gz
	cargobinutils_tar_url := "https://github.com/rust-embedded/cargo-binutils/archive/refs/tags/v0.3.6.tar.gz"
	cargobinutils_cmd_tar := exec.Command("curl", "-L", cargobinutils_tar_url, "-o", "package.tar.gz")
	err := cargobinutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargobinutils_zip_url := "https://github.com/rust-embedded/cargo-binutils/archive/refs/tags/v0.3.6.zip"
	cargobinutils_cmd_zip := exec.Command("curl", "-L", cargobinutils_zip_url, "-o", "package.zip")
	err = cargobinutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargobinutils_bin_url := "https://github.com/rust-embedded/cargo-binutils/archive/refs/tags/v0.3.6.bin"
	cargobinutils_cmd_bin := exec.Command("curl", "-L", cargobinutils_bin_url, "-o", "binary.bin")
	err = cargobinutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargobinutils_src_url := "https://github.com/rust-embedded/cargo-binutils/archive/refs/tags/v0.3.6.src.tar.gz"
	cargobinutils_cmd_src := exec.Command("curl", "-L", cargobinutils_src_url, "-o", "source.tar.gz")
	err = cargobinutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargobinutils_cmd_direct := exec.Command("./binary")
	err = cargobinutils_cmd_direct.Run()
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
