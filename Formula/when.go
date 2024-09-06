package main

// When - Tiny personal calendar
// Homepage: https://www.lightandmatter.com/when/when.html

import (
	"fmt"
	
	"os/exec"
)

func installWhen() {
	// Método 1: Descargar y extraer .tar.gz
	when_tar_url := "https://bitbucket.org/ben-crowell/when/get/1.1.45.tar.bz2"
	when_cmd_tar := exec.Command("curl", "-L", when_tar_url, "-o", "package.tar.gz")
	err := when_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	when_zip_url := "https://bitbucket.org/ben-crowell/when/get/1.1.45.tar.bz2"
	when_cmd_zip := exec.Command("curl", "-L", when_zip_url, "-o", "package.zip")
	err = when_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	when_bin_url := "https://bitbucket.org/ben-crowell/when/get/1.1.45.tar.bz2"
	when_cmd_bin := exec.Command("curl", "-L", when_bin_url, "-o", "binary.bin")
	err = when_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	when_src_url := "https://bitbucket.org/ben-crowell/when/get/1.1.45.tar.bz2"
	when_cmd_src := exec.Command("curl", "-L", when_src_url, "-o", "source.tar.gz")
	err = when_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	when_cmd_direct := exec.Command("./binary")
	err = when_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
