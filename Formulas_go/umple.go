package main

// Umple - Modeling tool/programming language that enables Model-Oriented Programming
// Homepage: https://www.umple.org

import (
	"fmt"
	
	"os/exec"
)

func installUmple() {
	// Método 1: Descargar y extraer .tar.gz
	umple_tar_url := "https://github.com/umple/umple/releases/download/v1.34.0/umple-1.34.0.7242.6b8819789.jar"
	umple_cmd_tar := exec.Command("curl", "-L", umple_tar_url, "-o", "package.tar.gz")
	err := umple_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	umple_zip_url := "https://github.com/umple/umple/releases/download/v1.34.0/umple-1.34.0.7242.6b8819789.jar"
	umple_cmd_zip := exec.Command("curl", "-L", umple_zip_url, "-o", "package.zip")
	err = umple_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	umple_bin_url := "https://github.com/umple/umple/releases/download/v1.34.0/umple-1.34.0.7242.6b8819789.jar"
	umple_cmd_bin := exec.Command("curl", "-L", umple_bin_url, "-o", "binary.bin")
	err = umple_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	umple_src_url := "https://github.com/umple/umple/releases/download/v1.34.0/umple-1.34.0.7242.6b8819789.jar"
	umple_cmd_src := exec.Command("curl", "-L", umple_src_url, "-o", "source.tar.gz")
	err = umple_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	umple_cmd_direct := exec.Command("./binary")
	err = umple_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
