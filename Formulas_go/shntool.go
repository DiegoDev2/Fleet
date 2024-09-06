package main

// Shntool - Multi-purpose tool for manipulating and analyzing WAV files
// Homepage: http://shnutils.freeshell.org/shntool/

import (
	"fmt"
	
	"os/exec"
)

func installShntool() {
	// Método 1: Descargar y extraer .tar.gz
	shntool_tar_url := "http://shnutils.freeshell.org/shntool/dist/src/shntool-3.0.10.tar.gz"
	shntool_cmd_tar := exec.Command("curl", "-L", shntool_tar_url, "-o", "package.tar.gz")
	err := shntool_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shntool_zip_url := "http://shnutils.freeshell.org/shntool/dist/src/shntool-3.0.10.zip"
	shntool_cmd_zip := exec.Command("curl", "-L", shntool_zip_url, "-o", "package.zip")
	err = shntool_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shntool_bin_url := "http://shnutils.freeshell.org/shntool/dist/src/shntool-3.0.10.bin"
	shntool_cmd_bin := exec.Command("curl", "-L", shntool_bin_url, "-o", "binary.bin")
	err = shntool_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shntool_src_url := "http://shnutils.freeshell.org/shntool/dist/src/shntool-3.0.10.src.tar.gz"
	shntool_cmd_src := exec.Command("curl", "-L", shntool_src_url, "-o", "source.tar.gz")
	err = shntool_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shntool_cmd_direct := exec.Command("./binary")
	err = shntool_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
