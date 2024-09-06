package main

// OilsForUnix - Bash-compatible Unix shell with more consistent syntax and semantics
// Homepage: https://www.oilshell.org/

import (
	"fmt"
	
	"os/exec"
)

func installOilsForUnix() {
	// Método 1: Descargar y extraer .tar.gz
	oilsforunix_tar_url := "https://www.oilshell.org/download/oils-for-unix-0.23.0.tar.gz"
	oilsforunix_cmd_tar := exec.Command("curl", "-L", oilsforunix_tar_url, "-o", "package.tar.gz")
	err := oilsforunix_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oilsforunix_zip_url := "https://www.oilshell.org/download/oils-for-unix-0.23.0.zip"
	oilsforunix_cmd_zip := exec.Command("curl", "-L", oilsforunix_zip_url, "-o", "package.zip")
	err = oilsforunix_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oilsforunix_bin_url := "https://www.oilshell.org/download/oils-for-unix-0.23.0.bin"
	oilsforunix_cmd_bin := exec.Command("curl", "-L", oilsforunix_bin_url, "-o", "binary.bin")
	err = oilsforunix_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oilsforunix_src_url := "https://www.oilshell.org/download/oils-for-unix-0.23.0.src.tar.gz"
	oilsforunix_cmd_src := exec.Command("curl", "-L", oilsforunix_src_url, "-o", "source.tar.gz")
	err = oilsforunix_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oilsforunix_cmd_direct := exec.Command("./binary")
	err = oilsforunix_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
