package main

// Jdupes - Duplicate file finder and an enhanced fork of 'fdupes'
// Homepage: https://codeberg.org/jbruchon/jdupes

import (
	"fmt"
	
	"os/exec"
)

func installJdupes() {
	// Método 1: Descargar y extraer .tar.gz
	jdupes_tar_url := "https://codeberg.org/jbruchon/jdupes/archive/v1.28.0.tar.gz"
	jdupes_cmd_tar := exec.Command("curl", "-L", jdupes_tar_url, "-o", "package.tar.gz")
	err := jdupes_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jdupes_zip_url := "https://codeberg.org/jbruchon/jdupes/archive/v1.28.0.zip"
	jdupes_cmd_zip := exec.Command("curl", "-L", jdupes_zip_url, "-o", "package.zip")
	err = jdupes_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jdupes_bin_url := "https://codeberg.org/jbruchon/jdupes/archive/v1.28.0.bin"
	jdupes_cmd_bin := exec.Command("curl", "-L", jdupes_bin_url, "-o", "binary.bin")
	err = jdupes_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jdupes_src_url := "https://codeberg.org/jbruchon/jdupes/archive/v1.28.0.src.tar.gz"
	jdupes_cmd_src := exec.Command("curl", "-L", jdupes_src_url, "-o", "source.tar.gz")
	err = jdupes_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jdupes_cmd_direct := exec.Command("./binary")
	err = jdupes_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
