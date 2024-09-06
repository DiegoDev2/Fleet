package main

// Gpx - Gcode to x3g converter for 3D printers running Sailfish
// Homepage: https://github.com/markwal/GPX/blob/HEAD/README.md

import (
	"fmt"
	
	"os/exec"
)

func installGpx() {
	// Método 1: Descargar y extraer .tar.gz
	gpx_tar_url := "https://github.com/markwal/GPX/archive/refs/tags/2.6.8.tar.gz"
	gpx_cmd_tar := exec.Command("curl", "-L", gpx_tar_url, "-o", "package.tar.gz")
	err := gpx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gpx_zip_url := "https://github.com/markwal/GPX/archive/refs/tags/2.6.8.zip"
	gpx_cmd_zip := exec.Command("curl", "-L", gpx_zip_url, "-o", "package.zip")
	err = gpx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gpx_bin_url := "https://github.com/markwal/GPX/archive/refs/tags/2.6.8.bin"
	gpx_cmd_bin := exec.Command("curl", "-L", gpx_bin_url, "-o", "binary.bin")
	err = gpx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gpx_src_url := "https://github.com/markwal/GPX/archive/refs/tags/2.6.8.src.tar.gz"
	gpx_cmd_src := exec.Command("curl", "-L", gpx_src_url, "-o", "source.tar.gz")
	err = gpx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gpx_cmd_direct := exec.Command("./binary")
	err = gpx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
