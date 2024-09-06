package main

// ProtobufC - Protocol buffers library
// Homepage: https://github.com/protobuf-c/protobuf-c

import (
	"fmt"
	
	"os/exec"
)

func installProtobufC() {
	// Método 1: Descargar y extraer .tar.gz
	protobufc_tar_url := "https://github.com/protobuf-c/protobuf-c/releases/download/v1.5.0/protobuf-c-1.5.0.tar.gz"
	protobufc_cmd_tar := exec.Command("curl", "-L", protobufc_tar_url, "-o", "package.tar.gz")
	err := protobufc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protobufc_zip_url := "https://github.com/protobuf-c/protobuf-c/releases/download/v1.5.0/protobuf-c-1.5.0.zip"
	protobufc_cmd_zip := exec.Command("curl", "-L", protobufc_zip_url, "-o", "package.zip")
	err = protobufc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protobufc_bin_url := "https://github.com/protobuf-c/protobuf-c/releases/download/v1.5.0/protobuf-c-1.5.0.bin"
	protobufc_cmd_bin := exec.Command("curl", "-L", protobufc_bin_url, "-o", "binary.bin")
	err = protobufc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protobufc_src_url := "https://github.com/protobuf-c/protobuf-c/releases/download/v1.5.0/protobuf-c-1.5.0.src.tar.gz"
	protobufc_cmd_src := exec.Command("curl", "-L", protobufc_src_url, "-o", "source.tar.gz")
	err = protobufc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protobufc_cmd_direct := exec.Command("./binary")
	err = protobufc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
	exec.Command("latte", "install", "asciidoc").Run()
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
}
