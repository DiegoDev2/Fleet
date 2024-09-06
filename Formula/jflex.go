package main

// Jflex - Lexical analyzer generator for Java, written in Java
// Homepage: https://jflex.de/

import (
	"fmt"
	
	"os/exec"
)

func installJflex() {
	// Método 1: Descargar y extraer .tar.gz
	jflex_tar_url := "https://github.com/jflex-de/jflex/releases/download/v1.9.1/jflex-1.9.1.tar.gz"
	jflex_cmd_tar := exec.Command("curl", "-L", jflex_tar_url, "-o", "package.tar.gz")
	err := jflex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jflex_zip_url := "https://github.com/jflex-de/jflex/releases/download/v1.9.1/jflex-1.9.1.zip"
	jflex_cmd_zip := exec.Command("curl", "-L", jflex_zip_url, "-o", "package.zip")
	err = jflex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jflex_bin_url := "https://github.com/jflex-de/jflex/releases/download/v1.9.1/jflex-1.9.1.bin"
	jflex_cmd_bin := exec.Command("curl", "-L", jflex_bin_url, "-o", "binary.bin")
	err = jflex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jflex_src_url := "https://github.com/jflex-de/jflex/releases/download/v1.9.1/jflex-1.9.1.src.tar.gz"
	jflex_cmd_src := exec.Command("curl", "-L", jflex_src_url, "-o", "source.tar.gz")
	err = jflex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jflex_cmd_direct := exec.Command("./binary")
	err = jflex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
