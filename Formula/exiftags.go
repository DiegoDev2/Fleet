package main

// Exiftags - Utility to read EXIF tags from a digital camera JPEG file
// Homepage: https://johnst.org/sw/exiftags/

import (
	"fmt"
	
	"os/exec"
)

func installExiftags() {
	// Método 1: Descargar y extraer .tar.gz
	exiftags_tar_url := "https://johnst.org/sw/exiftags/exiftags-1.01.tar.gz"
	exiftags_cmd_tar := exec.Command("curl", "-L", exiftags_tar_url, "-o", "package.tar.gz")
	err := exiftags_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exiftags_zip_url := "https://johnst.org/sw/exiftags/exiftags-1.01.zip"
	exiftags_cmd_zip := exec.Command("curl", "-L", exiftags_zip_url, "-o", "package.zip")
	err = exiftags_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exiftags_bin_url := "https://johnst.org/sw/exiftags/exiftags-1.01.bin"
	exiftags_cmd_bin := exec.Command("curl", "-L", exiftags_bin_url, "-o", "binary.bin")
	err = exiftags_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exiftags_src_url := "https://johnst.org/sw/exiftags/exiftags-1.01.src.tar.gz"
	exiftags_cmd_src := exec.Command("curl", "-L", exiftags_src_url, "-o", "source.tar.gz")
	err = exiftags_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exiftags_cmd_direct := exec.Command("./binary")
	err = exiftags_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
