package main

// Binaryen - Compiler infrastructure and toolchain library for WebAssembly
// Homepage: https://webassembly.org/

import (
	"fmt"
	
	"os/exec"
)

func installBinaryen() {
	// Método 1: Descargar y extraer .tar.gz
	binaryen_tar_url := "https://github.com/WebAssembly/binaryen/archive/refs/tags/version_119.tar.gz"
	binaryen_cmd_tar := exec.Command("curl", "-L", binaryen_tar_url, "-o", "package.tar.gz")
	err := binaryen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	binaryen_zip_url := "https://github.com/WebAssembly/binaryen/archive/refs/tags/version_119.zip"
	binaryen_cmd_zip := exec.Command("curl", "-L", binaryen_zip_url, "-o", "package.zip")
	err = binaryen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	binaryen_bin_url := "https://github.com/WebAssembly/binaryen/archive/refs/tags/version_119.bin"
	binaryen_cmd_bin := exec.Command("curl", "-L", binaryen_bin_url, "-o", "binary.bin")
	err = binaryen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	binaryen_src_url := "https://github.com/WebAssembly/binaryen/archive/refs/tags/version_119.src.tar.gz"
	binaryen_cmd_src := exec.Command("curl", "-L", binaryen_src_url, "-o", "source.tar.gz")
	err = binaryen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	binaryen_cmd_direct := exec.Command("./binary")
	err = binaryen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
