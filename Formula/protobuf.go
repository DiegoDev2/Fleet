package main

// Protobuf - Protocol buffers (Google's data interchange format)
// Homepage: https://protobuf.dev/

import (
	"fmt"
	
	"os/exec"
)

func installProtobuf() {
	// Método 1: Descargar y extraer .tar.gz
	protobuf_tar_url := "https://github.com/protocolbuffers/protobuf/releases/download/v28.0/protobuf-28.0.tar.gz"
	protobuf_cmd_tar := exec.Command("curl", "-L", protobuf_tar_url, "-o", "package.tar.gz")
	err := protobuf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protobuf_zip_url := "https://github.com/protocolbuffers/protobuf/releases/download/v28.0/protobuf-28.0.zip"
	protobuf_cmd_zip := exec.Command("curl", "-L", protobuf_zip_url, "-o", "package.zip")
	err = protobuf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protobuf_bin_url := "https://github.com/protocolbuffers/protobuf/releases/download/v28.0/protobuf-28.0.bin"
	protobuf_cmd_bin := exec.Command("curl", "-L", protobuf_bin_url, "-o", "binary.bin")
	err = protobuf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protobuf_src_url := "https://github.com/protocolbuffers/protobuf/releases/download/v28.0/protobuf-28.0.src.tar.gz"
	protobuf_cmd_src := exec.Command("curl", "-L", protobuf_src_url, "-o", "source.tar.gz")
	err = protobuf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protobuf_cmd_direct := exec.Command("./binary")
	err = protobuf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: googletest")
	exec.Command("latte", "install", "googletest").Run()
}
