package main

// Libxdiff - Implements diff functions for binary and text files
// Homepage: http://www.xmailserver.org/xdiff-lib.html

import (
	"fmt"
	
	"os/exec"
)

func installLibxdiff() {
	// Método 1: Descargar y extraer .tar.gz
	libxdiff_tar_url := "http://www.xmailserver.org/libxdiff-0.23.tar.gz"
	libxdiff_cmd_tar := exec.Command("curl", "-L", libxdiff_tar_url, "-o", "package.tar.gz")
	err := libxdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxdiff_zip_url := "http://www.xmailserver.org/libxdiff-0.23.zip"
	libxdiff_cmd_zip := exec.Command("curl", "-L", libxdiff_zip_url, "-o", "package.zip")
	err = libxdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxdiff_bin_url := "http://www.xmailserver.org/libxdiff-0.23.bin"
	libxdiff_cmd_bin := exec.Command("curl", "-L", libxdiff_bin_url, "-o", "binary.bin")
	err = libxdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxdiff_src_url := "http://www.xmailserver.org/libxdiff-0.23.src.tar.gz"
	libxdiff_cmd_src := exec.Command("curl", "-L", libxdiff_src_url, "-o", "source.tar.gz")
	err = libxdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxdiff_cmd_direct := exec.Command("./binary")
	err = libxdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
