package main

// Rure - C API for RUst's REgex engine
// Homepage: https://github.com/rust-lang/regex/tree/HEAD/regex-capi

import (
	"fmt"
	
	"os/exec"
)

func installRure() {
	// Método 1: Descargar y extraer .tar.gz
	rure_tar_url := "https://github.com/rust-lang/regex/archive/refs/tags/1.10.6.tar.gz"
	rure_cmd_tar := exec.Command("curl", "-L", rure_tar_url, "-o", "package.tar.gz")
	err := rure_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rure_zip_url := "https://github.com/rust-lang/regex/archive/refs/tags/1.10.6.zip"
	rure_cmd_zip := exec.Command("curl", "-L", rure_zip_url, "-o", "package.zip")
	err = rure_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rure_bin_url := "https://github.com/rust-lang/regex/archive/refs/tags/1.10.6.bin"
	rure_cmd_bin := exec.Command("curl", "-L", rure_bin_url, "-o", "binary.bin")
	err = rure_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rure_src_url := "https://github.com/rust-lang/regex/archive/refs/tags/1.10.6.src.tar.gz"
	rure_cmd_src := exec.Command("curl", "-L", rure_src_url, "-o", "source.tar.gz")
	err = rure_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rure_cmd_direct := exec.Command("./binary")
	err = rure_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
