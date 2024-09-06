package main

// Libmps - Memory Pool System
// Homepage: https://www.ravenbrook.com/project/mps/

import (
	"fmt"
	
	"os/exec"
)

func installLibmps() {
	// Método 1: Descargar y extraer .tar.gz
	libmps_tar_url := "https://github.com/Ravenbrook/mps/archive/refs/tags/release-1.118.0.tar.gz"
	libmps_cmd_tar := exec.Command("curl", "-L", libmps_tar_url, "-o", "package.tar.gz")
	err := libmps_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmps_zip_url := "https://github.com/Ravenbrook/mps/archive/refs/tags/release-1.118.0.zip"
	libmps_cmd_zip := exec.Command("curl", "-L", libmps_zip_url, "-o", "package.zip")
	err = libmps_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmps_bin_url := "https://github.com/Ravenbrook/mps/archive/refs/tags/release-1.118.0.bin"
	libmps_cmd_bin := exec.Command("curl", "-L", libmps_bin_url, "-o", "binary.bin")
	err = libmps_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmps_src_url := "https://github.com/Ravenbrook/mps/archive/refs/tags/release-1.118.0.src.tar.gz"
	libmps_cmd_src := exec.Command("curl", "-L", libmps_src_url, "-o", "source.tar.gz")
	err = libmps_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmps_cmd_direct := exec.Command("./binary")
	err = libmps_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
