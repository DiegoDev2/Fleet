package main

// RustlsFfi - FFI bindings for the rustls TLS library
// Homepage: https://github.com/rustls/rustls-ffi

import (
	"fmt"
	
	"os/exec"
)

func installRustlsFfi() {
	// Método 1: Descargar y extraer .tar.gz
	rustlsffi_tar_url := "https://github.com/rustls/rustls-ffi/archive/refs/tags/v0.13.0.tar.gz"
	rustlsffi_cmd_tar := exec.Command("curl", "-L", rustlsffi_tar_url, "-o", "package.tar.gz")
	err := rustlsffi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rustlsffi_zip_url := "https://github.com/rustls/rustls-ffi/archive/refs/tags/v0.13.0.zip"
	rustlsffi_cmd_zip := exec.Command("curl", "-L", rustlsffi_zip_url, "-o", "package.zip")
	err = rustlsffi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rustlsffi_bin_url := "https://github.com/rustls/rustls-ffi/archive/refs/tags/v0.13.0.bin"
	rustlsffi_cmd_bin := exec.Command("curl", "-L", rustlsffi_bin_url, "-o", "binary.bin")
	err = rustlsffi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rustlsffi_src_url := "https://github.com/rustls/rustls-ffi/archive/refs/tags/v0.13.0.src.tar.gz"
	rustlsffi_cmd_src := exec.Command("curl", "-L", rustlsffi_src_url, "-o", "source.tar.gz")
	err = rustlsffi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rustlsffi_cmd_direct := exec.Command("./binary")
	err = rustlsffi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cargo-c")
exec.Command("latte", "install", "cargo-c")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
