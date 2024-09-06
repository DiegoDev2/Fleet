package main

// Bcal - Storage conversion and expression calculator
// Homepage: https://github.com/jarun/bcal

import (
	"fmt"
	
	"os/exec"
)

func installBcal() {
	// Método 1: Descargar y extraer .tar.gz
	bcal_tar_url := "https://github.com/jarun/bcal/archive/refs/tags/v2.4.tar.gz"
	bcal_cmd_tar := exec.Command("curl", "-L", bcal_tar_url, "-o", "package.tar.gz")
	err := bcal_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bcal_zip_url := "https://github.com/jarun/bcal/archive/refs/tags/v2.4.zip"
	bcal_cmd_zip := exec.Command("curl", "-L", bcal_zip_url, "-o", "package.zip")
	err = bcal_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bcal_bin_url := "https://github.com/jarun/bcal/archive/refs/tags/v2.4.bin"
	bcal_cmd_bin := exec.Command("curl", "-L", bcal_bin_url, "-o", "binary.bin")
	err = bcal_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bcal_src_url := "https://github.com/jarun/bcal/archive/refs/tags/v2.4.src.tar.gz"
	bcal_cmd_src := exec.Command("curl", "-L", bcal_src_url, "-o", "source.tar.gz")
	err = bcal_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bcal_cmd_direct := exec.Command("./binary")
	err = bcal_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
}
