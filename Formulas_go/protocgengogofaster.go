package main

// ProtocGenGogofaster - Protocol Buffers for Go with Gadgets
// Homepage: https://github.com/gogo/protobuf

import (
	"fmt"
	
	"os/exec"
)

func installProtocGenGogofaster() {
	// Método 1: Descargar y extraer .tar.gz
	protocgengogofaster_tar_url := "https://github.com/gogo/protobuf/archive/refs/tags/v1.3.2.tar.gz"
	protocgengogofaster_cmd_tar := exec.Command("curl", "-L", protocgengogofaster_tar_url, "-o", "package.tar.gz")
	err := protocgengogofaster_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	protocgengogofaster_zip_url := "https://github.com/gogo/protobuf/archive/refs/tags/v1.3.2.zip"
	protocgengogofaster_cmd_zip := exec.Command("curl", "-L", protocgengogofaster_zip_url, "-o", "package.zip")
	err = protocgengogofaster_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	protocgengogofaster_bin_url := "https://github.com/gogo/protobuf/archive/refs/tags/v1.3.2.bin"
	protocgengogofaster_cmd_bin := exec.Command("curl", "-L", protocgengogofaster_bin_url, "-o", "binary.bin")
	err = protocgengogofaster_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	protocgengogofaster_src_url := "https://github.com/gogo/protobuf/archive/refs/tags/v1.3.2.src.tar.gz"
	protocgengogofaster_cmd_src := exec.Command("curl", "-L", protocgengogofaster_src_url, "-o", "source.tar.gz")
	err = protocgengogofaster_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	protocgengogofaster_cmd_direct := exec.Command("./binary")
	err = protocgengogofaster_cmd_direct.Run()
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
