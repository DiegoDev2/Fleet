package main

// NodeAT18 - Platform built on V8 to build network applications
// Homepage: https://nodejs.org/

import (
	"fmt"
	
	"os/exec"
)

func installNodeAT18() {
	// Método 1: Descargar y extraer .tar.gz
	nodeat18_tar_url := "https://nodejs.org/dist/v18.20.4/node-v18.20.4.tar.xz"
	nodeat18_cmd_tar := exec.Command("curl", "-L", nodeat18_tar_url, "-o", "package.tar.gz")
	err := nodeat18_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nodeat18_zip_url := "https://nodejs.org/dist/v18.20.4/node-v18.20.4.tar.xz"
	nodeat18_cmd_zip := exec.Command("curl", "-L", nodeat18_zip_url, "-o", "package.zip")
	err = nodeat18_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nodeat18_bin_url := "https://nodejs.org/dist/v18.20.4/node-v18.20.4.tar.xz"
	nodeat18_cmd_bin := exec.Command("curl", "-L", nodeat18_bin_url, "-o", "binary.bin")
	err = nodeat18_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nodeat18_src_url := "https://nodejs.org/dist/v18.20.4/node-v18.20.4.tar.xz"
	nodeat18_cmd_src := exec.Command("curl", "-L", nodeat18_src_url, "-o", "source.tar.gz")
	err = nodeat18_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nodeat18_cmd_direct := exec.Command("./binary")
	err = nodeat18_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
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
