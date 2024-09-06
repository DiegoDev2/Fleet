package main

// SwiftProtobuf - Plugin and runtime library for using protobuf with Swift
// Homepage: https://github.com/apple/swift-protobuf

import (
	"fmt"
	
	"os/exec"
)

func installSwiftProtobuf() {
	// Método 1: Descargar y extraer .tar.gz
	swiftprotobuf_tar_url := "https://github.com/apple/swift-protobuf/archive/refs/tags/1.28.1.tar.gz"
	swiftprotobuf_cmd_tar := exec.Command("curl", "-L", swiftprotobuf_tar_url, "-o", "package.tar.gz")
	err := swiftprotobuf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	swiftprotobuf_zip_url := "https://github.com/apple/swift-protobuf/archive/refs/tags/1.28.1.zip"
	swiftprotobuf_cmd_zip := exec.Command("curl", "-L", swiftprotobuf_zip_url, "-o", "package.zip")
	err = swiftprotobuf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	swiftprotobuf_bin_url := "https://github.com/apple/swift-protobuf/archive/refs/tags/1.28.1.bin"
	swiftprotobuf_cmd_bin := exec.Command("curl", "-L", swiftprotobuf_bin_url, "-o", "binary.bin")
	err = swiftprotobuf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	swiftprotobuf_src_url := "https://github.com/apple/swift-protobuf/archive/refs/tags/1.28.1.src.tar.gz"
	swiftprotobuf_cmd_src := exec.Command("curl", "-L", swiftprotobuf_src_url, "-o", "source.tar.gz")
	err = swiftprotobuf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	swiftprotobuf_cmd_direct := exec.Command("./binary")
	err = swiftprotobuf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
}
