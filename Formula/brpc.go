package main

// Brpc - Better RPC framework
// Homepage: https://brpc.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installBrpc() {
	// Método 1: Descargar y extraer .tar.gz
	brpc_tar_url := "https://dlcdn.apache.org/brpc/1.10.0/apache-brpc-1.10.0-src.tar.gz"
	brpc_cmd_tar := exec.Command("curl", "-L", brpc_tar_url, "-o", "package.tar.gz")
	err := brpc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brpc_zip_url := "https://dlcdn.apache.org/brpc/1.10.0/apache-brpc-1.10.0-src.zip"
	brpc_cmd_zip := exec.Command("curl", "-L", brpc_zip_url, "-o", "package.zip")
	err = brpc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brpc_bin_url := "https://dlcdn.apache.org/brpc/1.10.0/apache-brpc-1.10.0-src.bin"
	brpc_cmd_bin := exec.Command("curl", "-L", brpc_bin_url, "-o", "binary.bin")
	err = brpc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brpc_src_url := "https://dlcdn.apache.org/brpc/1.10.0/apache-brpc-1.10.0-src.src.tar.gz"
	brpc_cmd_src := exec.Command("curl", "-L", brpc_src_url, "-o", "source.tar.gz")
	err = brpc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brpc_cmd_direct := exec.Command("./binary")
	err = brpc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
	fmt.Println("Instalando dependencia: gperftools")
	exec.Command("latte", "install", "gperftools").Run()
	fmt.Println("Instalando dependencia: leveldb")
	exec.Command("latte", "install", "leveldb").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: protobuf@21")
	exec.Command("latte", "install", "protobuf@21").Run()
}
