package main

// Clib - Package manager for C programming
// Homepage: https://github.com/clibs/clib

import (
	"fmt"
	
	"os/exec"
)

func installClib() {
	// Método 1: Descargar y extraer .tar.gz
	clib_tar_url := "https://github.com/clibs/clib/archive/refs/tags/2.8.7.tar.gz"
	clib_cmd_tar := exec.Command("curl", "-L", clib_tar_url, "-o", "package.tar.gz")
	err := clib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clib_zip_url := "https://github.com/clibs/clib/archive/refs/tags/2.8.7.zip"
	clib_cmd_zip := exec.Command("curl", "-L", clib_zip_url, "-o", "package.zip")
	err = clib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clib_bin_url := "https://github.com/clibs/clib/archive/refs/tags/2.8.7.bin"
	clib_cmd_bin := exec.Command("curl", "-L", clib_bin_url, "-o", "binary.bin")
	err = clib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clib_src_url := "https://github.com/clibs/clib/archive/refs/tags/2.8.7.src.tar.gz"
	clib_cmd_src := exec.Command("curl", "-L", clib_src_url, "-o", "source.tar.gz")
	err = clib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clib_cmd_direct := exec.Command("./binary")
	err = clib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
