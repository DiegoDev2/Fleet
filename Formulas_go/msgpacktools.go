package main

// MsgpackTools - Command-line tools for converting between MessagePack and JSON
// Homepage: https://github.com/ludocode/msgpack-tools

import (
	"fmt"
	
	"os/exec"
)

func installMsgpackTools() {
	// Método 1: Descargar y extraer .tar.gz
	msgpacktools_tar_url := "https://github.com/ludocode/msgpack-tools/releases/download/v0.6/msgpack-tools-0.6.tar.gz"
	msgpacktools_cmd_tar := exec.Command("curl", "-L", msgpacktools_tar_url, "-o", "package.tar.gz")
	err := msgpacktools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	msgpacktools_zip_url := "https://github.com/ludocode/msgpack-tools/releases/download/v0.6/msgpack-tools-0.6.zip"
	msgpacktools_cmd_zip := exec.Command("curl", "-L", msgpacktools_zip_url, "-o", "package.zip")
	err = msgpacktools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	msgpacktools_bin_url := "https://github.com/ludocode/msgpack-tools/releases/download/v0.6/msgpack-tools-0.6.bin"
	msgpacktools_cmd_bin := exec.Command("curl", "-L", msgpacktools_bin_url, "-o", "binary.bin")
	err = msgpacktools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	msgpacktools_src_url := "https://github.com/ludocode/msgpack-tools/releases/download/v0.6/msgpack-tools-0.6.src.tar.gz"
	msgpacktools_cmd_src := exec.Command("curl", "-L", msgpacktools_src_url, "-o", "source.tar.gz")
	err = msgpacktools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	msgpacktools_cmd_direct := exec.Command("./binary")
	err = msgpacktools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
