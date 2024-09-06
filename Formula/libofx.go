package main

// Libofx - Library to support OFX command responses
// Homepage: https://github.com/libofx/libofx

import (
	"fmt"
	
	"os/exec"
)

func installLibofx() {
	// Método 1: Descargar y extraer .tar.gz
	libofx_tar_url := "https://github.com/libofx/libofx/releases/download/0.10.9/libofx-0.10.9.tar.gz"
	libofx_cmd_tar := exec.Command("curl", "-L", libofx_tar_url, "-o", "package.tar.gz")
	err := libofx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libofx_zip_url := "https://github.com/libofx/libofx/releases/download/0.10.9/libofx-0.10.9.zip"
	libofx_cmd_zip := exec.Command("curl", "-L", libofx_zip_url, "-o", "package.zip")
	err = libofx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libofx_bin_url := "https://github.com/libofx/libofx/releases/download/0.10.9/libofx-0.10.9.bin"
	libofx_cmd_bin := exec.Command("curl", "-L", libofx_bin_url, "-o", "binary.bin")
	err = libofx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libofx_src_url := "https://github.com/libofx/libofx/releases/download/0.10.9/libofx-0.10.9.src.tar.gz"
	libofx_cmd_src := exec.Command("curl", "-L", libofx_src_url, "-o", "source.tar.gz")
	err = libofx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libofx_cmd_direct := exec.Command("./binary")
	err = libofx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: gengetopt")
	exec.Command("latte", "install", "gengetopt").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: open-sp")
	exec.Command("latte", "install", "open-sp").Run()
}
