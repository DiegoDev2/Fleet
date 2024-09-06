package main

// CargoLlvmCov - Cargo subcommand to easily use LLVM source-based code coverage
// Homepage: https://github.com/taiki-e/cargo-llvm-cov

import (
	"fmt"
	
	"os/exec"
)

func installCargoLlvmCov() {
	// Método 1: Descargar y extraer .tar.gz
	cargollvmcov_tar_url := "https://static.crates.io/crates/cargo-llvm-cov/cargo-llvm-cov-0.6.12.crate"
	cargollvmcov_cmd_tar := exec.Command("curl", "-L", cargollvmcov_tar_url, "-o", "package.tar.gz")
	err := cargollvmcov_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargollvmcov_zip_url := "https://static.crates.io/crates/cargo-llvm-cov/cargo-llvm-cov-0.6.12.crate"
	cargollvmcov_cmd_zip := exec.Command("curl", "-L", cargollvmcov_zip_url, "-o", "package.zip")
	err = cargollvmcov_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargollvmcov_bin_url := "https://static.crates.io/crates/cargo-llvm-cov/cargo-llvm-cov-0.6.12.crate"
	cargollvmcov_cmd_bin := exec.Command("curl", "-L", cargollvmcov_bin_url, "-o", "binary.bin")
	err = cargollvmcov_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargollvmcov_src_url := "https://static.crates.io/crates/cargo-llvm-cov/cargo-llvm-cov-0.6.12.crate"
	cargollvmcov_cmd_src := exec.Command("curl", "-L", cargollvmcov_src_url, "-o", "source.tar.gz")
	err = cargollvmcov_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargollvmcov_cmd_direct := exec.Command("./binary")
	err = cargollvmcov_cmd_direct.Run()
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
