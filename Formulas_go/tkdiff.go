package main

// Tkdiff - Graphical side by side diff utility
// Homepage: https://tkdiff.sourceforge.io/

import (
	"fmt"
	
	"os/exec"
)

func installTkdiff() {
	// Método 1: Descargar y extraer .tar.gz
	tkdiff_tar_url := "https://downloads.sourceforge.net/project/tkdiff/tkdiff/5.7/tkdiff-5-7.zip"
	tkdiff_cmd_tar := exec.Command("curl", "-L", tkdiff_tar_url, "-o", "package.tar.gz")
	err := tkdiff_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tkdiff_zip_url := "https://downloads.sourceforge.net/project/tkdiff/tkdiff/5.7/tkdiff-5-7.zip"
	tkdiff_cmd_zip := exec.Command("curl", "-L", tkdiff_zip_url, "-o", "package.zip")
	err = tkdiff_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tkdiff_bin_url := "https://downloads.sourceforge.net/project/tkdiff/tkdiff/5.7/tkdiff-5-7.zip"
	tkdiff_cmd_bin := exec.Command("curl", "-L", tkdiff_bin_url, "-o", "binary.bin")
	err = tkdiff_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tkdiff_src_url := "https://downloads.sourceforge.net/project/tkdiff/tkdiff/5.7/tkdiff-5-7.zip"
	tkdiff_cmd_src := exec.Command("curl", "-L", tkdiff_src_url, "-o", "source.tar.gz")
	err = tkdiff_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tkdiff_cmd_direct := exec.Command("./binary")
	err = tkdiff_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tcl-tk")
exec.Command("latte", "install", "tcl-tk")
}
