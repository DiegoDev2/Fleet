package main

// Rust - Safe, concurrent, practical language
// Homepage: https://www.rust-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installRust() {
	// Método 1: Descargar y extraer .tar.gz
	rust_tar_url := "https://static.rust-lang.org/dist/rustc-1.80.1-src.tar.gz"
	rust_cmd_tar := exec.Command("curl", "-L", rust_tar_url, "-o", "package.tar.gz")
	err := rust_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rust_zip_url := "https://static.rust-lang.org/dist/rustc-1.80.1-src.zip"
	rust_cmd_zip := exec.Command("curl", "-L", rust_zip_url, "-o", "package.zip")
	err = rust_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rust_bin_url := "https://static.rust-lang.org/dist/rustc-1.80.1-src.bin"
	rust_cmd_bin := exec.Command("curl", "-L", rust_bin_url, "-o", "binary.bin")
	err = rust_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rust_src_url := "https://static.rust-lang.org/dist/rustc-1.80.1-src.src.tar.gz"
	rust_cmd_src := exec.Command("curl", "-L", rust_src_url, "-o", "source.tar.gz")
	err = rust_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rust_cmd_direct := exec.Command("./binary")
	err = rust_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libgit2@1.7")
exec.Command("latte", "install", "libgit2@1.7")
	fmt.Println("Instalando dependencia: libssh2")
exec.Command("latte", "install", "libssh2")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
