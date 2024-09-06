package main

// Libdc1394 - Provides API for IEEE 1394 cameras
// Homepage: https://damien.douxchamps.net/ieee1394/libdc1394/

import (
	"fmt"
	
	"os/exec"
)

func installLibdc1394() {
	// Método 1: Descargar y extraer .tar.gz
	libdc1394_tar_url := "https://downloads.sourceforge.net/project/libdc1394/libdc1394-2/2.2.7/libdc1394-2.2.7.tar.gz"
	libdc1394_cmd_tar := exec.Command("curl", "-L", libdc1394_tar_url, "-o", "package.tar.gz")
	err := libdc1394_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdc1394_zip_url := "https://downloads.sourceforge.net/project/libdc1394/libdc1394-2/2.2.7/libdc1394-2.2.7.zip"
	libdc1394_cmd_zip := exec.Command("curl", "-L", libdc1394_zip_url, "-o", "package.zip")
	err = libdc1394_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdc1394_bin_url := "https://downloads.sourceforge.net/project/libdc1394/libdc1394-2/2.2.7/libdc1394-2.2.7.bin"
	libdc1394_cmd_bin := exec.Command("curl", "-L", libdc1394_bin_url, "-o", "binary.bin")
	err = libdc1394_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdc1394_src_url := "https://downloads.sourceforge.net/project/libdc1394/libdc1394-2/2.2.7/libdc1394-2.2.7.src.tar.gz"
	libdc1394_cmd_src := exec.Command("curl", "-L", libdc1394_src_url, "-o", "source.tar.gz")
	err = libdc1394_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdc1394_cmd_direct := exec.Command("./binary")
	err = libdc1394_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sdl12-compat")
exec.Command("latte", "install", "sdl12-compat")
}
