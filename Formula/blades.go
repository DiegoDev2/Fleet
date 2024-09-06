package main

// Blades - Blazing fast dead simple static site generator
// Homepage: https://getblades.org/

import (
	"fmt"
	
	"os/exec"
)

func installBlades() {
	// Método 1: Descargar y extraer .tar.gz
	blades_tar_url := "https://github.com/grego/blades/archive/refs/tags/v0.5.0.tar.gz"
	blades_cmd_tar := exec.Command("curl", "-L", blades_tar_url, "-o", "package.tar.gz")
	err := blades_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	blades_zip_url := "https://github.com/grego/blades/archive/refs/tags/v0.5.0.zip"
	blades_cmd_zip := exec.Command("curl", "-L", blades_zip_url, "-o", "package.zip")
	err = blades_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	blades_bin_url := "https://github.com/grego/blades/archive/refs/tags/v0.5.0.bin"
	blades_cmd_bin := exec.Command("curl", "-L", blades_bin_url, "-o", "binary.bin")
	err = blades_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	blades_src_url := "https://github.com/grego/blades/archive/refs/tags/v0.5.0.src.tar.gz"
	blades_cmd_src := exec.Command("curl", "-L", blades_src_url, "-o", "source.tar.gz")
	err = blades_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	blades_cmd_direct := exec.Command("./binary")
	err = blades_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
