package main

// Par2 - Parchive: Parity Archive Volume Set for data recovery
// Homepage: https://github.com/Parchive/par2cmdline

import (
	"fmt"
	
	"os/exec"
)

func installPar2() {
	// Método 1: Descargar y extraer .tar.gz
	par2_tar_url := "https://github.com/Parchive/par2cmdline/releases/download/v0.8.1/par2cmdline-0.8.1.tar.bz2"
	par2_cmd_tar := exec.Command("curl", "-L", par2_tar_url, "-o", "package.tar.gz")
	err := par2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	par2_zip_url := "https://github.com/Parchive/par2cmdline/releases/download/v0.8.1/par2cmdline-0.8.1.tar.bz2"
	par2_cmd_zip := exec.Command("curl", "-L", par2_zip_url, "-o", "package.zip")
	err = par2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	par2_bin_url := "https://github.com/Parchive/par2cmdline/releases/download/v0.8.1/par2cmdline-0.8.1.tar.bz2"
	par2_cmd_bin := exec.Command("curl", "-L", par2_bin_url, "-o", "binary.bin")
	err = par2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	par2_src_url := "https://github.com/Parchive/par2cmdline/releases/download/v0.8.1/par2cmdline-0.8.1.tar.bz2"
	par2_cmd_src := exec.Command("curl", "-L", par2_src_url, "-o", "source.tar.gz")
	err = par2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	par2_cmd_direct := exec.Command("./binary")
	err = par2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
