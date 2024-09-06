package main

// GrpcSwift - Swift language implementation of gRPC
// Homepage: https://github.com/grpc/grpc-swift

import (
	"fmt"
	
	"os/exec"
)

func installGrpcSwift() {
	// Método 1: Descargar y extraer .tar.gz
	grpcswift_tar_url := "https://github.com/grpc/grpc-swift/archive/refs/tags/1.23.0.tar.gz"
	grpcswift_cmd_tar := exec.Command("curl", "-L", grpcswift_tar_url, "-o", "package.tar.gz")
	err := grpcswift_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grpcswift_zip_url := "https://github.com/grpc/grpc-swift/archive/refs/tags/1.23.0.zip"
	grpcswift_cmd_zip := exec.Command("curl", "-L", grpcswift_zip_url, "-o", "package.zip")
	err = grpcswift_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grpcswift_bin_url := "https://github.com/grpc/grpc-swift/archive/refs/tags/1.23.0.bin"
	grpcswift_cmd_bin := exec.Command("curl", "-L", grpcswift_bin_url, "-o", "binary.bin")
	err = grpcswift_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grpcswift_src_url := "https://github.com/grpc/grpc-swift/archive/refs/tags/1.23.0.src.tar.gz"
	grpcswift_cmd_src := exec.Command("curl", "-L", grpcswift_src_url, "-o", "source.tar.gz")
	err = grpcswift_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grpcswift_cmd_direct := exec.Command("./binary")
	err = grpcswift_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: swift-protobuf")
exec.Command("latte", "install", "swift-protobuf")
}
