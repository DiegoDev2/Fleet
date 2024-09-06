package main

// Libbinio - Binary I/O stream class library
// Homepage: https://adplug.github.io/libbinio/

import (
	"fmt"
	
	"os/exec"
)

func installLibbinio() {
	// Método 1: Descargar y extraer .tar.gz
	libbinio_tar_url := "https://github.com/adplug/libbinio/releases/download/libbinio-1.5/libbinio-1.5.tar.bz2"
	libbinio_cmd_tar := exec.Command("curl", "-L", libbinio_tar_url, "-o", "package.tar.gz")
	err := libbinio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libbinio_zip_url := "https://github.com/adplug/libbinio/releases/download/libbinio-1.5/libbinio-1.5.tar.bz2"
	libbinio_cmd_zip := exec.Command("curl", "-L", libbinio_zip_url, "-o", "package.zip")
	err = libbinio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libbinio_bin_url := "https://github.com/adplug/libbinio/releases/download/libbinio-1.5/libbinio-1.5.tar.bz2"
	libbinio_cmd_bin := exec.Command("curl", "-L", libbinio_bin_url, "-o", "binary.bin")
	err = libbinio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libbinio_src_url := "https://github.com/adplug/libbinio/releases/download/libbinio-1.5/libbinio-1.5.tar.bz2"
	libbinio_cmd_src := exec.Command("curl", "-L", libbinio_src_url, "-o", "source.tar.gz")
	err = libbinio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libbinio_cmd_direct := exec.Command("./binary")
	err = libbinio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
