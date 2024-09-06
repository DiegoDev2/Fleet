package main

// Libconfini - Yet another INI parser
// Homepage: https://madmurphy.github.io/libconfini/

import (
	"fmt"
	
	"os/exec"
)

func installLibconfini() {
	// Método 1: Descargar y extraer .tar.gz
	libconfini_tar_url := "https://github.com/madmurphy/libconfini/releases/download/1.16.4/libconfini-1.16.4-with-configure.tar.gz"
	libconfini_cmd_tar := exec.Command("curl", "-L", libconfini_tar_url, "-o", "package.tar.gz")
	err := libconfini_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libconfini_zip_url := "https://github.com/madmurphy/libconfini/releases/download/1.16.4/libconfini-1.16.4-with-configure.zip"
	libconfini_cmd_zip := exec.Command("curl", "-L", libconfini_zip_url, "-o", "package.zip")
	err = libconfini_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libconfini_bin_url := "https://github.com/madmurphy/libconfini/releases/download/1.16.4/libconfini-1.16.4-with-configure.bin"
	libconfini_cmd_bin := exec.Command("curl", "-L", libconfini_bin_url, "-o", "binary.bin")
	err = libconfini_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libconfini_src_url := "https://github.com/madmurphy/libconfini/releases/download/1.16.4/libconfini-1.16.4-with-configure.src.tar.gz"
	libconfini_cmd_src := exec.Command("curl", "-L", libconfini_src_url, "-o", "source.tar.gz")
	err = libconfini_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libconfini_cmd_direct := exec.Command("./binary")
	err = libconfini_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
