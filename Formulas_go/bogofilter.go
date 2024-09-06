package main

// Bogofilter - Mail filter via statistical analysis
// Homepage: https://bogofilter.sourceforge.io

import (
	"fmt"
	
	"os/exec"
)

func installBogofilter() {
	// Método 1: Descargar y extraer .tar.gz
	bogofilter_tar_url := "https://downloads.sourceforge.net/project/bogofilter/bogofilter-stable/bogofilter-1.2.5.tar.xz"
	bogofilter_cmd_tar := exec.Command("curl", "-L", bogofilter_tar_url, "-o", "package.tar.gz")
	err := bogofilter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bogofilter_zip_url := "https://downloads.sourceforge.net/project/bogofilter/bogofilter-stable/bogofilter-1.2.5.tar.xz"
	bogofilter_cmd_zip := exec.Command("curl", "-L", bogofilter_zip_url, "-o", "package.zip")
	err = bogofilter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bogofilter_bin_url := "https://downloads.sourceforge.net/project/bogofilter/bogofilter-stable/bogofilter-1.2.5.tar.xz"
	bogofilter_cmd_bin := exec.Command("curl", "-L", bogofilter_bin_url, "-o", "binary.bin")
	err = bogofilter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bogofilter_src_url := "https://downloads.sourceforge.net/project/bogofilter/bogofilter-stable/bogofilter-1.2.5.tar.xz"
	bogofilter_cmd_src := exec.Command("curl", "-L", bogofilter_src_url, "-o", "source.tar.gz")
	err = bogofilter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bogofilter_cmd_direct := exec.Command("./binary")
	err = bogofilter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
