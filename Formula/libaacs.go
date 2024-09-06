package main

// Libaacs - Implements the Advanced Access Content System specification
// Homepage: https://www.videolan.org/developers/libaacs.html

import (
	"fmt"
	
	"os/exec"
)

func installLibaacs() {
	// Método 1: Descargar y extraer .tar.gz
	libaacs_tar_url := "https://get.videolan.org/libaacs/0.11.1/libaacs-0.11.1.tar.bz2"
	libaacs_cmd_tar := exec.Command("curl", "-L", libaacs_tar_url, "-o", "package.tar.gz")
	err := libaacs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libaacs_zip_url := "https://get.videolan.org/libaacs/0.11.1/libaacs-0.11.1.tar.bz2"
	libaacs_cmd_zip := exec.Command("curl", "-L", libaacs_zip_url, "-o", "package.zip")
	err = libaacs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libaacs_bin_url := "https://get.videolan.org/libaacs/0.11.1/libaacs-0.11.1.tar.bz2"
	libaacs_cmd_bin := exec.Command("curl", "-L", libaacs_bin_url, "-o", "binary.bin")
	err = libaacs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libaacs_src_url := "https://get.videolan.org/libaacs/0.11.1/libaacs-0.11.1.tar.bz2"
	libaacs_cmd_src := exec.Command("curl", "-L", libaacs_src_url, "-o", "source.tar.gz")
	err = libaacs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libaacs_cmd_direct := exec.Command("./binary")
	err = libaacs_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
}
