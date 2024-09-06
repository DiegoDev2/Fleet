package main

// Rustfmt - Format Rust code
// Homepage: https://rust-lang.github.io/rustfmt/

import (
	"fmt"
	
	"os/exec"
)

func installRustfmt() {
	// Método 1: Descargar y extraer .tar.gz
	rustfmt_tar_url := "https://github.com/rust-lang/rustfmt/archive/refs/tags/v1.7.0.tar.gz"
	rustfmt_cmd_tar := exec.Command("curl", "-L", rustfmt_tar_url, "-o", "package.tar.gz")
	err := rustfmt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rustfmt_zip_url := "https://github.com/rust-lang/rustfmt/archive/refs/tags/v1.7.0.zip"
	rustfmt_cmd_zip := exec.Command("curl", "-L", rustfmt_zip_url, "-o", "package.zip")
	err = rustfmt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rustfmt_bin_url := "https://github.com/rust-lang/rustfmt/archive/refs/tags/v1.7.0.bin"
	rustfmt_cmd_bin := exec.Command("curl", "-L", rustfmt_bin_url, "-o", "binary.bin")
	err = rustfmt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rustfmt_src_url := "https://github.com/rust-lang/rustfmt/archive/refs/tags/v1.7.0.src.tar.gz"
	rustfmt_cmd_src := exec.Command("curl", "-L", rustfmt_src_url, "-o", "source.tar.gz")
	err = rustfmt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rustfmt_cmd_direct := exec.Command("./binary")
	err = rustfmt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
