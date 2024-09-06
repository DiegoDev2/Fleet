package main

// Glkterm - Terminal-window Glk library
// Homepage: https://www.eblong.com/zarf/glk/

import (
	"fmt"
	
	"os/exec"
)

func installGlkterm() {
	// Método 1: Descargar y extraer .tar.gz
	glkterm_tar_url := "https://www.eblong.com/zarf/glk/glkterm-104.tar.gz"
	glkterm_cmd_tar := exec.Command("curl", "-L", glkterm_tar_url, "-o", "package.tar.gz")
	err := glkterm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	glkterm_zip_url := "https://www.eblong.com/zarf/glk/glkterm-104.zip"
	glkterm_cmd_zip := exec.Command("curl", "-L", glkterm_zip_url, "-o", "package.zip")
	err = glkterm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	glkterm_bin_url := "https://www.eblong.com/zarf/glk/glkterm-104.bin"
	glkterm_cmd_bin := exec.Command("curl", "-L", glkterm_bin_url, "-o", "binary.bin")
	err = glkterm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	glkterm_src_url := "https://www.eblong.com/zarf/glk/glkterm-104.src.tar.gz"
	glkterm_cmd_src := exec.Command("curl", "-L", glkterm_src_url, "-o", "source.tar.gz")
	err = glkterm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	glkterm_cmd_direct := exec.Command("./binary")
	err = glkterm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
