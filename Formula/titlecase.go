package main

// Titlecase - Script to convert text to title case
// Homepage: http://plasmasturm.org/code/titlecase/

import (
	"fmt"
	
	"os/exec"
)

func installTitlecase() {
	// Método 1: Descargar y extraer .tar.gz
	titlecase_tar_url := "https://github.com/ap/titlecase/archive/refs/tags/v1.005.tar.gz"
	titlecase_cmd_tar := exec.Command("curl", "-L", titlecase_tar_url, "-o", "package.tar.gz")
	err := titlecase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	titlecase_zip_url := "https://github.com/ap/titlecase/archive/refs/tags/v1.005.zip"
	titlecase_cmd_zip := exec.Command("curl", "-L", titlecase_zip_url, "-o", "package.zip")
	err = titlecase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	titlecase_bin_url := "https://github.com/ap/titlecase/archive/refs/tags/v1.005.bin"
	titlecase_cmd_bin := exec.Command("curl", "-L", titlecase_bin_url, "-o", "binary.bin")
	err = titlecase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	titlecase_src_url := "https://github.com/ap/titlecase/archive/refs/tags/v1.005.src.tar.gz"
	titlecase_cmd_src := exec.Command("curl", "-L", titlecase_src_url, "-o", "source.tar.gz")
	err = titlecase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	titlecase_cmd_direct := exec.Command("./binary")
	err = titlecase_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
