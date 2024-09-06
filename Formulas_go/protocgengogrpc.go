package main

// ProtocGenGoGrpc - Protoc plugin that generates code for gRPC-Go clients
// Homepage: https://github.com/grpc/grpc-go

import (
	"fmt"
	
	"os/exec"
)

func installProtocGenGoGrpc() {
	// Método 1: Descargar y extraer .tar.gz
	protocgengogrpc_tar_url := "https://github.com/grpc/grpc-go/archive/refs/tags/cmd/protoc-gen-go-grpc/v1.5.1.tar.gz"
	protocgengogrpc_cmd_tar := exec.Command("curl", "-L", protocgengogrpc_tar_url, "-o", "package.tar.gz")
	err := protocgengogrpc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protocgengogrpc_zip_url := "https://github.com/grpc/grpc-go/archive/refs/tags/cmd/protoc-gen-go-grpc/v1.5.1.zip"
	protocgengogrpc_cmd_zip := exec.Command("curl", "-L", protocgengogrpc_zip_url, "-o", "package.zip")
	err = protocgengogrpc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protocgengogrpc_bin_url := "https://github.com/grpc/grpc-go/archive/refs/tags/cmd/protoc-gen-go-grpc/v1.5.1.bin"
	protocgengogrpc_cmd_bin := exec.Command("curl", "-L", protocgengogrpc_bin_url, "-o", "binary.bin")
	err = protocgengogrpc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protocgengogrpc_src_url := "https://github.com/grpc/grpc-go/archive/refs/tags/cmd/protoc-gen-go-grpc/v1.5.1.src.tar.gz"
	protocgengogrpc_cmd_src := exec.Command("curl", "-L", protocgengogrpc_src_url, "-o", "source.tar.gz")
	err = protocgengogrpc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protocgengogrpc_cmd_direct := exec.Command("./binary")
	err = protocgengogrpc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
}
