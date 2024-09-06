package main

// Envoy - Cloud-native high-performance edge/middle/service proxy
// Homepage: https://www.envoyproxy.io/index.html

import (
	"fmt"
	
	"os/exec"
)

func installEnvoy() {
	// Método 1: Descargar y extraer .tar.gz
	envoy_tar_url := "https://github.com/envoyproxy/envoy/archive/refs/tags/v1.31.0.tar.gz"
	envoy_cmd_tar := exec.Command("curl", "-L", envoy_tar_url, "-o", "package.tar.gz")
	err := envoy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	envoy_zip_url := "https://github.com/envoyproxy/envoy/archive/refs/tags/v1.31.0.zip"
	envoy_cmd_zip := exec.Command("curl", "-L", envoy_zip_url, "-o", "package.zip")
	err = envoy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	envoy_bin_url := "https://github.com/envoyproxy/envoy/archive/refs/tags/v1.31.0.bin"
	envoy_cmd_bin := exec.Command("curl", "-L", envoy_bin_url, "-o", "binary.bin")
	err = envoy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	envoy_src_url := "https://github.com/envoyproxy/envoy/archive/refs/tags/v1.31.0.src.tar.gz"
	envoy_cmd_src := exec.Command("curl", "-L", envoy_src_url, "-o", "source.tar.gz")
	err = envoy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	envoy_cmd_direct := exec.Command("./binary")
	err = envoy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bazelisk")
	exec.Command("latte", "install", "bazelisk").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
}
