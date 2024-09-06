package main

// Jpegoptim - Utility to optimize JPEG files
// Homepage: https://github.com/tjko/jpegoptim

import (
	"fmt"
	
	"os/exec"
)

func installJpegoptim() {
	// Método 1: Descargar y extraer .tar.gz
	jpegoptim_tar_url := "https://github.com/tjko/jpegoptim/archive/refs/tags/v1.5.5.tar.gz"
	jpegoptim_cmd_tar := exec.Command("curl", "-L", jpegoptim_tar_url, "-o", "package.tar.gz")
	err := jpegoptim_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jpegoptim_zip_url := "https://github.com/tjko/jpegoptim/archive/refs/tags/v1.5.5.zip"
	jpegoptim_cmd_zip := exec.Command("curl", "-L", jpegoptim_zip_url, "-o", "package.zip")
	err = jpegoptim_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jpegoptim_bin_url := "https://github.com/tjko/jpegoptim/archive/refs/tags/v1.5.5.bin"
	jpegoptim_cmd_bin := exec.Command("curl", "-L", jpegoptim_bin_url, "-o", "binary.bin")
	err = jpegoptim_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jpegoptim_src_url := "https://github.com/tjko/jpegoptim/archive/refs/tags/v1.5.5.src.tar.gz"
	jpegoptim_cmd_src := exec.Command("curl", "-L", jpegoptim_src_url, "-o", "source.tar.gz")
	err = jpegoptim_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jpegoptim_cmd_direct := exec.Command("./binary")
	err = jpegoptim_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
}
