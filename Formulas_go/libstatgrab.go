package main

// Libstatgrab - Provides cross-platform access to statistics about the system
// Homepage: https://libstatgrab.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibstatgrab() {
	// Método 1: Descargar y extraer .tar.gz
	libstatgrab_tar_url := "https://github.com/libstatgrab/libstatgrab/releases/download/LIBSTATGRAB_0_92_1/libstatgrab-0.92.1.tar.gz"
	libstatgrab_cmd_tar := exec.Command("curl", "-L", libstatgrab_tar_url, "-o", "package.tar.gz")
	err := libstatgrab_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libstatgrab_zip_url := "https://github.com/libstatgrab/libstatgrab/releases/download/LIBSTATGRAB_0_92_1/libstatgrab-0.92.1.zip"
	libstatgrab_cmd_zip := exec.Command("curl", "-L", libstatgrab_zip_url, "-o", "package.zip")
	err = libstatgrab_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libstatgrab_bin_url := "https://github.com/libstatgrab/libstatgrab/releases/download/LIBSTATGRAB_0_92_1/libstatgrab-0.92.1.bin"
	libstatgrab_cmd_bin := exec.Command("curl", "-L", libstatgrab_bin_url, "-o", "binary.bin")
	err = libstatgrab_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libstatgrab_src_url := "https://github.com/libstatgrab/libstatgrab/releases/download/LIBSTATGRAB_0_92_1/libstatgrab-0.92.1.src.tar.gz"
	libstatgrab_cmd_src := exec.Command("curl", "-L", libstatgrab_src_url, "-o", "source.tar.gz")
	err = libstatgrab_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libstatgrab_cmd_direct := exec.Command("./binary")
	err = libstatgrab_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
