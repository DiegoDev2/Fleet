package main

// ProtocGenGo - Go support for Google's protocol buffers
// Homepage: https://github.com/protocolbuffers/protobuf-go

import (
	"fmt"
	
	"os/exec"
)

func installProtocGenGo() {
	// Método 1: Descargar y extraer .tar.gz
	protocgengo_tar_url := "https://github.com/protocolbuffers/protobuf-go/archive/refs/tags/v1.34.2.tar.gz"
	protocgengo_cmd_tar := exec.Command("curl", "-L", protocgengo_tar_url, "-o", "package.tar.gz")
	err := protocgengo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protocgengo_zip_url := "https://github.com/protocolbuffers/protobuf-go/archive/refs/tags/v1.34.2.zip"
	protocgengo_cmd_zip := exec.Command("curl", "-L", protocgengo_zip_url, "-o", "package.zip")
	err = protocgengo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protocgengo_bin_url := "https://github.com/protocolbuffers/protobuf-go/archive/refs/tags/v1.34.2.bin"
	protocgengo_cmd_bin := exec.Command("curl", "-L", protocgengo_bin_url, "-o", "binary.bin")
	err = protocgengo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protocgengo_src_url := "https://github.com/protocolbuffers/protobuf-go/archive/refs/tags/v1.34.2.src.tar.gz"
	protocgengo_cmd_src := exec.Command("curl", "-L", protocgengo_src_url, "-o", "source.tar.gz")
	err = protocgengo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protocgengo_cmd_direct := exec.Command("./binary")
	err = protocgengo_cmd_direct.Run()
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
