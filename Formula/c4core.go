package main

// C4core - C++ utilities
// Homepage: https://github.com/biojppm/c4core

import (
	"fmt"
	
	"os/exec"
)

func installC4core() {
	// Método 1: Descargar y extraer .tar.gz
	c4core_tar_url := "https://github.com/biojppm/c4core/releases/download/v0.2.2/c4core-0.2.2-src.tgz"
	c4core_cmd_tar := exec.Command("curl", "-L", c4core_tar_url, "-o", "package.tar.gz")
	err := c4core_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	c4core_zip_url := "https://github.com/biojppm/c4core/releases/download/v0.2.2/c4core-0.2.2-src.tgz"
	c4core_cmd_zip := exec.Command("curl", "-L", c4core_zip_url, "-o", "package.zip")
	err = c4core_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	c4core_bin_url := "https://github.com/biojppm/c4core/releases/download/v0.2.2/c4core-0.2.2-src.tgz"
	c4core_cmd_bin := exec.Command("curl", "-L", c4core_bin_url, "-o", "binary.bin")
	err = c4core_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	c4core_src_url := "https://github.com/biojppm/c4core/releases/download/v0.2.2/c4core-0.2.2-src.tgz"
	c4core_cmd_src := exec.Command("curl", "-L", c4core_src_url, "-o", "source.tar.gz")
	err = c4core_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	c4core_cmd_direct := exec.Command("./binary")
	err = c4core_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
