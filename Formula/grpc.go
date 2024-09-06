package main

// Grpc - Next generation open source RPC library and framework
// Homepage: https://grpc.io/

import (
	"fmt"
	
	"os/exec"
)

func installGrpc() {
	// Método 1: Descargar y extraer .tar.gz
	grpc_tar_url := "https://github.com/grpc/grpc.git"
	grpc_cmd_tar := exec.Command("curl", "-L", grpc_tar_url, "-o", "package.tar.gz")
	err := grpc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grpc_zip_url := "https://github.com/grpc/grpc.git"
	grpc_cmd_zip := exec.Command("curl", "-L", grpc_zip_url, "-o", "package.zip")
	err = grpc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grpc_bin_url := "https://github.com/grpc/grpc.git"
	grpc_cmd_bin := exec.Command("curl", "-L", grpc_bin_url, "-o", "binary.bin")
	err = grpc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grpc_src_url := "https://github.com/grpc/grpc.git"
	grpc_cmd_src := exec.Command("curl", "-L", grpc_src_url, "-o", "source.tar.gz")
	err = grpc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grpc_cmd_direct := exec.Command("./binary")
	err = grpc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: c-ares")
	exec.Command("latte", "install", "c-ares").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: re2")
	exec.Command("latte", "install", "re2").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
