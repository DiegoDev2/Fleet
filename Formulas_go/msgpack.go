package main

// Msgpack - Library for a binary-based efficient data interchange format
// Homepage: https://msgpack.org/

import (
	"fmt"
	
	"os/exec"
)

func installMsgpack() {
	// Método 1: Descargar y extraer .tar.gz
	msgpack_tar_url := "https://github.com/msgpack/msgpack-c/releases/download/c-6.1.0/msgpack-c-6.1.0.tar.gz"
	msgpack_cmd_tar := exec.Command("curl", "-L", msgpack_tar_url, "-o", "package.tar.gz")
	err := msgpack_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	msgpack_zip_url := "https://github.com/msgpack/msgpack-c/releases/download/c-6.1.0/msgpack-c-6.1.0.zip"
	msgpack_cmd_zip := exec.Command("curl", "-L", msgpack_zip_url, "-o", "package.zip")
	err = msgpack_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	msgpack_bin_url := "https://github.com/msgpack/msgpack-c/releases/download/c-6.1.0/msgpack-c-6.1.0.bin"
	msgpack_cmd_bin := exec.Command("curl", "-L", msgpack_bin_url, "-o", "binary.bin")
	err = msgpack_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	msgpack_src_url := "https://github.com/msgpack/msgpack-c/releases/download/c-6.1.0/msgpack-c-6.1.0.src.tar.gz"
	msgpack_cmd_src := exec.Command("curl", "-L", msgpack_src_url, "-o", "source.tar.gz")
	err = msgpack_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	msgpack_cmd_direct := exec.Command("./binary")
	err = msgpack_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
