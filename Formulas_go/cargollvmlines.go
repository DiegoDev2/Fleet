package main

// CargoLlvmLines - Count lines of LLVM IR per generic function
// Homepage: https://github.com/dtolnay/cargo-llvm-lines

import (
	"fmt"
	
	"os/exec"
)

func installCargoLlvmLines() {
	// Método 1: Descargar y extraer .tar.gz
	cargollvmlines_tar_url := "https://github.com/dtolnay/cargo-llvm-lines/archive/refs/tags/0.4.40.tar.gz"
	cargollvmlines_cmd_tar := exec.Command("curl", "-L", cargollvmlines_tar_url, "-o", "package.tar.gz")
	err := cargollvmlines_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cargollvmlines_zip_url := "https://github.com/dtolnay/cargo-llvm-lines/archive/refs/tags/0.4.40.zip"
	cargollvmlines_cmd_zip := exec.Command("curl", "-L", cargollvmlines_zip_url, "-o", "package.zip")
	err = cargollvmlines_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cargollvmlines_bin_url := "https://github.com/dtolnay/cargo-llvm-lines/archive/refs/tags/0.4.40.bin"
	cargollvmlines_cmd_bin := exec.Command("curl", "-L", cargollvmlines_bin_url, "-o", "binary.bin")
	err = cargollvmlines_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cargollvmlines_src_url := "https://github.com/dtolnay/cargo-llvm-lines/archive/refs/tags/0.4.40.src.tar.gz"
	cargollvmlines_cmd_src := exec.Command("curl", "-L", cargollvmlines_src_url, "-o", "source.tar.gz")
	err = cargollvmlines_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cargollvmlines_cmd_direct := exec.Command("./binary")
	err = cargollvmlines_cmd_direct.Run()
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
