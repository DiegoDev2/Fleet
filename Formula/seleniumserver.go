package main

// SeleniumServer - Browser automation for testing purposes
// Homepage: https://www.selenium.dev/

import (
	"fmt"
	
	"os/exec"
)

func installSeleniumServer() {
	// Método 1: Descargar y extraer .tar.gz
	seleniumserver_tar_url := "https://github.com/SeleniumHQ/selenium/releases/download/selenium-4.24.0/selenium-server-4.24.0.jar"
	seleniumserver_cmd_tar := exec.Command("curl", "-L", seleniumserver_tar_url, "-o", "package.tar.gz")
	err := seleniumserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	seleniumserver_zip_url := "https://github.com/SeleniumHQ/selenium/releases/download/selenium-4.24.0/selenium-server-4.24.0.jar"
	seleniumserver_cmd_zip := exec.Command("curl", "-L", seleniumserver_zip_url, "-o", "package.zip")
	err = seleniumserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	seleniumserver_bin_url := "https://github.com/SeleniumHQ/selenium/releases/download/selenium-4.24.0/selenium-server-4.24.0.jar"
	seleniumserver_cmd_bin := exec.Command("curl", "-L", seleniumserver_bin_url, "-o", "binary.bin")
	err = seleniumserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	seleniumserver_src_url := "https://github.com/SeleniumHQ/selenium/releases/download/selenium-4.24.0/selenium-server-4.24.0.jar"
	seleniumserver_cmd_src := exec.Command("curl", "-L", seleniumserver_src_url, "-o", "source.tar.gz")
	err = seleniumserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	seleniumserver_cmd_direct := exec.Command("./binary")
	err = seleniumserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
