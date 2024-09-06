package main

// Fizz - C++14 implementation of the TLS-1.3 standard
// Homepage: https://github.com/facebookincubator/fizz

import (
	"fmt"
	
	"os/exec"
)

func installFizz() {
	// Método 1: Descargar y extraer .tar.gz
	fizz_tar_url := "https://github.com/facebookincubator/fizz/releases/download/v2024.08.26.00/fizz-v2024.08.26.00.tar.gz"
	fizz_cmd_tar := exec.Command("curl", "-L", fizz_tar_url, "-o", "package.tar.gz")
	err := fizz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fizz_zip_url := "https://github.com/facebookincubator/fizz/releases/download/v2024.08.26.00/fizz-v2024.08.26.00.zip"
	fizz_cmd_zip := exec.Command("curl", "-L", fizz_zip_url, "-o", "package.zip")
	err = fizz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fizz_bin_url := "https://github.com/facebookincubator/fizz/releases/download/v2024.08.26.00/fizz-v2024.08.26.00.bin"
	fizz_cmd_bin := exec.Command("curl", "-L", fizz_bin_url, "-o", "binary.bin")
	err = fizz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fizz_src_url := "https://github.com/facebookincubator/fizz/releases/download/v2024.08.26.00/fizz-v2024.08.26.00.src.tar.gz"
	fizz_cmd_src := exec.Command("curl", "-L", fizz_src_url, "-o", "source.tar.gz")
	err = fizz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fizz_cmd_direct := exec.Command("./binary")
	err = fizz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: double-conversion")
	exec.Command("latte", "install", "double-conversion").Run()
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
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
