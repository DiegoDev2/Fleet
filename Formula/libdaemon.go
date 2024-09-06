package main

// Libdaemon - C library that eases writing UNIX daemons
// Homepage: https://0pointer.de/lennart/projects/libdaemon/

import (
	"fmt"
	
	"os/exec"
)

func installLibdaemon() {
	// Método 1: Descargar y extraer .tar.gz
	libdaemon_tar_url := "https://0pointer.de/lennart/projects/libdaemon/libdaemon-0.14.tar.gz"
	libdaemon_cmd_tar := exec.Command("curl", "-L", libdaemon_tar_url, "-o", "package.tar.gz")
	err := libdaemon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdaemon_zip_url := "https://0pointer.de/lennart/projects/libdaemon/libdaemon-0.14.zip"
	libdaemon_cmd_zip := exec.Command("curl", "-L", libdaemon_zip_url, "-o", "package.zip")
	err = libdaemon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdaemon_bin_url := "https://0pointer.de/lennart/projects/libdaemon/libdaemon-0.14.bin"
	libdaemon_cmd_bin := exec.Command("curl", "-L", libdaemon_bin_url, "-o", "binary.bin")
	err = libdaemon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdaemon_src_url := "https://0pointer.de/lennart/projects/libdaemon/libdaemon-0.14.src.tar.gz"
	libdaemon_cmd_src := exec.Command("curl", "-L", libdaemon_src_url, "-o", "source.tar.gz")
	err = libdaemon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdaemon_cmd_direct := exec.Command("./binary")
	err = libdaemon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
