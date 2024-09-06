package main

// Inetutils - GNU utilities for networking
// Homepage: https://www.gnu.org/software/inetutils/

import (
	"fmt"
	
	"os/exec"
)

func installInetutils() {
	// Método 1: Descargar y extraer .tar.gz
	inetutils_tar_url := "https://ftp.gnu.org/gnu/inetutils/inetutils-2.5.tar.xz"
	inetutils_cmd_tar := exec.Command("curl", "-L", inetutils_tar_url, "-o", "package.tar.gz")
	err := inetutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	inetutils_zip_url := "https://ftp.gnu.org/gnu/inetutils/inetutils-2.5.tar.xz"
	inetutils_cmd_zip := exec.Command("curl", "-L", inetutils_zip_url, "-o", "package.zip")
	err = inetutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	inetutils_bin_url := "https://ftp.gnu.org/gnu/inetutils/inetutils-2.5.tar.xz"
	inetutils_cmd_bin := exec.Command("curl", "-L", inetutils_bin_url, "-o", "binary.bin")
	err = inetutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	inetutils_src_url := "https://ftp.gnu.org/gnu/inetutils/inetutils-2.5.tar.xz"
	inetutils_cmd_src := exec.Command("curl", "-L", inetutils_src_url, "-o", "source.tar.gz")
	err = inetutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	inetutils_cmd_direct := exec.Command("./binary")
	err = inetutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: help2man")
exec.Command("latte", "install", "help2man")
	fmt.Println("Instalando dependencia: libidn2")
exec.Command("latte", "install", "libidn2")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
