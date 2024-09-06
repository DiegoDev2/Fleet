package main

// Maturin - Build and publish Rust crates as Python packages
// Homepage: https://github.com/PyO3/maturin

import (
	"fmt"
	
	"os/exec"
)

func installMaturin() {
	// Método 1: Descargar y extraer .tar.gz
	maturin_tar_url := "https://github.com/PyO3/maturin/archive/refs/tags/v1.7.1.tar.gz"
	maturin_cmd_tar := exec.Command("curl", "-L", maturin_tar_url, "-o", "package.tar.gz")
	err := maturin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	maturin_zip_url := "https://github.com/PyO3/maturin/archive/refs/tags/v1.7.1.zip"
	maturin_cmd_zip := exec.Command("curl", "-L", maturin_zip_url, "-o", "package.zip")
	err = maturin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	maturin_bin_url := "https://github.com/PyO3/maturin/archive/refs/tags/v1.7.1.bin"
	maturin_cmd_bin := exec.Command("curl", "-L", maturin_bin_url, "-o", "binary.bin")
	err = maturin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	maturin_src_url := "https://github.com/PyO3/maturin/archive/refs/tags/v1.7.1.src.tar.gz"
	maturin_cmd_src := exec.Command("curl", "-L", maturin_src_url, "-o", "source.tar.gz")
	err = maturin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	maturin_cmd_direct := exec.Command("./binary")
	err = maturin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
