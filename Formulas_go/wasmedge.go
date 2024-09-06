package main

// Wasmedge - Lightweight, high-performance, and extensible WebAssembly runtime
// Homepage: https://WasmEdge.org/

import (
	"fmt"
	
	"os/exec"
)

func installWasmedge() {
	// Método 1: Descargar y extraer .tar.gz
	wasmedge_tar_url := "https://github.com/WasmEdge/WasmEdge/releases/download/0.14.0/WasmEdge-0.14.0-src.tar.gz"
	wasmedge_cmd_tar := exec.Command("curl", "-L", wasmedge_tar_url, "-o", "package.tar.gz")
	err := wasmedge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wasmedge_zip_url := "https://github.com/WasmEdge/WasmEdge/releases/download/0.14.0/WasmEdge-0.14.0-src.zip"
	wasmedge_cmd_zip := exec.Command("curl", "-L", wasmedge_zip_url, "-o", "package.zip")
	err = wasmedge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wasmedge_bin_url := "https://github.com/WasmEdge/WasmEdge/releases/download/0.14.0/WasmEdge-0.14.0-src.bin"
	wasmedge_cmd_bin := exec.Command("curl", "-L", wasmedge_bin_url, "-o", "binary.bin")
	err = wasmedge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wasmedge_src_url := "https://github.com/WasmEdge/WasmEdge/releases/download/0.14.0/WasmEdge-0.14.0-src.src.tar.gz"
	wasmedge_cmd_src := exec.Command("curl", "-L", wasmedge_src_url, "-o", "source.tar.gz")
	err = wasmedge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wasmedge_cmd_direct := exec.Command("./binary")
	err = wasmedge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: spdlog")
exec.Command("latte", "install", "spdlog")
}
