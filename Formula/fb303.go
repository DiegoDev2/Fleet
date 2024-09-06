package main

// Fb303 - Thrift functions for querying information from a service
// Homepage: https://github.com/facebook/fb303

import (
	"fmt"
	
	"os/exec"
)

func installFb303() {
	// Método 1: Descargar y extraer .tar.gz
	fb303_tar_url := "https://github.com/facebook/fb303/archive/refs/tags/v2024.08.26.00.tar.gz"
	fb303_cmd_tar := exec.Command("curl", "-L", fb303_tar_url, "-o", "package.tar.gz")
	err := fb303_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fb303_zip_url := "https://github.com/facebook/fb303/archive/refs/tags/v2024.08.26.00.zip"
	fb303_cmd_zip := exec.Command("curl", "-L", fb303_zip_url, "-o", "package.zip")
	err = fb303_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fb303_bin_url := "https://github.com/facebook/fb303/archive/refs/tags/v2024.08.26.00.bin"
	fb303_cmd_bin := exec.Command("curl", "-L", fb303_bin_url, "-o", "binary.bin")
	err = fb303_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fb303_src_url := "https://github.com/facebook/fb303/archive/refs/tags/v2024.08.26.00.src.tar.gz"
	fb303_cmd_src := exec.Command("curl", "-L", fb303_src_url, "-o", "source.tar.gz")
	err = fb303_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fb303_cmd_direct := exec.Command("./binary")
	err = fb303_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fbthrift")
	exec.Command("latte", "install", "fbthrift").Run()
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
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
