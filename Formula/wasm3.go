package main

// Wasm3 - High performance WebAssembly interpreter
// Homepage: https://github.com/wasm3/wasm3

import (
	"fmt"
	
	"os/exec"
)

func installWasm3() {
	// Método 1: Descargar y extraer .tar.gz
	wasm3_tar_url := "https://github.com/wasm3/wasm3/archive/refs/tags/v0.5.0.tar.gz"
	wasm3_cmd_tar := exec.Command("curl", "-L", wasm3_tar_url, "-o", "package.tar.gz")
	err := wasm3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wasm3_zip_url := "https://github.com/wasm3/wasm3/archive/refs/tags/v0.5.0.zip"
	wasm3_cmd_zip := exec.Command("curl", "-L", wasm3_zip_url, "-o", "package.zip")
	err = wasm3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wasm3_bin_url := "https://github.com/wasm3/wasm3/archive/refs/tags/v0.5.0.bin"
	wasm3_cmd_bin := exec.Command("curl", "-L", wasm3_bin_url, "-o", "binary.bin")
	err = wasm3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wasm3_src_url := "https://github.com/wasm3/wasm3/archive/refs/tags/v0.5.0.src.tar.gz"
	wasm3_cmd_src := exec.Command("curl", "-L", wasm3_src_url, "-o", "source.tar.gz")
	err = wasm3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wasm3_cmd_direct := exec.Command("./binary")
	err = wasm3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
