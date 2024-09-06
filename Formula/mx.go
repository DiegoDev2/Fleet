package main

// Mx - Command-line tool used for the development of Graal projects
// Homepage: https://github.com/graalvm/mx

import (
	"fmt"
	
	"os/exec"
)

func installMx() {
	// Método 1: Descargar y extraer .tar.gz
	mx_tar_url := "https://github.com/graalvm/mx/archive/refs/tags/7.29.7.tar.gz"
	mx_cmd_tar := exec.Command("curl", "-L", mx_tar_url, "-o", "package.tar.gz")
	err := mx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mx_zip_url := "https://github.com/graalvm/mx/archive/refs/tags/7.29.7.zip"
	mx_cmd_zip := exec.Command("curl", "-L", mx_zip_url, "-o", "package.zip")
	err = mx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mx_bin_url := "https://github.com/graalvm/mx/archive/refs/tags/7.29.7.bin"
	mx_cmd_bin := exec.Command("curl", "-L", mx_bin_url, "-o", "binary.bin")
	err = mx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mx_src_url := "https://github.com/graalvm/mx/archive/refs/tags/7.29.7.src.tar.gz"
	mx_cmd_src := exec.Command("curl", "-L", mx_src_url, "-o", "source.tar.gz")
	err = mx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mx_cmd_direct := exec.Command("./binary")
	err = mx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
