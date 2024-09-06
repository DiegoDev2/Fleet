package main

// Termcolor - Header-only C++ library for printing colored messages
// Homepage: https://termcolor.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installTermcolor() {
	// Método 1: Descargar y extraer .tar.gz
	termcolor_tar_url := "https://github.com/ikalnytskyi/termcolor/archive/refs/tags/v2.1.0.tar.gz"
	termcolor_cmd_tar := exec.Command("curl", "-L", termcolor_tar_url, "-o", "package.tar.gz")
	err := termcolor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	termcolor_zip_url := "https://github.com/ikalnytskyi/termcolor/archive/refs/tags/v2.1.0.zip"
	termcolor_cmd_zip := exec.Command("curl", "-L", termcolor_zip_url, "-o", "package.zip")
	err = termcolor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	termcolor_bin_url := "https://github.com/ikalnytskyi/termcolor/archive/refs/tags/v2.1.0.bin"
	termcolor_cmd_bin := exec.Command("curl", "-L", termcolor_bin_url, "-o", "binary.bin")
	err = termcolor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	termcolor_src_url := "https://github.com/ikalnytskyi/termcolor/archive/refs/tags/v2.1.0.src.tar.gz"
	termcolor_cmd_src := exec.Command("curl", "-L", termcolor_src_url, "-o", "source.tar.gz")
	err = termcolor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	termcolor_cmd_direct := exec.Command("./binary")
	err = termcolor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
