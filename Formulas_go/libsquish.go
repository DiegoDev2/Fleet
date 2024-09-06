package main

// Libsquish - Library for compressing images with the DXT standard
// Homepage: https://sourceforge.net/projects/libsquish/

import (
	"fmt"
	
	"os/exec"
)

func installLibsquish() {
	// Método 1: Descargar y extraer .tar.gz
	libsquish_tar_url := "https://downloads.sourceforge.net/project/libsquish/libsquish-1.15.tgz"
	libsquish_cmd_tar := exec.Command("curl", "-L", libsquish_tar_url, "-o", "package.tar.gz")
	err := libsquish_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsquish_zip_url := "https://downloads.sourceforge.net/project/libsquish/libsquish-1.15.tgz"
	libsquish_cmd_zip := exec.Command("curl", "-L", libsquish_zip_url, "-o", "package.zip")
	err = libsquish_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsquish_bin_url := "https://downloads.sourceforge.net/project/libsquish/libsquish-1.15.tgz"
	libsquish_cmd_bin := exec.Command("curl", "-L", libsquish_bin_url, "-o", "binary.bin")
	err = libsquish_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsquish_src_url := "https://downloads.sourceforge.net/project/libsquish/libsquish-1.15.tgz"
	libsquish_cmd_src := exec.Command("curl", "-L", libsquish_src_url, "-o", "source.tar.gz")
	err = libsquish_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsquish_cmd_direct := exec.Command("./binary")
	err = libsquish_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
