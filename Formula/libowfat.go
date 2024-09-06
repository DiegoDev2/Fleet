package main

// Libowfat - Reimplements libdjb
// Homepage: https://www.fefe.de/libowfat/

import (
	"fmt"
	
	"os/exec"
)

func installLibowfat() {
	// Método 1: Descargar y extraer .tar.gz
	libowfat_tar_url := "https://www.fefe.de/libowfat/libowfat-0.32.tar.xz"
	libowfat_cmd_tar := exec.Command("curl", "-L", libowfat_tar_url, "-o", "package.tar.gz")
	err := libowfat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libowfat_zip_url := "https://www.fefe.de/libowfat/libowfat-0.32.tar.xz"
	libowfat_cmd_zip := exec.Command("curl", "-L", libowfat_zip_url, "-o", "package.zip")
	err = libowfat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libowfat_bin_url := "https://www.fefe.de/libowfat/libowfat-0.32.tar.xz"
	libowfat_cmd_bin := exec.Command("curl", "-L", libowfat_bin_url, "-o", "binary.bin")
	err = libowfat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libowfat_src_url := "https://www.fefe.de/libowfat/libowfat-0.32.tar.xz"
	libowfat_cmd_src := exec.Command("curl", "-L", libowfat_src_url, "-o", "source.tar.gz")
	err = libowfat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libowfat_cmd_direct := exec.Command("./binary")
	err = libowfat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
