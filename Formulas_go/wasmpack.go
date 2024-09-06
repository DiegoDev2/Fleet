package main

// WasmPack - Your favorite rust -> wasm workflow tool!
// Homepage: https://rustwasm.github.io/wasm-pack/

import (
	"fmt"
	
	"os/exec"
)

func installWasmPack() {
	// Método 1: Descargar y extraer .tar.gz
	wasmpack_tar_url := "https://github.com/rustwasm/wasm-pack/archive/refs/tags/v0.13.0.tar.gz"
	wasmpack_cmd_tar := exec.Command("curl", "-L", wasmpack_tar_url, "-o", "package.tar.gz")
	err := wasmpack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wasmpack_zip_url := "https://github.com/rustwasm/wasm-pack/archive/refs/tags/v0.13.0.zip"
	wasmpack_cmd_zip := exec.Command("curl", "-L", wasmpack_zip_url, "-o", "package.zip")
	err = wasmpack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wasmpack_bin_url := "https://github.com/rustwasm/wasm-pack/archive/refs/tags/v0.13.0.bin"
	wasmpack_cmd_bin := exec.Command("curl", "-L", wasmpack_bin_url, "-o", "binary.bin")
	err = wasmpack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wasmpack_src_url := "https://github.com/rustwasm/wasm-pack/archive/refs/tags/v0.13.0.src.tar.gz"
	wasmpack_cmd_src := exec.Command("curl", "-L", wasmpack_src_url, "-o", "source.tar.gz")
	err = wasmpack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wasmpack_cmd_direct := exec.Command("./binary")
	err = wasmpack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: rustup")
exec.Command("latte", "install", "rustup")
}
