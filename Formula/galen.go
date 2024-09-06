package main

// Galen - Automated testing of look and feel for responsive websites
// Homepage: https://galenframework.com/

import (
	"fmt"
	
	"os/exec"
)

func installGalen() {
	// Método 1: Descargar y extraer .tar.gz
	galen_tar_url := "https://github.com/galenframework/galen/releases/download/galen-2.4.4/galen-bin-2.4.4.zip"
	galen_cmd_tar := exec.Command("curl", "-L", galen_tar_url, "-o", "package.tar.gz")
	err := galen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	galen_zip_url := "https://github.com/galenframework/galen/releases/download/galen-2.4.4/galen-bin-2.4.4.zip"
	galen_cmd_zip := exec.Command("curl", "-L", galen_zip_url, "-o", "package.zip")
	err = galen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	galen_bin_url := "https://github.com/galenframework/galen/releases/download/galen-2.4.4/galen-bin-2.4.4.zip"
	galen_cmd_bin := exec.Command("curl", "-L", galen_bin_url, "-o", "binary.bin")
	err = galen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	galen_src_url := "https://github.com/galenframework/galen/releases/download/galen-2.4.4/galen-bin-2.4.4.zip"
	galen_cmd_src := exec.Command("curl", "-L", galen_src_url, "-o", "source.tar.gz")
	err = galen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	galen_cmd_direct := exec.Command("./binary")
	err = galen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
