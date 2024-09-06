package main

// Jolie - Service-oriented programming language
// Homepage: https://www.jolie-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installJolie() {
	// Método 1: Descargar y extraer .tar.gz
	jolie_tar_url := "https://github.com/jolie/jolie/releases/download/v1.12.1/jolie-1.12.1.jar"
	jolie_cmd_tar := exec.Command("curl", "-L", jolie_tar_url, "-o", "package.tar.gz")
	err := jolie_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jolie_zip_url := "https://github.com/jolie/jolie/releases/download/v1.12.1/jolie-1.12.1.jar"
	jolie_cmd_zip := exec.Command("curl", "-L", jolie_zip_url, "-o", "package.zip")
	err = jolie_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jolie_bin_url := "https://github.com/jolie/jolie/releases/download/v1.12.1/jolie-1.12.1.jar"
	jolie_cmd_bin := exec.Command("curl", "-L", jolie_bin_url, "-o", "binary.bin")
	err = jolie_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jolie_src_url := "https://github.com/jolie/jolie/releases/download/v1.12.1/jolie-1.12.1.jar"
	jolie_cmd_src := exec.Command("curl", "-L", jolie_src_url, "-o", "source.tar.gz")
	err = jolie_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jolie_cmd_direct := exec.Command("./binary")
	err = jolie_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
