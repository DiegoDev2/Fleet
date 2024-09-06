package main

// Cbindgen - Project for generating C bindings from Rust code
// Homepage: https://github.com/mozilla/cbindgen

import (
	"fmt"
	
	"os/exec"
)

func installCbindgen() {
	// Método 1: Descargar y extraer .tar.gz
	cbindgen_tar_url := "https://github.com/mozilla/cbindgen/archive/refs/tags/v0.26.0.tar.gz"
	cbindgen_cmd_tar := exec.Command("curl", "-L", cbindgen_tar_url, "-o", "package.tar.gz")
	err := cbindgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cbindgen_zip_url := "https://github.com/mozilla/cbindgen/archive/refs/tags/v0.26.0.zip"
	cbindgen_cmd_zip := exec.Command("curl", "-L", cbindgen_zip_url, "-o", "package.zip")
	err = cbindgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cbindgen_bin_url := "https://github.com/mozilla/cbindgen/archive/refs/tags/v0.26.0.bin"
	cbindgen_cmd_bin := exec.Command("curl", "-L", cbindgen_bin_url, "-o", "binary.bin")
	err = cbindgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cbindgen_src_url := "https://github.com/mozilla/cbindgen/archive/refs/tags/v0.26.0.src.tar.gz"
	cbindgen_cmd_src := exec.Command("curl", "-L", cbindgen_src_url, "-o", "source.tar.gz")
	err = cbindgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cbindgen_cmd_direct := exec.Command("./binary")
	err = cbindgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
