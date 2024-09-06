package main

// Bwidget - Tcl/Tk script-only set of megawidgets to provide the developer additional tools
// Homepage: https://core.tcl-lang.org/bwidget/home

import (
	"fmt"
	
	"os/exec"
)

func installBwidget() {
	// Método 1: Descargar y extraer .tar.gz
	bwidget_tar_url := "https://downloads.sourceforge.net/project/tcllib/BWidget/1.9.16/bwidget-1.9.16.tar.gz"
	bwidget_cmd_tar := exec.Command("curl", "-L", bwidget_tar_url, "-o", "package.tar.gz")
	err := bwidget_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bwidget_zip_url := "https://downloads.sourceforge.net/project/tcllib/BWidget/1.9.16/bwidget-1.9.16.zip"
	bwidget_cmd_zip := exec.Command("curl", "-L", bwidget_zip_url, "-o", "package.zip")
	err = bwidget_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bwidget_bin_url := "https://downloads.sourceforge.net/project/tcllib/BWidget/1.9.16/bwidget-1.9.16.bin"
	bwidget_cmd_bin := exec.Command("curl", "-L", bwidget_bin_url, "-o", "binary.bin")
	err = bwidget_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bwidget_src_url := "https://downloads.sourceforge.net/project/tcllib/BWidget/1.9.16/bwidget-1.9.16.src.tar.gz"
	bwidget_cmd_src := exec.Command("curl", "-L", bwidget_src_url, "-o", "source.tar.gz")
	err = bwidget_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bwidget_cmd_direct := exec.Command("./binary")
	err = bwidget_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tcl-tk")
	exec.Command("latte", "install", "tcl-tk").Run()
}
