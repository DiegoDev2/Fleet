package main

// Hwatch - Modern alternative to the watch command
// Homepage: https://github.com/blacknon/hwatch

import (
	"fmt"
	
	"os/exec"
)

func installHwatch() {
	// Método 1: Descargar y extraer .tar.gz
	hwatch_tar_url := "https://github.com/blacknon/hwatch/archive/refs/tags/0.3.15.tar.gz"
	hwatch_cmd_tar := exec.Command("curl", "-L", hwatch_tar_url, "-o", "package.tar.gz")
	err := hwatch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hwatch_zip_url := "https://github.com/blacknon/hwatch/archive/refs/tags/0.3.15.zip"
	hwatch_cmd_zip := exec.Command("curl", "-L", hwatch_zip_url, "-o", "package.zip")
	err = hwatch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hwatch_bin_url := "https://github.com/blacknon/hwatch/archive/refs/tags/0.3.15.bin"
	hwatch_cmd_bin := exec.Command("curl", "-L", hwatch_bin_url, "-o", "binary.bin")
	err = hwatch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hwatch_src_url := "https://github.com/blacknon/hwatch/archive/refs/tags/0.3.15.src.tar.gz"
	hwatch_cmd_src := exec.Command("curl", "-L", hwatch_src_url, "-o", "source.tar.gz")
	err = hwatch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hwatch_cmd_direct := exec.Command("./binary")
	err = hwatch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
