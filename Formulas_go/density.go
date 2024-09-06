package main

// Density - Superfast compression library
// Homepage: https://github.com/g1mv/density

import (
	"fmt"
	
	"os/exec"
)

func installDensity() {
	// Método 1: Descargar y extraer .tar.gz
	density_tar_url := "https://github.com/g1mv/density/archive/refs/tags/density-0.14.2.tar.gz"
	density_cmd_tar := exec.Command("curl", "-L", density_tar_url, "-o", "package.tar.gz")
	err := density_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	density_zip_url := "https://github.com/g1mv/density/archive/refs/tags/density-0.14.2.zip"
	density_cmd_zip := exec.Command("curl", "-L", density_zip_url, "-o", "package.zip")
	err = density_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	density_bin_url := "https://github.com/g1mv/density/archive/refs/tags/density-0.14.2.bin"
	density_cmd_bin := exec.Command("curl", "-L", density_bin_url, "-o", "binary.bin")
	err = density_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	density_src_url := "https://github.com/g1mv/density/archive/refs/tags/density-0.14.2.src.tar.gz"
	density_cmd_src := exec.Command("curl", "-L", density_src_url, "-o", "source.tar.gz")
	err = density_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	density_cmd_direct := exec.Command("./binary")
	err = density_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
