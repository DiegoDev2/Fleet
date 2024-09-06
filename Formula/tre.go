package main

// Tre - Lightweight, POSIX-compliant regular expression (regex) library
// Homepage: https://laurikari.net/tre/

import (
	"fmt"
	
	"os/exec"
)

func installTre() {
	// Método 1: Descargar y extraer .tar.gz
	tre_tar_url := "https://laurikari.net/tre/tre-0.8.0.tar.bz2"
	tre_cmd_tar := exec.Command("curl", "-L", tre_tar_url, "-o", "package.tar.gz")
	err := tre_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tre_zip_url := "https://laurikari.net/tre/tre-0.8.0.tar.bz2"
	tre_cmd_zip := exec.Command("curl", "-L", tre_zip_url, "-o", "package.zip")
	err = tre_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tre_bin_url := "https://laurikari.net/tre/tre-0.8.0.tar.bz2"
	tre_cmd_bin := exec.Command("curl", "-L", tre_bin_url, "-o", "binary.bin")
	err = tre_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tre_src_url := "https://laurikari.net/tre/tre-0.8.0.tar.bz2"
	tre_cmd_src := exec.Command("curl", "-L", tre_src_url, "-o", "source.tar.gz")
	err = tre_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tre_cmd_direct := exec.Command("./binary")
	err = tre_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
