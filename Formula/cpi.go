package main

// Cpi - Tiny c++ interpreter
// Homepage: https://treefrogframework.github.io/cpi/

import (
	"fmt"
	
	"os/exec"
)

func installCpi() {
	// Método 1: Descargar y extraer .tar.gz
	cpi_tar_url := "https://github.com/treefrogframework/cpi/archive/refs/tags/v2.1.0.tar.gz"
	cpi_cmd_tar := exec.Command("curl", "-L", cpi_tar_url, "-o", "package.tar.gz")
	err := cpi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpi_zip_url := "https://github.com/treefrogframework/cpi/archive/refs/tags/v2.1.0.zip"
	cpi_cmd_zip := exec.Command("curl", "-L", cpi_zip_url, "-o", "package.zip")
	err = cpi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpi_bin_url := "https://github.com/treefrogframework/cpi/archive/refs/tags/v2.1.0.bin"
	cpi_cmd_bin := exec.Command("curl", "-L", cpi_bin_url, "-o", "binary.bin")
	err = cpi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpi_src_url := "https://github.com/treefrogframework/cpi/archive/refs/tags/v2.1.0.src.tar.gz"
	cpi_cmd_src := exec.Command("curl", "-L", cpi_src_url, "-o", "source.tar.gz")
	err = cpi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpi_cmd_direct := exec.Command("./binary")
	err = cpi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: qt")
	exec.Command("latte", "install", "qt").Run()
}
