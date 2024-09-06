package main

// Node - Platform built on V8 to build network applications
// Homepage: https://nodejs.org/

import (
	"fmt"
	
	"os/exec"
)

func installNode() {
	// Método 1: Descargar y extraer .tar.gz
	node_tar_url := "https://nodejs.org/dist/v22.8.0/node-v22.8.0.tar.xz"
	node_cmd_tar := exec.Command("curl", "-L", node_tar_url, "-o", "package.tar.gz")
	err := node_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	node_zip_url := "https://nodejs.org/dist/v22.8.0/node-v22.8.0.tar.xz"
	node_cmd_zip := exec.Command("curl", "-L", node_zip_url, "-o", "package.zip")
	err = node_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	node_bin_url := "https://nodejs.org/dist/v22.8.0/node-v22.8.0.tar.xz"
	node_cmd_bin := exec.Command("curl", "-L", node_bin_url, "-o", "binary.bin")
	err = node_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	node_src_url := "https://nodejs.org/dist/v22.8.0/node-v22.8.0.tar.xz"
	node_cmd_src := exec.Command("curl", "-L", node_src_url, "-o", "source.tar.gz")
	err = node_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	node_cmd_direct := exec.Command("./binary")
	err = node_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
	fmt.Println("Instalando dependencia: c-ares")
exec.Command("latte", "install", "c-ares")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libnghttp2")
exec.Command("latte", "install", "libnghttp2")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
