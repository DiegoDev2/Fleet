package main

// Picat - Simple, and yet powerful, logic-based multi-paradigm programming language
// Homepage: http://picat-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installPicat() {
	// Método 1: Descargar y extraer .tar.gz
	picat_tar_url := "http://picat-lang.org/download/picat37_src.tar.gz"
	picat_cmd_tar := exec.Command("curl", "-L", picat_tar_url, "-o", "package.tar.gz")
	err := picat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	picat_zip_url := "http://picat-lang.org/download/picat37_src.zip"
	picat_cmd_zip := exec.Command("curl", "-L", picat_zip_url, "-o", "package.zip")
	err = picat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	picat_bin_url := "http://picat-lang.org/download/picat37_src.bin"
	picat_cmd_bin := exec.Command("curl", "-L", picat_bin_url, "-o", "binary.bin")
	err = picat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	picat_src_url := "http://picat-lang.org/download/picat37_src.src.tar.gz"
	picat_cmd_src := exec.Command("curl", "-L", picat_src_url, "-o", "source.tar.gz")
	err = picat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	picat_cmd_direct := exec.Command("./binary")
	err = picat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
