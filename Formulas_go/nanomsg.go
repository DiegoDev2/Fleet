package main

// Nanomsg - Socket library in C
// Homepage: https://nanomsg.org/

import (
	"fmt"
	
	"os/exec"
)

func installNanomsg() {
	// Método 1: Descargar y extraer .tar.gz
	nanomsg_tar_url := "https://github.com/nanomsg/nanomsg/archive/refs/tags/1.2.tar.gz"
	nanomsg_cmd_tar := exec.Command("curl", "-L", nanomsg_tar_url, "-o", "package.tar.gz")
	err := nanomsg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nanomsg_zip_url := "https://github.com/nanomsg/nanomsg/archive/refs/tags/1.2.zip"
	nanomsg_cmd_zip := exec.Command("curl", "-L", nanomsg_zip_url, "-o", "package.zip")
	err = nanomsg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nanomsg_bin_url := "https://github.com/nanomsg/nanomsg/archive/refs/tags/1.2.bin"
	nanomsg_cmd_bin := exec.Command("curl", "-L", nanomsg_bin_url, "-o", "binary.bin")
	err = nanomsg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nanomsg_src_url := "https://github.com/nanomsg/nanomsg/archive/refs/tags/1.2.src.tar.gz"
	nanomsg_cmd_src := exec.Command("curl", "-L", nanomsg_src_url, "-o", "source.tar.gz")
	err = nanomsg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nanomsg_cmd_direct := exec.Command("./binary")
	err = nanomsg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
