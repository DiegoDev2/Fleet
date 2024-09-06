package main

// Runit - Collection of tools for managing UNIX services
// Homepage: https://smarden.org/runit/

import (
	"fmt"
	
	"os/exec"
)

func installRunit() {
	// Método 1: Descargar y extraer .tar.gz
	runit_tar_url := "https://smarden.org/runit/runit-2.1.2.tar.gz"
	runit_cmd_tar := exec.Command("curl", "-L", runit_tar_url, "-o", "package.tar.gz")
	err := runit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	runit_zip_url := "https://smarden.org/runit/runit-2.1.2.zip"
	runit_cmd_zip := exec.Command("curl", "-L", runit_zip_url, "-o", "package.zip")
	err = runit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	runit_bin_url := "https://smarden.org/runit/runit-2.1.2.bin"
	runit_cmd_bin := exec.Command("curl", "-L", runit_bin_url, "-o", "binary.bin")
	err = runit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	runit_src_url := "https://smarden.org/runit/runit-2.1.2.src.tar.gz"
	runit_cmd_src := exec.Command("curl", "-L", runit_src_url, "-o", "source.tar.gz")
	err = runit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	runit_cmd_direct := exec.Command("./binary")
	err = runit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
