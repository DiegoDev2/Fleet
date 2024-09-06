package main

// Scala - JVM-based programming language
// Homepage: https://www.scala-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installScala() {
	// Método 1: Descargar y extraer .tar.gz
	scala_tar_url := "https://github.com/lampepfl/dotty/releases/download/3.5.0/scala3-3.5.0.tar.gz"
	scala_cmd_tar := exec.Command("curl", "-L", scala_tar_url, "-o", "package.tar.gz")
	err := scala_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scala_zip_url := "https://github.com/lampepfl/dotty/releases/download/3.5.0/scala3-3.5.0.zip"
	scala_cmd_zip := exec.Command("curl", "-L", scala_zip_url, "-o", "package.zip")
	err = scala_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scala_bin_url := "https://github.com/lampepfl/dotty/releases/download/3.5.0/scala3-3.5.0.bin"
	scala_cmd_bin := exec.Command("curl", "-L", scala_bin_url, "-o", "binary.bin")
	err = scala_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scala_src_url := "https://github.com/lampepfl/dotty/releases/download/3.5.0/scala3-3.5.0.src.tar.gz"
	scala_cmd_src := exec.Command("curl", "-L", scala_src_url, "-o", "source.tar.gz")
	err = scala_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scala_cmd_direct := exec.Command("./binary")
	err = scala_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
