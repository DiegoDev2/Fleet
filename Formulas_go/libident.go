package main

// Libident - Ident protocol library
// Homepage: https://www.remlab.net/libident/

import (
	"fmt"
	
	"os/exec"
)

func installLibident() {
	// Método 1: Descargar y extraer .tar.gz
	libident_tar_url := "https://www.remlab.net/files/libident/libident-0.32.tar.gz"
	libident_cmd_tar := exec.Command("curl", "-L", libident_tar_url, "-o", "package.tar.gz")
	err := libident_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libident_zip_url := "https://www.remlab.net/files/libident/libident-0.32.zip"
	libident_cmd_zip := exec.Command("curl", "-L", libident_zip_url, "-o", "package.zip")
	err = libident_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libident_bin_url := "https://www.remlab.net/files/libident/libident-0.32.bin"
	libident_cmd_bin := exec.Command("curl", "-L", libident_bin_url, "-o", "binary.bin")
	err = libident_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libident_src_url := "https://www.remlab.net/files/libident/libident-0.32.src.tar.gz"
	libident_cmd_src := exec.Command("curl", "-L", libident_src_url, "-o", "source.tar.gz")
	err = libident_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libident_cmd_direct := exec.Command("./binary")
	err = libident_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
