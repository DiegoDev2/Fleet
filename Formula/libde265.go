package main

// Libde265 - Open h.265 video codec implementation
// Homepage: https://github.com/strukturag/libde265

import (
	"fmt"
	
	"os/exec"
)

func installLibde265() {
	// Método 1: Descargar y extraer .tar.gz
	libde265_tar_url := "https://github.com/strukturag/libde265/releases/download/v1.0.15/libde265-1.0.15.tar.gz"
	libde265_cmd_tar := exec.Command("curl", "-L", libde265_tar_url, "-o", "package.tar.gz")
	err := libde265_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libde265_zip_url := "https://github.com/strukturag/libde265/releases/download/v1.0.15/libde265-1.0.15.zip"
	libde265_cmd_zip := exec.Command("curl", "-L", libde265_zip_url, "-o", "package.zip")
	err = libde265_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libde265_bin_url := "https://github.com/strukturag/libde265/releases/download/v1.0.15/libde265-1.0.15.bin"
	libde265_cmd_bin := exec.Command("curl", "-L", libde265_bin_url, "-o", "binary.bin")
	err = libde265_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libde265_src_url := "https://github.com/strukturag/libde265/releases/download/v1.0.15/libde265-1.0.15.src.tar.gz"
	libde265_cmd_src := exec.Command("curl", "-L", libde265_src_url, "-o", "source.tar.gz")
	err = libde265_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libde265_cmd_direct := exec.Command("./binary")
	err = libde265_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
