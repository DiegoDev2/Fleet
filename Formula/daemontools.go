package main

// Daemontools - Collection of tools for managing UNIX services
// Homepage: https://cr.yp.to/daemontools.html

import (
	"fmt"
	
	"os/exec"
)

func installDaemontools() {
	// Método 1: Descargar y extraer .tar.gz
	daemontools_tar_url := "https://cr.yp.to/daemontools/daemontools-0.76.tar.gz"
	daemontools_cmd_tar := exec.Command("curl", "-L", daemontools_tar_url, "-o", "package.tar.gz")
	err := daemontools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	daemontools_zip_url := "https://cr.yp.to/daemontools/daemontools-0.76.zip"
	daemontools_cmd_zip := exec.Command("curl", "-L", daemontools_zip_url, "-o", "package.zip")
	err = daemontools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	daemontools_bin_url := "https://cr.yp.to/daemontools/daemontools-0.76.bin"
	daemontools_cmd_bin := exec.Command("curl", "-L", daemontools_bin_url, "-o", "binary.bin")
	err = daemontools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	daemontools_src_url := "https://cr.yp.to/daemontools/daemontools-0.76.src.tar.gz"
	daemontools_cmd_src := exec.Command("curl", "-L", daemontools_src_url, "-o", "source.tar.gz")
	err = daemontools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	daemontools_cmd_direct := exec.Command("./binary")
	err = daemontools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
