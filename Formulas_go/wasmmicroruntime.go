package main

// WasmMicroRuntime - WebAssembly Micro Runtime (WAMR)
// Homepage: https://github.com/bytecodealliance/wasm-micro-runtime

import (
	"fmt"
	
	"os/exec"
)

func installWasmMicroRuntime() {
	// Método 1: Descargar y extraer .tar.gz
	wasmmicroruntime_tar_url := "https://github.com/bytecodealliance/wasm-micro-runtime/archive/refs/tags/WAMR-2.1.2.tar.gz"
	wasmmicroruntime_cmd_tar := exec.Command("curl", "-L", wasmmicroruntime_tar_url, "-o", "package.tar.gz")
	err := wasmmicroruntime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wasmmicroruntime_zip_url := "https://github.com/bytecodealliance/wasm-micro-runtime/archive/refs/tags/WAMR-2.1.2.zip"
	wasmmicroruntime_cmd_zip := exec.Command("curl", "-L", wasmmicroruntime_zip_url, "-o", "package.zip")
	err = wasmmicroruntime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wasmmicroruntime_bin_url := "https://github.com/bytecodealliance/wasm-micro-runtime/archive/refs/tags/WAMR-2.1.2.bin"
	wasmmicroruntime_cmd_bin := exec.Command("curl", "-L", wasmmicroruntime_bin_url, "-o", "binary.bin")
	err = wasmmicroruntime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wasmmicroruntime_src_url := "https://github.com/bytecodealliance/wasm-micro-runtime/archive/refs/tags/WAMR-2.1.2.src.tar.gz"
	wasmmicroruntime_cmd_src := exec.Command("curl", "-L", wasmmicroruntime_src_url, "-o", "source.tar.gz")
	err = wasmmicroruntime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wasmmicroruntime_cmd_direct := exec.Command("./binary")
	err = wasmmicroruntime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
