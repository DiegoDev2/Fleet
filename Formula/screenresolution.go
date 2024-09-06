package main

// Screenresolution - Get, set, and list display resolution
// Homepage: https://github.com/jhford/screenresolution

import (
	"fmt"
	
	"os/exec"
)

func installScreenresolution() {
	// Método 1: Descargar y extraer .tar.gz
	screenresolution_tar_url := "https://github.com/jhford/screenresolution/archive/refs/tags/v1.6.tar.gz"
	screenresolution_cmd_tar := exec.Command("curl", "-L", screenresolution_tar_url, "-o", "package.tar.gz")
	err := screenresolution_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	screenresolution_zip_url := "https://github.com/jhford/screenresolution/archive/refs/tags/v1.6.zip"
	screenresolution_cmd_zip := exec.Command("curl", "-L", screenresolution_zip_url, "-o", "package.zip")
	err = screenresolution_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	screenresolution_bin_url := "https://github.com/jhford/screenresolution/archive/refs/tags/v1.6.bin"
	screenresolution_cmd_bin := exec.Command("curl", "-L", screenresolution_bin_url, "-o", "binary.bin")
	err = screenresolution_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	screenresolution_src_url := "https://github.com/jhford/screenresolution/archive/refs/tags/v1.6.src.tar.gz"
	screenresolution_cmd_src := exec.Command("curl", "-L", screenresolution_src_url, "-o", "source.tar.gz")
	err = screenresolution_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	screenresolution_cmd_direct := exec.Command("./binary")
	err = screenresolution_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
