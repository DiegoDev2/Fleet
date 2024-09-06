package main

// JpegArchive - Utilities for archiving JPEGs for long term storage
// Homepage: https://github.com/danielgtaylor/jpeg-archive

import (
	"fmt"
	
	"os/exec"
)

func installJpegArchive() {
	// Método 1: Descargar y extraer .tar.gz
	jpegarchive_tar_url := "https://github.com/danielgtaylor/jpeg-archive/archive/refs/tags/v2.2.0.tar.gz"
	jpegarchive_cmd_tar := exec.Command("curl", "-L", jpegarchive_tar_url, "-o", "package.tar.gz")
	err := jpegarchive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jpegarchive_zip_url := "https://github.com/danielgtaylor/jpeg-archive/archive/refs/tags/v2.2.0.zip"
	jpegarchive_cmd_zip := exec.Command("curl", "-L", jpegarchive_zip_url, "-o", "package.zip")
	err = jpegarchive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jpegarchive_bin_url := "https://github.com/danielgtaylor/jpeg-archive/archive/refs/tags/v2.2.0.bin"
	jpegarchive_cmd_bin := exec.Command("curl", "-L", jpegarchive_bin_url, "-o", "binary.bin")
	err = jpegarchive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jpegarchive_src_url := "https://github.com/danielgtaylor/jpeg-archive/archive/refs/tags/v2.2.0.src.tar.gz"
	jpegarchive_cmd_src := exec.Command("curl", "-L", jpegarchive_src_url, "-o", "source.tar.gz")
	err = jpegarchive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jpegarchive_cmd_direct := exec.Command("./binary")
	err = jpegarchive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mozjpeg")
exec.Command("latte", "install", "mozjpeg")
}
