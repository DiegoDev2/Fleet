package main

// Libmowgli - Core framework for Atheme applications
// Homepage: https://github.com/atheme/libmowgli-2

import (
	"fmt"
	
	"os/exec"
)

func installLibmowgli() {
	// Método 1: Descargar y extraer .tar.gz
	libmowgli_tar_url := "https://github.com/atheme/libmowgli-2/archive/refs/tags/v2.1.3.tar.gz"
	libmowgli_cmd_tar := exec.Command("curl", "-L", libmowgli_tar_url, "-o", "package.tar.gz")
	err := libmowgli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmowgli_zip_url := "https://github.com/atheme/libmowgli-2/archive/refs/tags/v2.1.3.zip"
	libmowgli_cmd_zip := exec.Command("curl", "-L", libmowgli_zip_url, "-o", "package.zip")
	err = libmowgli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmowgli_bin_url := "https://github.com/atheme/libmowgli-2/archive/refs/tags/v2.1.3.bin"
	libmowgli_cmd_bin := exec.Command("curl", "-L", libmowgli_bin_url, "-o", "binary.bin")
	err = libmowgli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmowgli_src_url := "https://github.com/atheme/libmowgli-2/archive/refs/tags/v2.1.3.src.tar.gz"
	libmowgli_cmd_src := exec.Command("curl", "-L", libmowgli_src_url, "-o", "source.tar.gz")
	err = libmowgli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmowgli_cmd_direct := exec.Command("./binary")
	err = libmowgli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
}
