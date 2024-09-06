package main

// Micronaut - Modern JVM-based framework for building modular microservices
// Homepage: https://micronaut.io/

import (
	"fmt"
	
	"os/exec"
)

func installMicronaut() {
	// Método 1: Descargar y extraer .tar.gz
	micronaut_tar_url := "https://github.com/micronaut-projects/micronaut-starter/archive/refs/tags/v4.6.1.tar.gz"
	micronaut_cmd_tar := exec.Command("curl", "-L", micronaut_tar_url, "-o", "package.tar.gz")
	err := micronaut_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	micronaut_zip_url := "https://github.com/micronaut-projects/micronaut-starter/archive/refs/tags/v4.6.1.zip"
	micronaut_cmd_zip := exec.Command("curl", "-L", micronaut_zip_url, "-o", "package.zip")
	err = micronaut_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	micronaut_bin_url := "https://github.com/micronaut-projects/micronaut-starter/archive/refs/tags/v4.6.1.bin"
	micronaut_cmd_bin := exec.Command("curl", "-L", micronaut_bin_url, "-o", "binary.bin")
	err = micronaut_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	micronaut_src_url := "https://github.com/micronaut-projects/micronaut-starter/archive/refs/tags/v4.6.1.src.tar.gz"
	micronaut_cmd_src := exec.Command("curl", "-L", micronaut_src_url, "-o", "source.tar.gz")
	err = micronaut_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	micronaut_cmd_direct := exec.Command("./binary")
	err = micronaut_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gradle")
	exec.Command("latte", "install", "gradle").Run()
	fmt.Println("Instalando dependencia: openjdk@17")
	exec.Command("latte", "install", "openjdk@17").Run()
}
