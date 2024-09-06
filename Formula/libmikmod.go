package main

// Libmikmod - Portable sound library
// Homepage: https://mikmod.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibmikmod() {
	// Método 1: Descargar y extraer .tar.gz
	libmikmod_tar_url := "https://downloads.sourceforge.net/project/mikmod/libmikmod/3.3.11.1/libmikmod-3.3.11.1.tar.gz"
	libmikmod_cmd_tar := exec.Command("curl", "-L", libmikmod_tar_url, "-o", "package.tar.gz")
	err := libmikmod_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmikmod_zip_url := "https://downloads.sourceforge.net/project/mikmod/libmikmod/3.3.11.1/libmikmod-3.3.11.1.zip"
	libmikmod_cmd_zip := exec.Command("curl", "-L", libmikmod_zip_url, "-o", "package.zip")
	err = libmikmod_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmikmod_bin_url := "https://downloads.sourceforge.net/project/mikmod/libmikmod/3.3.11.1/libmikmod-3.3.11.1.bin"
	libmikmod_cmd_bin := exec.Command("curl", "-L", libmikmod_bin_url, "-o", "binary.bin")
	err = libmikmod_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmikmod_src_url := "https://downloads.sourceforge.net/project/mikmod/libmikmod/3.3.11.1/libmikmod-3.3.11.1.src.tar.gz"
	libmikmod_cmd_src := exec.Command("curl", "-L", libmikmod_src_url, "-o", "source.tar.gz")
	err = libmikmod_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmikmod_cmd_direct := exec.Command("./binary")
	err = libmikmod_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
