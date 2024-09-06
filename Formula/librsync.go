package main

// Librsync - Library that implements the rsync remote-delta algorithm
// Homepage: https://librsync.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installLibrsync() {
	// Método 1: Descargar y extraer .tar.gz
	librsync_tar_url := "https://github.com/librsync/librsync/archive/refs/tags/v2.3.4.tar.gz"
	librsync_cmd_tar := exec.Command("curl", "-L", librsync_tar_url, "-o", "package.tar.gz")
	err := librsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	librsync_zip_url := "https://github.com/librsync/librsync/archive/refs/tags/v2.3.4.zip"
	librsync_cmd_zip := exec.Command("curl", "-L", librsync_zip_url, "-o", "package.zip")
	err = librsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	librsync_bin_url := "https://github.com/librsync/librsync/archive/refs/tags/v2.3.4.bin"
	librsync_cmd_bin := exec.Command("curl", "-L", librsync_bin_url, "-o", "binary.bin")
	err = librsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	librsync_src_url := "https://github.com/librsync/librsync/archive/refs/tags/v2.3.4.src.tar.gz"
	librsync_cmd_src := exec.Command("curl", "-L", librsync_src_url, "-o", "source.tar.gz")
	err = librsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	librsync_cmd_direct := exec.Command("./binary")
	err = librsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: popt")
	exec.Command("latte", "install", "popt").Run()
}
