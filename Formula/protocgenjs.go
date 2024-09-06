package main

// ProtocGenJs - Protocol buffers JavaScript generator plugin
// Homepage: https://github.com/protocolbuffers/protobuf-javascript

import (
	"fmt"
	
	"os/exec"
)

func installProtocGenJs() {
	// Método 1: Descargar y extraer .tar.gz
	protocgenjs_tar_url := "https://github.com/protocolbuffers/protobuf-javascript/archive/refs/tags/v3.21.2.tar.gz"
	protocgenjs_cmd_tar := exec.Command("curl", "-L", protocgenjs_tar_url, "-o", "package.tar.gz")
	err := protocgenjs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protocgenjs_zip_url := "https://github.com/protocolbuffers/protobuf-javascript/archive/refs/tags/v3.21.2.zip"
	protocgenjs_cmd_zip := exec.Command("curl", "-L", protocgenjs_zip_url, "-o", "package.zip")
	err = protocgenjs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protocgenjs_bin_url := "https://github.com/protocolbuffers/protobuf-javascript/archive/refs/tags/v3.21.2.bin"
	protocgenjs_cmd_bin := exec.Command("curl", "-L", protocgenjs_bin_url, "-o", "binary.bin")
	err = protocgenjs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protocgenjs_src_url := "https://github.com/protocolbuffers/protobuf-javascript/archive/refs/tags/v3.21.2.src.tar.gz"
	protocgenjs_cmd_src := exec.Command("curl", "-L", protocgenjs_src_url, "-o", "source.tar.gz")
	err = protocgenjs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protocgenjs_cmd_direct := exec.Command("./binary")
	err = protocgenjs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bazelisk")
	exec.Command("latte", "install", "bazelisk").Run()
	fmt.Println("Instalando dependencia: protobuf@21")
	exec.Command("latte", "install", "protobuf@21").Run()
}
