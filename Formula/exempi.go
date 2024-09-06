package main

// Exempi - Library to parse XMP metadata
// Homepage: https://wiki.freedesktop.org/libopenraw/Exempi/

import (
	"fmt"
	
	"os/exec"
)

func installExempi() {
	// Método 1: Descargar y extraer .tar.gz
	exempi_tar_url := "https://libopenraw.freedesktop.org/download/exempi-2.6.5.tar.bz2"
	exempi_cmd_tar := exec.Command("curl", "-L", exempi_tar_url, "-o", "package.tar.gz")
	err := exempi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exempi_zip_url := "https://libopenraw.freedesktop.org/download/exempi-2.6.5.tar.bz2"
	exempi_cmd_zip := exec.Command("curl", "-L", exempi_zip_url, "-o", "package.zip")
	err = exempi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exempi_bin_url := "https://libopenraw.freedesktop.org/download/exempi-2.6.5.tar.bz2"
	exempi_cmd_bin := exec.Command("curl", "-L", exempi_bin_url, "-o", "binary.bin")
	err = exempi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exempi_src_url := "https://libopenraw.freedesktop.org/download/exempi-2.6.5.tar.bz2"
	exempi_cmd_src := exec.Command("curl", "-L", exempi_src_url, "-o", "source.tar.gz")
	err = exempi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exempi_cmd_direct := exec.Command("./binary")
	err = exempi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
