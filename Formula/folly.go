package main

// Folly - Collection of reusable C++ library artifacts developed at Facebook
// Homepage: https://github.com/facebook/folly

import (
	"fmt"
	
	"os/exec"
)

func installFolly() {
	// Método 1: Descargar y extraer .tar.gz
	folly_tar_url := "https://github.com/facebook/folly/archive/refs/tags/v2024.08.26.00.tar.gz"
	folly_cmd_tar := exec.Command("curl", "-L", folly_tar_url, "-o", "package.tar.gz")
	err := folly_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	folly_zip_url := "https://github.com/facebook/folly/archive/refs/tags/v2024.08.26.00.zip"
	folly_cmd_zip := exec.Command("curl", "-L", folly_zip_url, "-o", "package.zip")
	err = folly_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	folly_bin_url := "https://github.com/facebook/folly/archive/refs/tags/v2024.08.26.00.bin"
	folly_cmd_bin := exec.Command("curl", "-L", folly_bin_url, "-o", "binary.bin")
	err = folly_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	folly_src_url := "https://github.com/facebook/folly/archive/refs/tags/v2024.08.26.00.src.tar.gz"
	folly_cmd_src := exec.Command("curl", "-L", folly_src_url, "-o", "source.tar.gz")
	err = folly_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	folly_cmd_direct := exec.Command("./binary")
	err = folly_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: fast_float")
	exec.Command("latte", "install", "fast_float").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: double-conversion")
	exec.Command("latte", "install", "double-conversion").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
	fmt.Println("Instalando dependencia: glog")
	exec.Command("latte", "install", "glog").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
