package main

// ArgpStandalone - Standalone version of arguments parsing functions from GLIBC
// Homepage: https://www.lysator.liu.se/~nisse/misc/

import (
	"fmt"
	
	"os/exec"
)

func installArgpStandalone() {
	// Método 1: Descargar y extraer .tar.gz
	argpstandalone_tar_url := "https://www.lysator.liu.se/~nisse/misc/argp-standalone-1.3.tar.gz"
	argpstandalone_cmd_tar := exec.Command("curl", "-L", argpstandalone_tar_url, "-o", "package.tar.gz")
	err := argpstandalone_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	argpstandalone_zip_url := "https://www.lysator.liu.se/~nisse/misc/argp-standalone-1.3.zip"
	argpstandalone_cmd_zip := exec.Command("curl", "-L", argpstandalone_zip_url, "-o", "package.zip")
	err = argpstandalone_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	argpstandalone_bin_url := "https://www.lysator.liu.se/~nisse/misc/argp-standalone-1.3.bin"
	argpstandalone_cmd_bin := exec.Command("curl", "-L", argpstandalone_bin_url, "-o", "binary.bin")
	err = argpstandalone_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	argpstandalone_src_url := "https://www.lysator.liu.se/~nisse/misc/argp-standalone-1.3.src.tar.gz"
	argpstandalone_cmd_src := exec.Command("curl", "-L", argpstandalone_src_url, "-o", "source.tar.gz")
	err = argpstandalone_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	argpstandalone_cmd_direct := exec.Command("./binary")
	err = argpstandalone_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
