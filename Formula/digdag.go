package main

// Digdag - Workload Automation System
// Homepage: https://www.digdag.io/

import (
	"fmt"
	
	"os/exec"
)

func installDigdag() {
	// Método 1: Descargar y extraer .tar.gz
	digdag_tar_url := "https://dl.digdag.io/digdag-0.10.5.1.jar"
	digdag_cmd_tar := exec.Command("curl", "-L", digdag_tar_url, "-o", "package.tar.gz")
	err := digdag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	digdag_zip_url := "https://dl.digdag.io/digdag-0.10.5.1.jar"
	digdag_cmd_zip := exec.Command("curl", "-L", digdag_zip_url, "-o", "package.zip")
	err = digdag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	digdag_bin_url := "https://dl.digdag.io/digdag-0.10.5.1.jar"
	digdag_cmd_bin := exec.Command("curl", "-L", digdag_bin_url, "-o", "binary.bin")
	err = digdag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	digdag_src_url := "https://dl.digdag.io/digdag-0.10.5.1.jar"
	digdag_cmd_src := exec.Command("curl", "-L", digdag_src_url, "-o", "source.tar.gz")
	err = digdag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	digdag_cmd_direct := exec.Command("./binary")
	err = digdag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@11")
	exec.Command("latte", "install", "openjdk@11").Run()
}
