package main

// Fbthrift - Facebook's branch of Apache Thrift, including a new C++ server
// Homepage: https://github.com/facebook/fbthrift

import (
	"fmt"
	
	"os/exec"
)

func installFbthrift() {
	// Método 1: Descargar y extraer .tar.gz
	fbthrift_tar_url := "https://github.com/facebook/fbthrift/archive/refs/tags/v2024.08.26.00.tar.gz"
	fbthrift_cmd_tar := exec.Command("curl", "-L", fbthrift_tar_url, "-o", "package.tar.gz")
	err := fbthrift_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fbthrift_zip_url := "https://github.com/facebook/fbthrift/archive/refs/tags/v2024.08.26.00.zip"
	fbthrift_cmd_zip := exec.Command("curl", "-L", fbthrift_zip_url, "-o", "package.zip")
	err = fbthrift_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fbthrift_bin_url := "https://github.com/facebook/fbthrift/archive/refs/tags/v2024.08.26.00.bin"
	fbthrift_cmd_bin := exec.Command("curl", "-L", fbthrift_bin_url, "-o", "binary.bin")
	err = fbthrift_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fbthrift_src_url := "https://github.com/facebook/fbthrift/archive/refs/tags/v2024.08.26.00.src.tar.gz"
	fbthrift_cmd_src := exec.Command("curl", "-L", fbthrift_src_url, "-o", "source.tar.gz")
	err = fbthrift_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fbthrift_cmd_direct := exec.Command("./binary")
	err = fbthrift_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: double-conversion")
	exec.Command("latte", "install", "double-conversion").Run()
	fmt.Println("Instalando dependencia: fizz")
	exec.Command("latte", "install", "fizz").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: folly")
	exec.Command("latte", "install", "folly").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
	fmt.Println("Instalando dependencia: glog")
	exec.Command("latte", "install", "glog").Run()
	fmt.Println("Instalando dependencia: mvfst")
	exec.Command("latte", "install", "mvfst").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: wangle")
	exec.Command("latte", "install", "wangle").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
