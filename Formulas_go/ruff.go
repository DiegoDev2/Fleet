package main

// Ruff - Extremely fast Python linter, written in Rust
// Homepage: https://docs.astral.sh/ruff/

import (
	"fmt"
	
	"os/exec"
)

func installRuff() {
	// Método 1: Descargar y extraer .tar.gz
	ruff_tar_url := "https://github.com/astral-sh/ruff/archive/refs/tags/0.6.4.tar.gz"
	ruff_cmd_tar := exec.Command("curl", "-L", ruff_tar_url, "-o", "package.tar.gz")
	err := ruff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ruff_zip_url := "https://github.com/astral-sh/ruff/archive/refs/tags/0.6.4.zip"
	ruff_cmd_zip := exec.Command("curl", "-L", ruff_zip_url, "-o", "package.zip")
	err = ruff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ruff_bin_url := "https://github.com/astral-sh/ruff/archive/refs/tags/0.6.4.bin"
	ruff_cmd_bin := exec.Command("curl", "-L", ruff_bin_url, "-o", "binary.bin")
	err = ruff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ruff_src_url := "https://github.com/astral-sh/ruff/archive/refs/tags/0.6.4.src.tar.gz"
	ruff_cmd_src := exec.Command("curl", "-L", ruff_src_url, "-o", "source.tar.gz")
	err = ruff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ruff_cmd_direct := exec.Command("./binary")
	err = ruff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
