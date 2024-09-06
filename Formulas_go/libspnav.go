package main

// Libspnav - Client library for connecting to 3Dconnexion's 3D input devices
// Homepage: https://spacenav.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibspnav() {
	// Método 1: Descargar y extraer .tar.gz
	libspnav_tar_url := "https://github.com/FreeSpacenav/libspnav/releases/download/v1.1/libspnav-1.1.tar.gz"
	libspnav_cmd_tar := exec.Command("curl", "-L", libspnav_tar_url, "-o", "package.tar.gz")
	err := libspnav_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libspnav_zip_url := "https://github.com/FreeSpacenav/libspnav/releases/download/v1.1/libspnav-1.1.zip"
	libspnav_cmd_zip := exec.Command("curl", "-L", libspnav_zip_url, "-o", "package.zip")
	err = libspnav_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libspnav_bin_url := "https://github.com/FreeSpacenav/libspnav/releases/download/v1.1/libspnav-1.1.bin"
	libspnav_cmd_bin := exec.Command("curl", "-L", libspnav_bin_url, "-o", "binary.bin")
	err = libspnav_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libspnav_src_url := "https://github.com/FreeSpacenav/libspnav/releases/download/v1.1/libspnav-1.1.src.tar.gz"
	libspnav_cmd_src := exec.Command("curl", "-L", libspnav_src_url, "-o", "source.tar.gz")
	err = libspnav_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libspnav_cmd_direct := exec.Command("./binary")
	err = libspnav_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
