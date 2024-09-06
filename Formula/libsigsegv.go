package main

// Libsigsegv - Library for handling page faults in user mode
// Homepage: https://www.gnu.org/software/libsigsegv/

import (
	"fmt"
	
	"os/exec"
)

func installLibsigsegv() {
	// Método 1: Descargar y extraer .tar.gz
	libsigsegv_tar_url := "https://ftp.gnu.org/gnu/libsigsegv/libsigsegv-2.14.tar.gz"
	libsigsegv_cmd_tar := exec.Command("curl", "-L", libsigsegv_tar_url, "-o", "package.tar.gz")
	err := libsigsegv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsigsegv_zip_url := "https://ftp.gnu.org/gnu/libsigsegv/libsigsegv-2.14.zip"
	libsigsegv_cmd_zip := exec.Command("curl", "-L", libsigsegv_zip_url, "-o", "package.zip")
	err = libsigsegv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsigsegv_bin_url := "https://ftp.gnu.org/gnu/libsigsegv/libsigsegv-2.14.bin"
	libsigsegv_cmd_bin := exec.Command("curl", "-L", libsigsegv_bin_url, "-o", "binary.bin")
	err = libsigsegv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsigsegv_src_url := "https://ftp.gnu.org/gnu/libsigsegv/libsigsegv-2.14.src.tar.gz"
	libsigsegv_cmd_src := exec.Command("curl", "-L", libsigsegv_src_url, "-o", "source.tar.gz")
	err = libsigsegv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsigsegv_cmd_direct := exec.Command("./binary")
	err = libsigsegv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
}
