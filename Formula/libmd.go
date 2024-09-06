package main

// Libmd - BSD Message Digest library
// Homepage: https://www.hadrons.org/software/libmd/

import (
	"fmt"
	
	"os/exec"
)

func installLibmd() {
	// Método 1: Descargar y extraer .tar.gz
	libmd_tar_url := "https://libbsd.freedesktop.org/releases/libmd-1.1.0.tar.xz"
	libmd_cmd_tar := exec.Command("curl", "-L", libmd_tar_url, "-o", "package.tar.gz")
	err := libmd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmd_zip_url := "https://libbsd.freedesktop.org/releases/libmd-1.1.0.tar.xz"
	libmd_cmd_zip := exec.Command("curl", "-L", libmd_zip_url, "-o", "package.zip")
	err = libmd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmd_bin_url := "https://libbsd.freedesktop.org/releases/libmd-1.1.0.tar.xz"
	libmd_cmd_bin := exec.Command("curl", "-L", libmd_bin_url, "-o", "binary.bin")
	err = libmd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmd_src_url := "https://libbsd.freedesktop.org/releases/libmd-1.1.0.tar.xz"
	libmd_cmd_src := exec.Command("curl", "-L", libmd_src_url, "-o", "source.tar.gz")
	err = libmd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmd_cmd_direct := exec.Command("./binary")
	err = libmd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
