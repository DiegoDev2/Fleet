package main

// Watchman - Watch files and take action when they change
// Homepage: https://github.com/facebook/watchman

import (
	"fmt"
	
	"os/exec"
)

func installWatchman() {
	// Método 1: Descargar y extraer .tar.gz
	watchman_tar_url := "https://github.com/facebook/watchman/archive/refs/tags/v2024.08.26.00.tar.gz"
	watchman_cmd_tar := exec.Command("curl", "-L", watchman_tar_url, "-o", "package.tar.gz")
	err := watchman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	watchman_zip_url := "https://github.com/facebook/watchman/archive/refs/tags/v2024.08.26.00.zip"
	watchman_cmd_zip := exec.Command("curl", "-L", watchman_zip_url, "-o", "package.zip")
	err = watchman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	watchman_bin_url := "https://github.com/facebook/watchman/archive/refs/tags/v2024.08.26.00.bin"
	watchman_cmd_bin := exec.Command("curl", "-L", watchman_bin_url, "-o", "binary.bin")
	err = watchman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	watchman_src_url := "https://github.com/facebook/watchman/archive/refs/tags/v2024.08.26.00.src.tar.gz"
	watchman_cmd_src := exec.Command("curl", "-L", watchman_src_url, "-o", "source.tar.gz")
	err = watchman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	watchman_cmd_direct := exec.Command("./binary")
	err = watchman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: cpptoml")
exec.Command("latte", "install", "cpptoml")
	fmt.Println("Instalando dependencia: googletest")
exec.Command("latte", "install", "googletest")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: edencommon")
exec.Command("latte", "install", "edencommon")
	fmt.Println("Instalando dependencia: fb303")
exec.Command("latte", "install", "fb303")
	fmt.Println("Instalando dependencia: fbthrift")
exec.Command("latte", "install", "fbthrift")
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
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: pcre2")
exec.Command("latte", "install", "pcre2")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: libunwind")
exec.Command("latte", "install", "libunwind")
}
