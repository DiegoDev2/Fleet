package main

// AngleGrinder - Slice and dice log files on the command-line
// Homepage: https://github.com/rcoh/angle-grinder

import (
	"fmt"
	
	"os/exec"
)

func installAngleGrinder() {
	// Método 1: Descargar y extraer .tar.gz
	anglegrinder_tar_url := "https://github.com/rcoh/angle-grinder/archive/refs/tags/v0.19.4.tar.gz"
	anglegrinder_cmd_tar := exec.Command("curl", "-L", anglegrinder_tar_url, "-o", "package.tar.gz")
	err := anglegrinder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	anglegrinder_zip_url := "https://github.com/rcoh/angle-grinder/archive/refs/tags/v0.19.4.zip"
	anglegrinder_cmd_zip := exec.Command("curl", "-L", anglegrinder_zip_url, "-o", "package.zip")
	err = anglegrinder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	anglegrinder_bin_url := "https://github.com/rcoh/angle-grinder/archive/refs/tags/v0.19.4.bin"
	anglegrinder_cmd_bin := exec.Command("curl", "-L", anglegrinder_bin_url, "-o", "binary.bin")
	err = anglegrinder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	anglegrinder_src_url := "https://github.com/rcoh/angle-grinder/archive/refs/tags/v0.19.4.src.tar.gz"
	anglegrinder_cmd_src := exec.Command("curl", "-L", anglegrinder_src_url, "-o", "source.tar.gz")
	err = anglegrinder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	anglegrinder_cmd_direct := exec.Command("./binary")
	err = anglegrinder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
