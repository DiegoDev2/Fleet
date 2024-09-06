package main

// ProtocGenGrpcWeb - Protoc plugin that generates code for gRPC-Web clients
// Homepage: https://github.com/grpc/grpc-web

import (
	"fmt"
	
	"os/exec"
)

func installProtocGenGrpcWeb() {
	// Método 1: Descargar y extraer .tar.gz
	protocgengrpcweb_tar_url := "https://github.com/grpc/grpc-web/archive/refs/tags/1.5.0.tar.gz"
	protocgengrpcweb_cmd_tar := exec.Command("curl", "-L", protocgengrpcweb_tar_url, "-o", "package.tar.gz")
	err := protocgengrpcweb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protocgengrpcweb_zip_url := "https://github.com/grpc/grpc-web/archive/refs/tags/1.5.0.zip"
	protocgengrpcweb_cmd_zip := exec.Command("curl", "-L", protocgengrpcweb_zip_url, "-o", "package.zip")
	err = protocgengrpcweb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protocgengrpcweb_bin_url := "https://github.com/grpc/grpc-web/archive/refs/tags/1.5.0.bin"
	protocgengrpcweb_cmd_bin := exec.Command("curl", "-L", protocgengrpcweb_bin_url, "-o", "binary.bin")
	err = protocgengrpcweb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protocgengrpcweb_src_url := "https://github.com/grpc/grpc-web/archive/refs/tags/1.5.0.src.tar.gz"
	protocgengrpcweb_cmd_src := exec.Command("curl", "-L", protocgengrpcweb_src_url, "-o", "source.tar.gz")
	err = protocgengrpcweb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protocgengrpcweb_cmd_direct := exec.Command("./binary")
	err = protocgengrpcweb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: typescript")
exec.Command("latte", "install", "typescript")
	fmt.Println("Instalando dependencia: protobuf@21")
exec.Command("latte", "install", "protobuf@21")
	fmt.Println("Instalando dependencia: protoc-gen-js")
exec.Command("latte", "install", "protoc-gen-js")
}
