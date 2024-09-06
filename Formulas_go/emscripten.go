package main

// Emscripten - LLVM bytecode to JavaScript compiler
// Homepage: https://emscripten.org/

import (
	"fmt"
	
	"os/exec"
)

func installEmscripten() {
	// Método 1: Descargar y extraer .tar.gz
	emscripten_tar_url := "https://github.com/emscripten-core/emscripten/archive/refs/tags/3.1.65.tar.gz"
	emscripten_cmd_tar := exec.Command("curl", "-L", emscripten_tar_url, "-o", "package.tar.gz")
	err := emscripten_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	emscripten_zip_url := "https://github.com/emscripten-core/emscripten/archive/refs/tags/3.1.65.zip"
	emscripten_cmd_zip := exec.Command("curl", "-L", emscripten_zip_url, "-o", "package.zip")
	err = emscripten_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	emscripten_bin_url := "https://github.com/emscripten-core/emscripten/archive/refs/tags/3.1.65.bin"
	emscripten_cmd_bin := exec.Command("curl", "-L", emscripten_bin_url, "-o", "binary.bin")
	err = emscripten_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	emscripten_src_url := "https://github.com/emscripten-core/emscripten/archive/refs/tags/3.1.65.src.tar.gz"
	emscripten_cmd_src := exec.Command("curl", "-L", emscripten_src_url, "-o", "source.tar.gz")
	err = emscripten_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	emscripten_cmd_direct := exec.Command("./binary")
	err = emscripten_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: yuicompressor")
exec.Command("latte", "install", "yuicompressor")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
