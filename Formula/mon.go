package main

// Mon - Monitor hosts/services/whatever and alert about problems
// Homepage: https://github.com/tj/mon

import (
	"fmt"
	
	"os/exec"
)

func installMon() {
	// Método 1: Descargar y extraer .tar.gz
	mon_tar_url := "https://github.com/tj/mon/archive/refs/tags/1.2.3.tar.gz"
	mon_cmd_tar := exec.Command("curl", "-L", mon_tar_url, "-o", "package.tar.gz")
	err := mon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mon_zip_url := "https://github.com/tj/mon/archive/refs/tags/1.2.3.zip"
	mon_cmd_zip := exec.Command("curl", "-L", mon_zip_url, "-o", "package.zip")
	err = mon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mon_bin_url := "https://github.com/tj/mon/archive/refs/tags/1.2.3.bin"
	mon_cmd_bin := exec.Command("curl", "-L", mon_bin_url, "-o", "binary.bin")
	err = mon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mon_src_url := "https://github.com/tj/mon/archive/refs/tags/1.2.3.src.tar.gz"
	mon_cmd_src := exec.Command("curl", "-L", mon_src_url, "-o", "source.tar.gz")
	err = mon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mon_cmd_direct := exec.Command("./binary")
	err = mon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
