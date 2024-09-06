package main

// CargoAudit - Audit Cargo.lock files for crates with security vulnerabilities
// Homepage: https://rustsec.org/

import (
	"fmt"
	
	"os/exec"
)

func installCargoAudit() {
	// Método 1: Descargar y extraer .tar.gz
	cargoaudit_tar_url := "https://github.com/RustSec/rustsec/archive/refs/tags/cargo-audit/v0.20.0.tar.gz"
	cargoaudit_cmd_tar := exec.Command("curl", "-L", cargoaudit_tar_url, "-o", "package.tar.gz")
	err := cargoaudit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargoaudit_zip_url := "https://github.com/RustSec/rustsec/archive/refs/tags/cargo-audit/v0.20.0.zip"
	cargoaudit_cmd_zip := exec.Command("curl", "-L", cargoaudit_zip_url, "-o", "package.zip")
	err = cargoaudit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargoaudit_bin_url := "https://github.com/RustSec/rustsec/archive/refs/tags/cargo-audit/v0.20.0.bin"
	cargoaudit_cmd_bin := exec.Command("curl", "-L", cargoaudit_bin_url, "-o", "binary.bin")
	err = cargoaudit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargoaudit_src_url := "https://github.com/RustSec/rustsec/archive/refs/tags/cargo-audit/v0.20.0.src.tar.gz"
	cargoaudit_cmd_src := exec.Command("curl", "-L", cargoaudit_src_url, "-o", "source.tar.gz")
	err = cargoaudit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargoaudit_cmd_direct := exec.Command("./binary")
	err = cargoaudit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
