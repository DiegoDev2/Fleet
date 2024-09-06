package main

// Newman - Command-line collection runner for Postman
// Homepage: https://www.getpostman.com

import (
	"fmt"
	
	"os/exec"
)

func installNewman() {
	// Método 1: Descargar y extraer .tar.gz
	newman_tar_url := "https://registry.npmjs.org/newman/-/newman-6.2.1.tgz"
	newman_cmd_tar := exec.Command("curl", "-L", newman_tar_url, "-o", "package.tar.gz")
	err := newman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	newman_zip_url := "https://registry.npmjs.org/newman/-/newman-6.2.1.tgz"
	newman_cmd_zip := exec.Command("curl", "-L", newman_zip_url, "-o", "package.zip")
	err = newman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	newman_bin_url := "https://registry.npmjs.org/newman/-/newman-6.2.1.tgz"
	newman_cmd_bin := exec.Command("curl", "-L", newman_bin_url, "-o", "binary.bin")
	err = newman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	newman_src_url := "https://registry.npmjs.org/newman/-/newman-6.2.1.tgz"
	newman_cmd_src := exec.Command("curl", "-L", newman_src_url, "-o", "source.tar.gz")
	err = newman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	newman_cmd_direct := exec.Command("./binary")
	err = newman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
