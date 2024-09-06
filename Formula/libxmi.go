package main

// Libxmi - C/C++ function library for rasterizing 2D vector graphics
// Homepage: https://www.gnu.org/software/libxmi/

import (
	"fmt"
	
	"os/exec"
)

func installLibxmi() {
	// Método 1: Descargar y extraer .tar.gz
	libxmi_tar_url := "https://ftp.gnu.org/gnu/libxmi/libxmi-1.2.tar.gz"
	libxmi_cmd_tar := exec.Command("curl", "-L", libxmi_tar_url, "-o", "package.tar.gz")
	err := libxmi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxmi_zip_url := "https://ftp.gnu.org/gnu/libxmi/libxmi-1.2.zip"
	libxmi_cmd_zip := exec.Command("curl", "-L", libxmi_zip_url, "-o", "package.zip")
	err = libxmi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxmi_bin_url := "https://ftp.gnu.org/gnu/libxmi/libxmi-1.2.bin"
	libxmi_cmd_bin := exec.Command("curl", "-L", libxmi_bin_url, "-o", "binary.bin")
	err = libxmi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxmi_src_url := "https://ftp.gnu.org/gnu/libxmi/libxmi-1.2.src.tar.gz"
	libxmi_cmd_src := exec.Command("curl", "-L", libxmi_src_url, "-o", "source.tar.gz")
	err = libxmi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxmi_cmd_direct := exec.Command("./binary")
	err = libxmi_cmd_direct.Run()
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
