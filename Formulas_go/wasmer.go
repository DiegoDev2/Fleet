package main

// Wasmer - Universal WebAssembly Runtime
// Homepage: https://wasmer.io

import (
	"fmt"
	
	"os/exec"
)

func installWasmer() {
	// Método 1: Descargar y extraer .tar.gz
	wasmer_tar_url := "https://github.com/wasmerio/wasmer/archive/refs/tags/v4.3.6.tar.gz"
	wasmer_cmd_tar := exec.Command("curl", "-L", wasmer_tar_url, "-o", "package.tar.gz")
	err := wasmer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wasmer_zip_url := "https://github.com/wasmerio/wasmer/archive/refs/tags/v4.3.6.zip"
	wasmer_cmd_zip := exec.Command("curl", "-L", wasmer_zip_url, "-o", "package.zip")
	err = wasmer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wasmer_bin_url := "https://github.com/wasmerio/wasmer/archive/refs/tags/v4.3.6.bin"
	wasmer_cmd_bin := exec.Command("curl", "-L", wasmer_bin_url, "-o", "binary.bin")
	err = wasmer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wasmer_src_url := "https://github.com/wasmerio/wasmer/archive/refs/tags/v4.3.6.src.tar.gz"
	wasmer_cmd_src := exec.Command("curl", "-L", wasmer_src_url, "-o", "source.tar.gz")
	err = wasmer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wasmer_cmd_direct := exec.Command("./binary")
	err = wasmer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: wabt")
exec.Command("latte", "install", "wabt")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxkbcommon")
exec.Command("latte", "install", "libxkbcommon")
}
