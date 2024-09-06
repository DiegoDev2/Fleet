package main

// F3 - Test various flash cards
// Homepage: https://fight-flash-fraud.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installF3() {
	// Método 1: Descargar y extraer .tar.gz
	f3_tar_url := "https://github.com/AltraMayor/f3/archive/refs/tags/v8.0.tar.gz"
	f3_cmd_tar := exec.Command("curl", "-L", f3_tar_url, "-o", "package.tar.gz")
	err := f3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	f3_zip_url := "https://github.com/AltraMayor/f3/archive/refs/tags/v8.0.zip"
	f3_cmd_zip := exec.Command("curl", "-L", f3_zip_url, "-o", "package.zip")
	err = f3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	f3_bin_url := "https://github.com/AltraMayor/f3/archive/refs/tags/v8.0.bin"
	f3_cmd_bin := exec.Command("curl", "-L", f3_bin_url, "-o", "binary.bin")
	err = f3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	f3_src_url := "https://github.com/AltraMayor/f3/archive/refs/tags/v8.0.src.tar.gz"
	f3_cmd_src := exec.Command("curl", "-L", f3_src_url, "-o", "source.tar.gz")
	err = f3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	f3_cmd_direct := exec.Command("./binary")
	err = f3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: argp-standalone")
	exec.Command("latte", "install", "argp-standalone").Run()
}
