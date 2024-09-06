package main

// Ipcalc - Calculate various network masks, etc. from a given IP address
// Homepage: https://jodies.de/ipcalc

import (
	"fmt"
	
	"os/exec"
)

func installIpcalc() {
	// Método 1: Descargar y extraer .tar.gz
	ipcalc_tar_url := "https://github.com/kjokjo/ipcalc/archive/refs/tags/0.51.tar.gz"
	ipcalc_cmd_tar := exec.Command("curl", "-L", ipcalc_tar_url, "-o", "package.tar.gz")
	err := ipcalc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ipcalc_zip_url := "https://github.com/kjokjo/ipcalc/archive/refs/tags/0.51.zip"
	ipcalc_cmd_zip := exec.Command("curl", "-L", ipcalc_zip_url, "-o", "package.zip")
	err = ipcalc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ipcalc_bin_url := "https://github.com/kjokjo/ipcalc/archive/refs/tags/0.51.bin"
	ipcalc_cmd_bin := exec.Command("curl", "-L", ipcalc_bin_url, "-o", "binary.bin")
	err = ipcalc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ipcalc_src_url := "https://github.com/kjokjo/ipcalc/archive/refs/tags/0.51.src.tar.gz"
	ipcalc_cmd_src := exec.Command("curl", "-L", ipcalc_src_url, "-o", "source.tar.gz")
	err = ipcalc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ipcalc_cmd_direct := exec.Command("./binary")
	err = ipcalc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
