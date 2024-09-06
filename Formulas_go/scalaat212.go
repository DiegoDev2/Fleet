package main

// ScalaAT212 - JVM-based programming language
// Homepage: https://www.scala-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installScalaAT212() {
	// Método 1: Descargar y extraer .tar.gz
	scalaat212_tar_url := "https://downloads.lightbend.com/scala/2.12.20/scala-2.12.20.tgz"
	scalaat212_cmd_tar := exec.Command("curl", "-L", scalaat212_tar_url, "-o", "package.tar.gz")
	err := scalaat212_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scalaat212_zip_url := "https://downloads.lightbend.com/scala/2.12.20/scala-2.12.20.tgz"
	scalaat212_cmd_zip := exec.Command("curl", "-L", scalaat212_zip_url, "-o", "package.zip")
	err = scalaat212_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scalaat212_bin_url := "https://downloads.lightbend.com/scala/2.12.20/scala-2.12.20.tgz"
	scalaat212_cmd_bin := exec.Command("curl", "-L", scalaat212_bin_url, "-o", "binary.bin")
	err = scalaat212_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scalaat212_src_url := "https://downloads.lightbend.com/scala/2.12.20/scala-2.12.20.tgz"
	scalaat212_cmd_src := exec.Command("curl", "-L", scalaat212_src_url, "-o", "source.tar.gz")
	err = scalaat212_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scalaat212_cmd_direct := exec.Command("./binary")
	err = scalaat212_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
