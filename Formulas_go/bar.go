package main

// Bar - Provide progress bars for shell scripts
// Homepage: http://www.theiling.de/projects/bar.html

import (
	"fmt"
	
	"os/exec"
)

func installBar() {
	// Método 1: Descargar y extraer .tar.gz
	bar_tar_url := "http://www.theiling.de/downloads/bar-1.4-src.tar.bz2"
	bar_cmd_tar := exec.Command("curl", "-L", bar_tar_url, "-o", "package.tar.gz")
	err := bar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bar_zip_url := "http://www.theiling.de/downloads/bar-1.4-src.tar.bz2"
	bar_cmd_zip := exec.Command("curl", "-L", bar_zip_url, "-o", "package.zip")
	err = bar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bar_bin_url := "http://www.theiling.de/downloads/bar-1.4-src.tar.bz2"
	bar_cmd_bin := exec.Command("curl", "-L", bar_bin_url, "-o", "binary.bin")
	err = bar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bar_src_url := "http://www.theiling.de/downloads/bar-1.4-src.tar.bz2"
	bar_cmd_src := exec.Command("curl", "-L", bar_src_url, "-o", "source.tar.gz")
	err = bar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bar_cmd_direct := exec.Command("./binary")
	err = bar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
