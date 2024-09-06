package main

// Clpbar - Command-line progress bar
// Homepage: https://clpbar.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installClpbar() {
	// Método 1: Descargar y extraer .tar.gz
	clpbar_tar_url := "https://downloads.sourceforge.net/project/clpbar/clpbar/bar-1.11.1/bar_1.11.1.tar.gz"
	clpbar_cmd_tar := exec.Command("curl", "-L", clpbar_tar_url, "-o", "package.tar.gz")
	err := clpbar_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clpbar_zip_url := "https://downloads.sourceforge.net/project/clpbar/clpbar/bar-1.11.1/bar_1.11.1.zip"
	clpbar_cmd_zip := exec.Command("curl", "-L", clpbar_zip_url, "-o", "package.zip")
	err = clpbar_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clpbar_bin_url := "https://downloads.sourceforge.net/project/clpbar/clpbar/bar-1.11.1/bar_1.11.1.bin"
	clpbar_cmd_bin := exec.Command("curl", "-L", clpbar_bin_url, "-o", "binary.bin")
	err = clpbar_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clpbar_src_url := "https://downloads.sourceforge.net/project/clpbar/clpbar/bar-1.11.1/bar_1.11.1.src.tar.gz"
	clpbar_cmd_src := exec.Command("curl", "-L", clpbar_src_url, "-o", "source.tar.gz")
	err = clpbar_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clpbar_cmd_direct := exec.Command("./binary")
	err = clpbar_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
