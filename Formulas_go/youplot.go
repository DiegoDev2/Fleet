package main

// Youplot - Command-line tool that draw plots on the terminal
// Homepage: https://github.com/red-data-tools/YouPlot/

import (
	"fmt"
	
	"os/exec"
)

func installYouplot() {
	// Método 1: Descargar y extraer .tar.gz
	youplot_tar_url := "https://github.com/red-data-tools/YouPlot/archive/refs/tags/v0.4.6.tar.gz"
	youplot_cmd_tar := exec.Command("curl", "-L", youplot_tar_url, "-o", "package.tar.gz")
	err := youplot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	youplot_zip_url := "https://github.com/red-data-tools/YouPlot/archive/refs/tags/v0.4.6.zip"
	youplot_cmd_zip := exec.Command("curl", "-L", youplot_zip_url, "-o", "package.zip")
	err = youplot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	youplot_bin_url := "https://github.com/red-data-tools/YouPlot/archive/refs/tags/v0.4.6.bin"
	youplot_cmd_bin := exec.Command("curl", "-L", youplot_bin_url, "-o", "binary.bin")
	err = youplot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	youplot_src_url := "https://github.com/red-data-tools/YouPlot/archive/refs/tags/v0.4.6.src.tar.gz"
	youplot_cmd_src := exec.Command("curl", "-L", youplot_src_url, "-o", "source.tar.gz")
	err = youplot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	youplot_cmd_direct := exec.Command("./binary")
	err = youplot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
