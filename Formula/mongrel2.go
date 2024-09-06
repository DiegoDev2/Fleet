package main

// Mongrel2 - Application, language, and network architecture agnostic web server
// Homepage: https://mongrel2.org/

import (
	"fmt"
	
	"os/exec"
)

func installMongrel2() {
	// Método 1: Descargar y extraer .tar.gz
	mongrel2_tar_url := "https://github.com/mongrel2/mongrel2/releases/download/v1.13.0/mongrel2-v1.13.0.tar.bz2"
	mongrel2_cmd_tar := exec.Command("curl", "-L", mongrel2_tar_url, "-o", "package.tar.gz")
	err := mongrel2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mongrel2_zip_url := "https://github.com/mongrel2/mongrel2/releases/download/v1.13.0/mongrel2-v1.13.0.tar.bz2"
	mongrel2_cmd_zip := exec.Command("curl", "-L", mongrel2_zip_url, "-o", "package.zip")
	err = mongrel2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mongrel2_bin_url := "https://github.com/mongrel2/mongrel2/releases/download/v1.13.0/mongrel2-v1.13.0.tar.bz2"
	mongrel2_cmd_bin := exec.Command("curl", "-L", mongrel2_bin_url, "-o", "binary.bin")
	err = mongrel2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mongrel2_src_url := "https://github.com/mongrel2/mongrel2/releases/download/v1.13.0/mongrel2-v1.13.0.tar.bz2"
	mongrel2_cmd_src := exec.Command("curl", "-L", mongrel2_src_url, "-o", "source.tar.gz")
	err = mongrel2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mongrel2_cmd_direct := exec.Command("./binary")
	err = mongrel2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: zeromq")
	exec.Command("latte", "install", "zeromq").Run()
}
