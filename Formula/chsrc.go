package main

// Chsrc - Change Source for every software on every platform from the command-line
// Homepage: https://github.com/RubyMetric/chsrc

import (
	"fmt"
	
	"os/exec"
)

func installChsrc() {
	// Método 1: Descargar y extraer .tar.gz
	chsrc_tar_url := "https://github.com/RubyMetric/chsrc/archive/refs/tags/v0.1.8.1.tar.gz"
	chsrc_cmd_tar := exec.Command("curl", "-L", chsrc_tar_url, "-o", "package.tar.gz")
	err := chsrc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chsrc_zip_url := "https://github.com/RubyMetric/chsrc/archive/refs/tags/v0.1.8.1.zip"
	chsrc_cmd_zip := exec.Command("curl", "-L", chsrc_zip_url, "-o", "package.zip")
	err = chsrc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chsrc_bin_url := "https://github.com/RubyMetric/chsrc/archive/refs/tags/v0.1.8.1.bin"
	chsrc_cmd_bin := exec.Command("curl", "-L", chsrc_bin_url, "-o", "binary.bin")
	err = chsrc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chsrc_src_url := "https://github.com/RubyMetric/chsrc/archive/refs/tags/v0.1.8.1.src.tar.gz"
	chsrc_cmd_src := exec.Command("curl", "-L", chsrc_src_url, "-o", "source.tar.gz")
	err = chsrc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chsrc_cmd_direct := exec.Command("./binary")
	err = chsrc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
