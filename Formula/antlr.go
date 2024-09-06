package main

// Antlr - ANother Tool for Language Recognition
// Homepage: https://www.antlr.org/

import (
	"fmt"
	
	"os/exec"
)

func installAntlr() {
	// Método 1: Descargar y extraer .tar.gz
	antlr_tar_url := "https://www.antlr.org/download/antlr-4.13.2-complete.jar"
	antlr_cmd_tar := exec.Command("curl", "-L", antlr_tar_url, "-o", "package.tar.gz")
	err := antlr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	antlr_zip_url := "https://www.antlr.org/download/antlr-4.13.2-complete.jar"
	antlr_cmd_zip := exec.Command("curl", "-L", antlr_zip_url, "-o", "package.zip")
	err = antlr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	antlr_bin_url := "https://www.antlr.org/download/antlr-4.13.2-complete.jar"
	antlr_cmd_bin := exec.Command("curl", "-L", antlr_bin_url, "-o", "binary.bin")
	err = antlr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	antlr_src_url := "https://www.antlr.org/download/antlr-4.13.2-complete.jar"
	antlr_cmd_src := exec.Command("curl", "-L", antlr_src_url, "-o", "source.tar.gz")
	err = antlr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	antlr_cmd_direct := exec.Command("./binary")
	err = antlr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
