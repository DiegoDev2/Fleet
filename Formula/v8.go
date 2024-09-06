package main

// V8 - Google's JavaScript engine
// Homepage: https://v8.dev/docs

import (
	"fmt"
	
	"os/exec"
)

func installV8() {
	// Método 1: Descargar y extraer .tar.gz
	v8_tar_url := "https://github.com/v8/v8/archive/refs/tags/12.7.224.16.tar.gz"
	v8_cmd_tar := exec.Command("curl", "-L", v8_tar_url, "-o", "package.tar.gz")
	err := v8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	v8_zip_url := "https://github.com/v8/v8/archive/refs/tags/12.7.224.16.zip"
	v8_cmd_zip := exec.Command("curl", "-L", v8_zip_url, "-o", "package.zip")
	err = v8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	v8_bin_url := "https://github.com/v8/v8/archive/refs/tags/12.7.224.16.bin"
	v8_cmd_bin := exec.Command("curl", "-L", v8_bin_url, "-o", "binary.bin")
	err = v8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	v8_src_url := "https://github.com/v8/v8/archive/refs/tags/12.7.224.16.src.tar.gz"
	v8_cmd_src := exec.Command("curl", "-L", v8_src_url, "-o", "source.tar.gz")
	err = v8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	v8_cmd_direct := exec.Command("./binary")
	err = v8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
}
