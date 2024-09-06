package main

// VisionmediaWatch - Periodically executes the given command
// Homepage: https://github.com/tj/watch

import (
	"fmt"
	
	"os/exec"
)

func installVisionmediaWatch() {
	// Método 1: Descargar y extraer .tar.gz
	visionmediawatch_tar_url := "https://github.com/tj/watch/archive/refs/tags/0.4.0.tar.gz"
	visionmediawatch_cmd_tar := exec.Command("curl", "-L", visionmediawatch_tar_url, "-o", "package.tar.gz")
	err := visionmediawatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	visionmediawatch_zip_url := "https://github.com/tj/watch/archive/refs/tags/0.4.0.zip"
	visionmediawatch_cmd_zip := exec.Command("curl", "-L", visionmediawatch_zip_url, "-o", "package.zip")
	err = visionmediawatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	visionmediawatch_bin_url := "https://github.com/tj/watch/archive/refs/tags/0.4.0.bin"
	visionmediawatch_cmd_bin := exec.Command("curl", "-L", visionmediawatch_bin_url, "-o", "binary.bin")
	err = visionmediawatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	visionmediawatch_src_url := "https://github.com/tj/watch/archive/refs/tags/0.4.0.src.tar.gz"
	visionmediawatch_cmd_src := exec.Command("curl", "-L", visionmediawatch_src_url, "-o", "source.tar.gz")
	err = visionmediawatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	visionmediawatch_cmd_direct := exec.Command("./binary")
	err = visionmediawatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
