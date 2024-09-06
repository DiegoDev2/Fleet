package main

// Mvfst - QUIC transport protocol implementation
// Homepage: https://github.com/facebook/mvfst

import (
	"fmt"
	
	"os/exec"
)

func installMvfst() {
	// Método 1: Descargar y extraer .tar.gz
	mvfst_tar_url := "https://github.com/facebook/mvfst/archive/refs/tags/v2024.08.26.00.tar.gz"
	mvfst_cmd_tar := exec.Command("curl", "-L", mvfst_tar_url, "-o", "package.tar.gz")
	err := mvfst_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mvfst_zip_url := "https://github.com/facebook/mvfst/archive/refs/tags/v2024.08.26.00.zip"
	mvfst_cmd_zip := exec.Command("curl", "-L", mvfst_zip_url, "-o", "package.zip")
	err = mvfst_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mvfst_bin_url := "https://github.com/facebook/mvfst/archive/refs/tags/v2024.08.26.00.bin"
	mvfst_cmd_bin := exec.Command("curl", "-L", mvfst_bin_url, "-o", "binary.bin")
	err = mvfst_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mvfst_src_url := "https://github.com/facebook/mvfst/archive/refs/tags/v2024.08.26.00.src.tar.gz"
	mvfst_cmd_src := exec.Command("curl", "-L", mvfst_src_url, "-o", "source.tar.gz")
	err = mvfst_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mvfst_cmd_direct := exec.Command("./binary")
	err = mvfst_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: googletest")
	exec.Command("latte", "install", "googletest").Run()
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
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
