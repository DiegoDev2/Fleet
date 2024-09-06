package main

// GlibcAT213 - GNU C Library
// Homepage: https://www.gnu.org/software/libc/

import (
	"fmt"
	
	"os/exec"
)

func installGlibcAT213() {
	// Método 1: Descargar y extraer .tar.gz
	glibcat213_tar_url := "https://ftp.gnu.org/gnu/glibc/glibc-2.13.tar.gz"
	glibcat213_cmd_tar := exec.Command("curl", "-L", glibcat213_tar_url, "-o", "package.tar.gz")
	err := glibcat213_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glibcat213_zip_url := "https://ftp.gnu.org/gnu/glibc/glibc-2.13.zip"
	glibcat213_cmd_zip := exec.Command("curl", "-L", glibcat213_zip_url, "-o", "package.zip")
	err = glibcat213_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glibcat213_bin_url := "https://ftp.gnu.org/gnu/glibc/glibc-2.13.bin"
	glibcat213_cmd_bin := exec.Command("curl", "-L", glibcat213_bin_url, "-o", "binary.bin")
	err = glibcat213_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glibcat213_src_url := "https://ftp.gnu.org/gnu/glibc/glibc-2.13.src.tar.gz"
	glibcat213_cmd_src := exec.Command("curl", "-L", glibcat213_src_url, "-o", "source.tar.gz")
	err = glibcat213_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glibcat213_cmd_direct := exec.Command("./binary")
	err = glibcat213_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: linux-headers@4.4")
exec.Command("latte", "install", "linux-headers@4.4")
}
