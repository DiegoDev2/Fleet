package main

// Smartmontools - SMART hard drive monitoring
// Homepage: https://www.smartmontools.org/

import (
	"fmt"
	
	"os/exec"
)

func installSmartmontools() {
	// Método 1: Descargar y extraer .tar.gz
	smartmontools_tar_url := "https://downloads.sourceforge.net/project/smartmontools/smartmontools/7.4/smartmontools-7.4.tar.gz"
	smartmontools_cmd_tar := exec.Command("curl", "-L", smartmontools_tar_url, "-o", "package.tar.gz")
	err := smartmontools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smartmontools_zip_url := "https://downloads.sourceforge.net/project/smartmontools/smartmontools/7.4/smartmontools-7.4.zip"
	smartmontools_cmd_zip := exec.Command("curl", "-L", smartmontools_zip_url, "-o", "package.zip")
	err = smartmontools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smartmontools_bin_url := "https://downloads.sourceforge.net/project/smartmontools/smartmontools/7.4/smartmontools-7.4.bin"
	smartmontools_cmd_bin := exec.Command("curl", "-L", smartmontools_bin_url, "-o", "binary.bin")
	err = smartmontools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smartmontools_src_url := "https://downloads.sourceforge.net/project/smartmontools/smartmontools/7.4/smartmontools-7.4.src.tar.gz"
	smartmontools_cmd_src := exec.Command("curl", "-L", smartmontools_src_url, "-o", "source.tar.gz")
	err = smartmontools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smartmontools_cmd_direct := exec.Command("./binary")
	err = smartmontools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
