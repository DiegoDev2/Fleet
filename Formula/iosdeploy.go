package main

// IosDeploy - Install and debug iPhone apps from the command-line
// Homepage: https://github.com/ios-control/ios-deploy

import (
	"fmt"
	
	"os/exec"
)

func installIosDeploy() {
	// Método 1: Descargar y extraer .tar.gz
	iosdeploy_tar_url := "https://github.com/ios-control/ios-deploy/archive/refs/tags/1.12.2.tar.gz"
	iosdeploy_cmd_tar := exec.Command("curl", "-L", iosdeploy_tar_url, "-o", "package.tar.gz")
	err := iosdeploy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iosdeploy_zip_url := "https://github.com/ios-control/ios-deploy/archive/refs/tags/1.12.2.zip"
	iosdeploy_cmd_zip := exec.Command("curl", "-L", iosdeploy_zip_url, "-o", "package.zip")
	err = iosdeploy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iosdeploy_bin_url := "https://github.com/ios-control/ios-deploy/archive/refs/tags/1.12.2.bin"
	iosdeploy_cmd_bin := exec.Command("curl", "-L", iosdeploy_bin_url, "-o", "binary.bin")
	err = iosdeploy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iosdeploy_src_url := "https://github.com/ios-control/ios-deploy/archive/refs/tags/1.12.2.src.tar.gz"
	iosdeploy_cmd_src := exec.Command("curl", "-L", iosdeploy_src_url, "-o", "source.tar.gz")
	err = iosdeploy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iosdeploy_cmd_direct := exec.Command("./binary")
	err = iosdeploy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
