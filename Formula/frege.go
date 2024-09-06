package main

// Frege - Non-strict, functional programming language in the spirit of Haskell
// Homepage: https://github.com/Frege/frege/

import (
	"fmt"
	
	"os/exec"
)

func installFrege() {
	// Método 1: Descargar y extraer .tar.gz
	frege_tar_url := "https://github.com/Frege/frege/releases/download/3.24public/frege3.24.405.jar"
	frege_cmd_tar := exec.Command("curl", "-L", frege_tar_url, "-o", "package.tar.gz")
	err := frege_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	frege_zip_url := "https://github.com/Frege/frege/releases/download/3.24public/frege3.24.405.jar"
	frege_cmd_zip := exec.Command("curl", "-L", frege_zip_url, "-o", "package.zip")
	err = frege_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	frege_bin_url := "https://github.com/Frege/frege/releases/download/3.24public/frege3.24.405.jar"
	frege_cmd_bin := exec.Command("curl", "-L", frege_bin_url, "-o", "binary.bin")
	err = frege_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	frege_src_url := "https://github.com/Frege/frege/releases/download/3.24public/frege3.24.405.jar"
	frege_cmd_src := exec.Command("curl", "-L", frege_src_url, "-o", "source.tar.gz")
	err = frege_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	frege_cmd_direct := exec.Command("./binary")
	err = frege_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
