package main

// Libcello - Higher-level programming in C
// Homepage: https://libcello.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibcello() {
	// Método 1: Descargar y extraer .tar.gz
	libcello_tar_url := "https://libcello.org/static/libCello-2.1.0.tar.gz"
	libcello_cmd_tar := exec.Command("curl", "-L", libcello_tar_url, "-o", "package.tar.gz")
	err := libcello_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcello_zip_url := "https://libcello.org/static/libCello-2.1.0.zip"
	libcello_cmd_zip := exec.Command("curl", "-L", libcello_zip_url, "-o", "package.zip")
	err = libcello_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcello_bin_url := "https://libcello.org/static/libCello-2.1.0.bin"
	libcello_cmd_bin := exec.Command("curl", "-L", libcello_bin_url, "-o", "binary.bin")
	err = libcello_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcello_src_url := "https://libcello.org/static/libCello-2.1.0.src.tar.gz"
	libcello_cmd_src := exec.Command("curl", "-L", libcello_src_url, "-o", "source.tar.gz")
	err = libcello_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcello_cmd_direct := exec.Command("./binary")
	err = libcello_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
