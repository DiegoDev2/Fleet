package main

// Masscan - TCP port scanner, scans entire Internet in under 5 minutes
// Homepage: https://github.com/robertdavidgraham/masscan/

import (
	"fmt"
	
	"os/exec"
)

func installMasscan() {
	// Método 1: Descargar y extraer .tar.gz
	masscan_tar_url := "https://github.com/robertdavidgraham/masscan/archive/refs/tags/1.3.2.tar.gz"
	masscan_cmd_tar := exec.Command("curl", "-L", masscan_tar_url, "-o", "package.tar.gz")
	err := masscan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	masscan_zip_url := "https://github.com/robertdavidgraham/masscan/archive/refs/tags/1.3.2.zip"
	masscan_cmd_zip := exec.Command("curl", "-L", masscan_zip_url, "-o", "package.zip")
	err = masscan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	masscan_bin_url := "https://github.com/robertdavidgraham/masscan/archive/refs/tags/1.3.2.bin"
	masscan_cmd_bin := exec.Command("curl", "-L", masscan_bin_url, "-o", "binary.bin")
	err = masscan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	masscan_src_url := "https://github.com/robertdavidgraham/masscan/archive/refs/tags/1.3.2.src.tar.gz"
	masscan_cmd_src := exec.Command("curl", "-L", masscan_src_url, "-o", "source.tar.gz")
	err = masscan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	masscan_cmd_direct := exec.Command("./binary")
	err = masscan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
