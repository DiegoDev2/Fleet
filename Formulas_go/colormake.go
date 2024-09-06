package main

// Colormake - Wrapper around make to colorize the output
// Homepage: https://github.com/pagekite/Colormake

import (
	"fmt"
	
	"os/exec"
)

func installColormake() {
	// Método 1: Descargar y extraer .tar.gz
	colormake_tar_url := "https://github.com/pagekite/Colormake/archive/refs/tags/0.9.20140503.tar.gz"
	colormake_cmd_tar := exec.Command("curl", "-L", colormake_tar_url, "-o", "package.tar.gz")
	err := colormake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	colormake_zip_url := "https://github.com/pagekite/Colormake/archive/refs/tags/0.9.20140503.zip"
	colormake_cmd_zip := exec.Command("curl", "-L", colormake_zip_url, "-o", "package.zip")
	err = colormake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	colormake_bin_url := "https://github.com/pagekite/Colormake/archive/refs/tags/0.9.20140503.bin"
	colormake_cmd_bin := exec.Command("curl", "-L", colormake_bin_url, "-o", "binary.bin")
	err = colormake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	colormake_src_url := "https://github.com/pagekite/Colormake/archive/refs/tags/0.9.20140503.src.tar.gz"
	colormake_cmd_src := exec.Command("curl", "-L", colormake_src_url, "-o", "source.tar.gz")
	err = colormake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	colormake_cmd_direct := exec.Command("./binary")
	err = colormake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
