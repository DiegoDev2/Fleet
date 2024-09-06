package main

// ConsoleBridge - Robot Operating System-independent package for logging
// Homepage: https://wiki.ros.org/console_bridge/

import (
	"fmt"
	
	"os/exec"
)

func installConsoleBridge() {
	// Método 1: Descargar y extraer .tar.gz
	consolebridge_tar_url := "https://github.com/ros/console_bridge/archive/refs/tags/1.0.2.tar.gz"
	consolebridge_cmd_tar := exec.Command("curl", "-L", consolebridge_tar_url, "-o", "package.tar.gz")
	err := consolebridge_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	consolebridge_zip_url := "https://github.com/ros/console_bridge/archive/refs/tags/1.0.2.zip"
	consolebridge_cmd_zip := exec.Command("curl", "-L", consolebridge_zip_url, "-o", "package.zip")
	err = consolebridge_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	consolebridge_bin_url := "https://github.com/ros/console_bridge/archive/refs/tags/1.0.2.bin"
	consolebridge_cmd_bin := exec.Command("curl", "-L", consolebridge_bin_url, "-o", "binary.bin")
	err = consolebridge_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	consolebridge_src_url := "https://github.com/ros/console_bridge/archive/refs/tags/1.0.2.src.tar.gz"
	consolebridge_cmd_src := exec.Command("curl", "-L", consolebridge_src_url, "-o", "source.tar.gz")
	err = consolebridge_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	consolebridge_cmd_direct := exec.Command("./binary")
	err = consolebridge_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
