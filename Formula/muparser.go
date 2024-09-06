package main

// Muparser - C++ math expression parser library
// Homepage: https://github.com/beltoforion/muparser

import (
	"fmt"
	
	"os/exec"
)

func installMuparser() {
	// Método 1: Descargar y extraer .tar.gz
	muparser_tar_url := "https://github.com/beltoforion/muparser/archive/refs/tags/v2.3.4.tar.gz"
	muparser_cmd_tar := exec.Command("curl", "-L", muparser_tar_url, "-o", "package.tar.gz")
	err := muparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	muparser_zip_url := "https://github.com/beltoforion/muparser/archive/refs/tags/v2.3.4.zip"
	muparser_cmd_zip := exec.Command("curl", "-L", muparser_zip_url, "-o", "package.zip")
	err = muparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	muparser_bin_url := "https://github.com/beltoforion/muparser/archive/refs/tags/v2.3.4.bin"
	muparser_cmd_bin := exec.Command("curl", "-L", muparser_bin_url, "-o", "binary.bin")
	err = muparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	muparser_src_url := "https://github.com/beltoforion/muparser/archive/refs/tags/v2.3.4.src.tar.gz"
	muparser_cmd_src := exec.Command("curl", "-L", muparser_src_url, "-o", "source.tar.gz")
	err = muparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	muparser_cmd_direct := exec.Command("./binary")
	err = muparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
