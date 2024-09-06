package main

// LibjsonRpcCpp - C++ framework for json-rpc
// Homepage: https://github.com/cinemast/libjson-rpc-cpp

import (
	"fmt"
	
	"os/exec"
)

func installLibjsonRpcCpp() {
	// Método 1: Descargar y extraer .tar.gz
	libjsonrpccpp_tar_url := "https://github.com/cinemast/libjson-rpc-cpp/archive/refs/tags/v1.4.1.tar.gz"
	libjsonrpccpp_cmd_tar := exec.Command("curl", "-L", libjsonrpccpp_tar_url, "-o", "package.tar.gz")
	err := libjsonrpccpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libjsonrpccpp_zip_url := "https://github.com/cinemast/libjson-rpc-cpp/archive/refs/tags/v1.4.1.zip"
	libjsonrpccpp_cmd_zip := exec.Command("curl", "-L", libjsonrpccpp_zip_url, "-o", "package.zip")
	err = libjsonrpccpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libjsonrpccpp_bin_url := "https://github.com/cinemast/libjson-rpc-cpp/archive/refs/tags/v1.4.1.bin"
	libjsonrpccpp_cmd_bin := exec.Command("curl", "-L", libjsonrpccpp_bin_url, "-o", "binary.bin")
	err = libjsonrpccpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libjsonrpccpp_src_url := "https://github.com/cinemast/libjson-rpc-cpp/archive/refs/tags/v1.4.1.src.tar.gz"
	libjsonrpccpp_cmd_src := exec.Command("curl", "-L", libjsonrpccpp_src_url, "-o", "source.tar.gz")
	err = libjsonrpccpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libjsonrpccpp_cmd_direct := exec.Command("./binary")
	err = libjsonrpccpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: argtable")
exec.Command("latte", "install", "argtable")
	fmt.Println("Instalando dependencia: hiredis")
exec.Command("latte", "install", "hiredis")
	fmt.Println("Instalando dependencia: jsoncpp")
exec.Command("latte", "install", "jsoncpp")
	fmt.Println("Instalando dependencia: libmicrohttpd")
exec.Command("latte", "install", "libmicrohttpd")
}
