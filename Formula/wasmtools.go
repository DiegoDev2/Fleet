package main

// WasmTools - Low level tooling for WebAssembly in Rust
// Homepage: https://github.com/bytecodealliance/wasm-tools

import (
	"fmt"
	
	"os/exec"
)

func installWasmTools() {
	// Método 1: Descargar y extraer .tar.gz
	wasmtools_tar_url := "https://github.com/bytecodealliance/wasm-tools/archive/refs/tags/v1.216.0.tar.gz"
	wasmtools_cmd_tar := exec.Command("curl", "-L", wasmtools_tar_url, "-o", "package.tar.gz")
	err := wasmtools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wasmtools_zip_url := "https://github.com/bytecodealliance/wasm-tools/archive/refs/tags/v1.216.0.zip"
	wasmtools_cmd_zip := exec.Command("curl", "-L", wasmtools_zip_url, "-o", "package.zip")
	err = wasmtools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wasmtools_bin_url := "https://github.com/bytecodealliance/wasm-tools/archive/refs/tags/v1.216.0.bin"
	wasmtools_cmd_bin := exec.Command("curl", "-L", wasmtools_bin_url, "-o", "binary.bin")
	err = wasmtools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wasmtools_src_url := "https://github.com/bytecodealliance/wasm-tools/archive/refs/tags/v1.216.0.src.tar.gz"
	wasmtools_cmd_src := exec.Command("curl", "-L", wasmtools_src_url, "-o", "source.tar.gz")
	err = wasmtools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wasmtools_cmd_direct := exec.Command("./binary")
	err = wasmtools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
