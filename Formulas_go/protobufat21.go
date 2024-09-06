package main

// ProtobufAT21 - Protocol buffers (Google's data interchange format)
// Homepage: https://protobuf.dev/

import (
	"fmt"
	
	"os/exec"
)

func installProtobufAT21() {
	// Método 1: Descargar y extraer .tar.gz
	protobufat21_tar_url := "https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protobuf-all-21.12.tar.gz"
	protobufat21_cmd_tar := exec.Command("curl", "-L", protobufat21_tar_url, "-o", "package.tar.gz")
	err := protobufat21_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protobufat21_zip_url := "https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protobuf-all-21.12.zip"
	protobufat21_cmd_zip := exec.Command("curl", "-L", protobufat21_zip_url, "-o", "package.zip")
	err = protobufat21_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protobufat21_bin_url := "https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protobuf-all-21.12.bin"
	protobufat21_cmd_bin := exec.Command("curl", "-L", protobufat21_bin_url, "-o", "binary.bin")
	err = protobufat21_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protobufat21_src_url := "https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protobuf-all-21.12.src.tar.gz"
	protobufat21_cmd_src := exec.Command("curl", "-L", protobufat21_src_url, "-o", "source.tar.gz")
	err = protobufat21_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protobufat21_cmd_direct := exec.Command("./binary")
	err = protobufat21_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
