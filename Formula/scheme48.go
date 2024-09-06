package main

// Scheme48 - Scheme byte-code interpreter
// Homepage: https://www.s48.org/

import (
	"fmt"
	
	"os/exec"
)

func installScheme48() {
	// Método 1: Descargar y extraer .tar.gz
	scheme48_tar_url := "https://s48.org/1.9.2/scheme48-1.9.2.tgz"
	scheme48_cmd_tar := exec.Command("curl", "-L", scheme48_tar_url, "-o", "package.tar.gz")
	err := scheme48_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scheme48_zip_url := "https://s48.org/1.9.2/scheme48-1.9.2.tgz"
	scheme48_cmd_zip := exec.Command("curl", "-L", scheme48_zip_url, "-o", "package.zip")
	err = scheme48_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scheme48_bin_url := "https://s48.org/1.9.2/scheme48-1.9.2.tgz"
	scheme48_cmd_bin := exec.Command("curl", "-L", scheme48_bin_url, "-o", "binary.bin")
	err = scheme48_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scheme48_src_url := "https://s48.org/1.9.2/scheme48-1.9.2.tgz"
	scheme48_cmd_src := exec.Command("curl", "-L", scheme48_src_url, "-o", "source.tar.gz")
	err = scheme48_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scheme48_cmd_direct := exec.Command("./binary")
	err = scheme48_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
