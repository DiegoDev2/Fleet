package main

// Freeimage - Library for FreeImage, a dependency-free graphics library
// Homepage: https://sourceforge.net/projects/freeimage/

import (
	"fmt"
	
	"os/exec"
)

func installFreeimage() {
	// Método 1: Descargar y extraer .tar.gz
	freeimage_tar_url := "https://downloads.sourceforge.net/project/freeimage/Source%20Distribution/3.18.0/FreeImage3180.zip"
	freeimage_cmd_tar := exec.Command("curl", "-L", freeimage_tar_url, "-o", "package.tar.gz")
	err := freeimage_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	freeimage_zip_url := "https://downloads.sourceforge.net/project/freeimage/Source%20Distribution/3.18.0/FreeImage3180.zip"
	freeimage_cmd_zip := exec.Command("curl", "-L", freeimage_zip_url, "-o", "package.zip")
	err = freeimage_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	freeimage_bin_url := "https://downloads.sourceforge.net/project/freeimage/Source%20Distribution/3.18.0/FreeImage3180.zip"
	freeimage_cmd_bin := exec.Command("curl", "-L", freeimage_bin_url, "-o", "binary.bin")
	err = freeimage_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	freeimage_src_url := "https://downloads.sourceforge.net/project/freeimage/Source%20Distribution/3.18.0/FreeImage3180.zip"
	freeimage_cmd_src := exec.Command("curl", "-L", freeimage_src_url, "-o", "source.tar.gz")
	err = freeimage_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	freeimage_cmd_direct := exec.Command("./binary")
	err = freeimage_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
