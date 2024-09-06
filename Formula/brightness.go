package main

// Brightness - Change macOS display brightness from the command-line
// Homepage: https://github.com/nriley/brightness

import (
	"fmt"
	
	"os/exec"
)

func installBrightness() {
	// Método 1: Descargar y extraer .tar.gz
	brightness_tar_url := "https://github.com/nriley/brightness/archive/refs/tags/1.2.tar.gz"
	brightness_cmd_tar := exec.Command("curl", "-L", brightness_tar_url, "-o", "package.tar.gz")
	err := brightness_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brightness_zip_url := "https://github.com/nriley/brightness/archive/refs/tags/1.2.zip"
	brightness_cmd_zip := exec.Command("curl", "-L", brightness_zip_url, "-o", "package.zip")
	err = brightness_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brightness_bin_url := "https://github.com/nriley/brightness/archive/refs/tags/1.2.bin"
	brightness_cmd_bin := exec.Command("curl", "-L", brightness_bin_url, "-o", "binary.bin")
	err = brightness_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brightness_src_url := "https://github.com/nriley/brightness/archive/refs/tags/1.2.src.tar.gz"
	brightness_cmd_src := exec.Command("curl", "-L", brightness_src_url, "-o", "source.tar.gz")
	err = brightness_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brightness_cmd_direct := exec.Command("./binary")
	err = brightness_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
