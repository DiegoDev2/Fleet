package main

// CargoNextest - Next-generation test runner for Rust
// Homepage: https://nexte.st

import (
	"fmt"
	
	"os/exec"
)

func installCargoNextest() {
	// Método 1: Descargar y extraer .tar.gz
	cargonextest_tar_url := "https://github.com/nextest-rs/nextest/archive/refs/tags/cargo-nextest-0.9.78.tar.gz"
	cargonextest_cmd_tar := exec.Command("curl", "-L", cargonextest_tar_url, "-o", "package.tar.gz")
	err := cargonextest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargonextest_zip_url := "https://github.com/nextest-rs/nextest/archive/refs/tags/cargo-nextest-0.9.78.zip"
	cargonextest_cmd_zip := exec.Command("curl", "-L", cargonextest_zip_url, "-o", "package.zip")
	err = cargonextest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargonextest_bin_url := "https://github.com/nextest-rs/nextest/archive/refs/tags/cargo-nextest-0.9.78.bin"
	cargonextest_cmd_bin := exec.Command("curl", "-L", cargonextest_bin_url, "-o", "binary.bin")
	err = cargonextest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargonextest_src_url := "https://github.com/nextest-rs/nextest/archive/refs/tags/cargo-nextest-0.9.78.src.tar.gz"
	cargonextest_cmd_src := exec.Command("curl", "-L", cargonextest_src_url, "-o", "source.tar.gz")
	err = cargonextest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargonextest_cmd_direct := exec.Command("./binary")
	err = cargonextest_cmd_direct.Run()
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
