package main

// Jshon - Parse, read, and create JSON from the shell
// Homepage: http://kmkeen.com/jshon/

import (
	"fmt"
	
	"os/exec"
)

func installJshon() {
	// Método 1: Descargar y extraer .tar.gz
	jshon_tar_url := "https://github.com/keenerd/jshon/archive/refs/tags/20131105.tar.gz"
	jshon_cmd_tar := exec.Command("curl", "-L", jshon_tar_url, "-o", "package.tar.gz")
	err := jshon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jshon_zip_url := "https://github.com/keenerd/jshon/archive/refs/tags/20131105.zip"
	jshon_cmd_zip := exec.Command("curl", "-L", jshon_zip_url, "-o", "package.zip")
	err = jshon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jshon_bin_url := "https://github.com/keenerd/jshon/archive/refs/tags/20131105.bin"
	jshon_cmd_bin := exec.Command("curl", "-L", jshon_bin_url, "-o", "binary.bin")
	err = jshon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jshon_src_url := "https://github.com/keenerd/jshon/archive/refs/tags/20131105.src.tar.gz"
	jshon_cmd_src := exec.Command("curl", "-L", jshon_src_url, "-o", "source.tar.gz")
	err = jshon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jshon_cmd_direct := exec.Command("./binary")
	err = jshon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
}
