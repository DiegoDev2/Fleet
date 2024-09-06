package main

// Haiti - Hash type identifier
// Homepage: https://noraj.github.io/haiti/#/

import (
	"fmt"
	
	"os/exec"
)

func installHaiti() {
	// Método 1: Descargar y extraer .tar.gz
	haiti_tar_url := "https://github.com/noraj/haiti/archive/refs/tags/v2.1.0.tar.gz"
	haiti_cmd_tar := exec.Command("curl", "-L", haiti_tar_url, "-o", "package.tar.gz")
	err := haiti_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	haiti_zip_url := "https://github.com/noraj/haiti/archive/refs/tags/v2.1.0.zip"
	haiti_cmd_zip := exec.Command("curl", "-L", haiti_zip_url, "-o", "package.zip")
	err = haiti_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	haiti_bin_url := "https://github.com/noraj/haiti/archive/refs/tags/v2.1.0.bin"
	haiti_cmd_bin := exec.Command("curl", "-L", haiti_bin_url, "-o", "binary.bin")
	err = haiti_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	haiti_src_url := "https://github.com/noraj/haiti/archive/refs/tags/v2.1.0.src.tar.gz"
	haiti_cmd_src := exec.Command("curl", "-L", haiti_src_url, "-o", "source.tar.gz")
	err = haiti_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	haiti_cmd_direct := exec.Command("./binary")
	err = haiti_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
