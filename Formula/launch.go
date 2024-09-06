package main

// Launch - Command-line launcher for macOS, in the spirit of `open`
// Homepage: https://sabi.net/nriley/software/#launch

import (
	"fmt"
	
	"os/exec"
)

func installLaunch() {
	// Método 1: Descargar y extraer .tar.gz
	launch_tar_url := "https://sabi.net/nriley/software/launch-1.2.5.tar.gz"
	launch_cmd_tar := exec.Command("curl", "-L", launch_tar_url, "-o", "package.tar.gz")
	err := launch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	launch_zip_url := "https://sabi.net/nriley/software/launch-1.2.5.zip"
	launch_cmd_zip := exec.Command("curl", "-L", launch_zip_url, "-o", "package.zip")
	err = launch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	launch_bin_url := "https://sabi.net/nriley/software/launch-1.2.5.bin"
	launch_cmd_bin := exec.Command("curl", "-L", launch_bin_url, "-o", "binary.bin")
	err = launch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	launch_src_url := "https://sabi.net/nriley/software/launch-1.2.5.src.tar.gz"
	launch_cmd_src := exec.Command("curl", "-L", launch_src_url, "-o", "source.tar.gz")
	err = launch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	launch_cmd_direct := exec.Command("./binary")
	err = launch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
