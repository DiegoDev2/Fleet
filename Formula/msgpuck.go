package main

// Msgpuck - Simple and efficient MsgPack binary serialization library
// Homepage: https://rtsisyk.github.io/msgpuck/

import (
	"fmt"
	
	"os/exec"
)

func installMsgpuck() {
	// Método 1: Descargar y extraer .tar.gz
	msgpuck_tar_url := "https://github.com/rtsisyk/msgpuck/archive/refs/tags/2.0.tar.gz"
	msgpuck_cmd_tar := exec.Command("curl", "-L", msgpuck_tar_url, "-o", "package.tar.gz")
	err := msgpuck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	msgpuck_zip_url := "https://github.com/rtsisyk/msgpuck/archive/refs/tags/2.0.zip"
	msgpuck_cmd_zip := exec.Command("curl", "-L", msgpuck_zip_url, "-o", "package.zip")
	err = msgpuck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	msgpuck_bin_url := "https://github.com/rtsisyk/msgpuck/archive/refs/tags/2.0.bin"
	msgpuck_cmd_bin := exec.Command("curl", "-L", msgpuck_bin_url, "-o", "binary.bin")
	err = msgpuck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	msgpuck_src_url := "https://github.com/rtsisyk/msgpuck/archive/refs/tags/2.0.src.tar.gz"
	msgpuck_cmd_src := exec.Command("curl", "-L", msgpuck_src_url, "-o", "source.tar.gz")
	err = msgpuck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	msgpuck_cmd_direct := exec.Command("./binary")
	err = msgpuck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
