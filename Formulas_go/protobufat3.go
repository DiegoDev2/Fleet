package main

// ProtobufAT3 - Protocol buffers (Google's data interchange format)
// Homepage: https://github.com/protocolbuffers/protobuf/

import (
	"fmt"
	
	"os/exec"
)

func installProtobufAT3() {
	// Método 1: Descargar y extraer .tar.gz
	protobufat3_tar_url := "https://github.com/protocolbuffers/protobuf/releases/download/v3.20.3/protobuf-all-3.20.3.tar.gz"
	protobufat3_cmd_tar := exec.Command("curl", "-L", protobufat3_tar_url, "-o", "package.tar.gz")
	err := protobufat3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protobufat3_zip_url := "https://github.com/protocolbuffers/protobuf/releases/download/v3.20.3/protobuf-all-3.20.3.zip"
	protobufat3_cmd_zip := exec.Command("curl", "-L", protobufat3_zip_url, "-o", "package.zip")
	err = protobufat3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protobufat3_bin_url := "https://github.com/protocolbuffers/protobuf/releases/download/v3.20.3/protobuf-all-3.20.3.bin"
	protobufat3_cmd_bin := exec.Command("curl", "-L", protobufat3_bin_url, "-o", "binary.bin")
	err = protobufat3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protobufat3_src_url := "https://github.com/protocolbuffers/protobuf/releases/download/v3.20.3/protobuf-all-3.20.3.src.tar.gz"
	protobufat3_cmd_src := exec.Command("curl", "-L", protobufat3_src_url, "-o", "source.tar.gz")
	err = protobufat3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protobufat3_cmd_direct := exec.Command("./binary")
	err = protobufat3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
