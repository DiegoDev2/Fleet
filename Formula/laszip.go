package main

// Laszip - Lossless LiDAR compression
// Homepage: https://laszip.org/

import (
	"fmt"
	
	"os/exec"
)

func installLaszip() {
	// Método 1: Descargar y extraer .tar.gz
	laszip_tar_url := "https://github.com/LASzip/LASzip/archive/refs/tags/3.4.4.tar.gz"
	laszip_cmd_tar := exec.Command("curl", "-L", laszip_tar_url, "-o", "package.tar.gz")
	err := laszip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	laszip_zip_url := "https://github.com/LASzip/LASzip/archive/refs/tags/3.4.4.zip"
	laszip_cmd_zip := exec.Command("curl", "-L", laszip_zip_url, "-o", "package.zip")
	err = laszip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	laszip_bin_url := "https://github.com/LASzip/LASzip/archive/refs/tags/3.4.4.bin"
	laszip_cmd_bin := exec.Command("curl", "-L", laszip_bin_url, "-o", "binary.bin")
	err = laszip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	laszip_src_url := "https://github.com/LASzip/LASzip/archive/refs/tags/3.4.4.src.tar.gz"
	laszip_cmd_src := exec.Command("curl", "-L", laszip_src_url, "-o", "source.tar.gz")
	err = laszip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	laszip_cmd_direct := exec.Command("./binary")
	err = laszip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
