package main

// Zplug - Next-generation plugin manager for zsh
// Homepage: https://github.com/zplug/zplug/

import (
	"fmt"
	
	"os/exec"
)

func installZplug() {
	// Método 1: Descargar y extraer .tar.gz
	zplug_tar_url := "https://github.com/zplug/zplug/archive/refs/tags/2.4.2.tar.gz"
	zplug_cmd_tar := exec.Command("curl", "-L", zplug_tar_url, "-o", "package.tar.gz")
	err := zplug_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zplug_zip_url := "https://github.com/zplug/zplug/archive/refs/tags/2.4.2.zip"
	zplug_cmd_zip := exec.Command("curl", "-L", zplug_zip_url, "-o", "package.zip")
	err = zplug_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zplug_bin_url := "https://github.com/zplug/zplug/archive/refs/tags/2.4.2.bin"
	zplug_cmd_bin := exec.Command("curl", "-L", zplug_bin_url, "-o", "binary.bin")
	err = zplug_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zplug_src_url := "https://github.com/zplug/zplug/archive/refs/tags/2.4.2.src.tar.gz"
	zplug_cmd_src := exec.Command("curl", "-L", zplug_src_url, "-o", "source.tar.gz")
	err = zplug_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zplug_cmd_direct := exec.Command("./binary")
	err = zplug_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
