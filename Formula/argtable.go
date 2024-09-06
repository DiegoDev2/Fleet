package main

// Argtable - ANSI C library for parsing GNU-style command-line options
// Homepage: https://argtable.sourceforge.io

import (
	"fmt"
	
	"os/exec"
)

func installArgtable() {
	// Método 1: Descargar y extraer .tar.gz
	argtable_tar_url := "https://downloads.sourceforge.net/project/argtable/argtable/argtable-2.13/argtable2-13.tar.gz"
	argtable_cmd_tar := exec.Command("curl", "-L", argtable_tar_url, "-o", "package.tar.gz")
	err := argtable_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	argtable_zip_url := "https://downloads.sourceforge.net/project/argtable/argtable/argtable-2.13/argtable2-13.zip"
	argtable_cmd_zip := exec.Command("curl", "-L", argtable_zip_url, "-o", "package.zip")
	err = argtable_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	argtable_bin_url := "https://downloads.sourceforge.net/project/argtable/argtable/argtable-2.13/argtable2-13.bin"
	argtable_cmd_bin := exec.Command("curl", "-L", argtable_bin_url, "-o", "binary.bin")
	err = argtable_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	argtable_src_url := "https://downloads.sourceforge.net/project/argtable/argtable/argtable-2.13/argtable2-13.src.tar.gz"
	argtable_cmd_src := exec.Command("curl", "-L", argtable_src_url, "-o", "source.tar.gz")
	err = argtable_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	argtable_cmd_direct := exec.Command("./binary")
	err = argtable_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
