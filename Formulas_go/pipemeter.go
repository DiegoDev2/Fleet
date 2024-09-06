package main

// Pipemeter - Shows speed of data moving from input to output
// Homepage: https://launchpad.net/pipemeter

import (
	"fmt"
	
	"os/exec"
)

func installPipemeter() {
	// Método 1: Descargar y extraer .tar.gz
	pipemeter_tar_url := "https://launchpad.net/pipemeter/trunk/1.1.5/+download/pipemeter-1.1.5.tar.gz"
	pipemeter_cmd_tar := exec.Command("curl", "-L", pipemeter_tar_url, "-o", "package.tar.gz")
	err := pipemeter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pipemeter_zip_url := "https://launchpad.net/pipemeter/trunk/1.1.5/+download/pipemeter-1.1.5.zip"
	pipemeter_cmd_zip := exec.Command("curl", "-L", pipemeter_zip_url, "-o", "package.zip")
	err = pipemeter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pipemeter_bin_url := "https://launchpad.net/pipemeter/trunk/1.1.5/+download/pipemeter-1.1.5.bin"
	pipemeter_cmd_bin := exec.Command("curl", "-L", pipemeter_bin_url, "-o", "binary.bin")
	err = pipemeter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pipemeter_src_url := "https://launchpad.net/pipemeter/trunk/1.1.5/+download/pipemeter-1.1.5.src.tar.gz"
	pipemeter_cmd_src := exec.Command("curl", "-L", pipemeter_src_url, "-o", "source.tar.gz")
	err = pipemeter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pipemeter_cmd_direct := exec.Command("./binary")
	err = pipemeter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
