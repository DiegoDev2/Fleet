package main

// Wangle - Modular, composable client/server abstractions framework
// Homepage: https://github.com/facebook/wangle

import (
	"fmt"
	
	"os/exec"
)

func installWangle() {
	// Método 1: Descargar y extraer .tar.gz
	wangle_tar_url := "https://github.com/facebook/wangle/archive/refs/tags/v2024.08.26.00.tar.gz"
	wangle_cmd_tar := exec.Command("curl", "-L", wangle_tar_url, "-o", "package.tar.gz")
	err := wangle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wangle_zip_url := "https://github.com/facebook/wangle/archive/refs/tags/v2024.08.26.00.zip"
	wangle_cmd_zip := exec.Command("curl", "-L", wangle_zip_url, "-o", "package.zip")
	err = wangle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wangle_bin_url := "https://github.com/facebook/wangle/archive/refs/tags/v2024.08.26.00.bin"
	wangle_cmd_bin := exec.Command("curl", "-L", wangle_bin_url, "-o", "binary.bin")
	err = wangle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wangle_src_url := "https://github.com/facebook/wangle/archive/refs/tags/v2024.08.26.00.src.tar.gz"
	wangle_cmd_src := exec.Command("curl", "-L", wangle_src_url, "-o", "source.tar.gz")
	err = wangle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wangle_cmd_direct := exec.Command("./binary")
	err = wangle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: double-conversion")
exec.Command("latte", "install", "double-conversion")
	fmt.Println("Instalando dependencia: fizz")
exec.Command("latte", "install", "fizz")
	fmt.Println("Instalando dependencia: fmt")
exec.Command("latte", "install", "fmt")
	fmt.Println("Instalando dependencia: folly")
exec.Command("latte", "install", "folly")
	fmt.Println("Instalando dependencia: gflags")
exec.Command("latte", "install", "gflags")
	fmt.Println("Instalando dependencia: glog")
exec.Command("latte", "install", "glog")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: lz4")
exec.Command("latte", "install", "lz4")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
