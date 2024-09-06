package main

// Hoverfly - API simulations for development and testing
// Homepage: https://hoverfly.io/

import (
	"fmt"
	
	"os/exec"
)

func installHoverfly() {
	// Método 1: Descargar y extraer .tar.gz
	hoverfly_tar_url := "https://github.com/SpectoLabs/hoverfly/archive/refs/tags/v1.10.4.tar.gz"
	hoverfly_cmd_tar := exec.Command("curl", "-L", hoverfly_tar_url, "-o", "package.tar.gz")
	err := hoverfly_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hoverfly_zip_url := "https://github.com/SpectoLabs/hoverfly/archive/refs/tags/v1.10.4.zip"
	hoverfly_cmd_zip := exec.Command("curl", "-L", hoverfly_zip_url, "-o", "package.zip")
	err = hoverfly_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hoverfly_bin_url := "https://github.com/SpectoLabs/hoverfly/archive/refs/tags/v1.10.4.bin"
	hoverfly_cmd_bin := exec.Command("curl", "-L", hoverfly_bin_url, "-o", "binary.bin")
	err = hoverfly_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hoverfly_src_url := "https://github.com/SpectoLabs/hoverfly/archive/refs/tags/v1.10.4.src.tar.gz"
	hoverfly_cmd_src := exec.Command("curl", "-L", hoverfly_src_url, "-o", "source.tar.gz")
	err = hoverfly_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hoverfly_cmd_direct := exec.Command("./binary")
	err = hoverfly_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
