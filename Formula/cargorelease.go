package main

// CargoRelease - Cargo subcommand `release`: everything about releasing a rust crate
// Homepage: https://github.com/crate-ci/cargo-release

import (
	"fmt"
	
	"os/exec"
)

func installCargoRelease() {
	// Método 1: Descargar y extraer .tar.gz
	cargorelease_tar_url := "https://github.com/crate-ci/cargo-release/archive/refs/tags/v0.25.10.tar.gz"
	cargorelease_cmd_tar := exec.Command("curl", "-L", cargorelease_tar_url, "-o", "package.tar.gz")
	err := cargorelease_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargorelease_zip_url := "https://github.com/crate-ci/cargo-release/archive/refs/tags/v0.25.10.zip"
	cargorelease_cmd_zip := exec.Command("curl", "-L", cargorelease_zip_url, "-o", "package.zip")
	err = cargorelease_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargorelease_bin_url := "https://github.com/crate-ci/cargo-release/archive/refs/tags/v0.25.10.bin"
	cargorelease_cmd_bin := exec.Command("curl", "-L", cargorelease_bin_url, "-o", "binary.bin")
	err = cargorelease_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargorelease_src_url := "https://github.com/crate-ci/cargo-release/archive/refs/tags/v0.25.10.src.tar.gz"
	cargorelease_cmd_src := exec.Command("curl", "-L", cargorelease_src_url, "-o", "source.tar.gz")
	err = cargorelease_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargorelease_cmd_direct := exec.Command("./binary")
	err = cargorelease_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: rustup")
	exec.Command("latte", "install", "rustup").Run()
	fmt.Println("Instalando dependencia: libgit2@1.7")
	exec.Command("latte", "install", "libgit2@1.7").Run()
}
