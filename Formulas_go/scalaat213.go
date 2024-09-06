package main

// ScalaAT213 - JVM-based programming language
// Homepage: https://www.scala-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installScalaAT213() {
	// Método 1: Descargar y extraer .tar.gz
	scalaat213_tar_url := "https://downloads.lightbend.com/scala/2.13.14/scala-2.13.14.tgz"
	scalaat213_cmd_tar := exec.Command("curl", "-L", scalaat213_tar_url, "-o", "package.tar.gz")
	err := scalaat213_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scalaat213_zip_url := "https://downloads.lightbend.com/scala/2.13.14/scala-2.13.14.tgz"
	scalaat213_cmd_zip := exec.Command("curl", "-L", scalaat213_zip_url, "-o", "package.zip")
	err = scalaat213_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scalaat213_bin_url := "https://downloads.lightbend.com/scala/2.13.14/scala-2.13.14.tgz"
	scalaat213_cmd_bin := exec.Command("curl", "-L", scalaat213_bin_url, "-o", "binary.bin")
	err = scalaat213_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scalaat213_src_url := "https://downloads.lightbend.com/scala/2.13.14/scala-2.13.14.tgz"
	scalaat213_cmd_src := exec.Command("curl", "-L", scalaat213_src_url, "-o", "source.tar.gz")
	err = scalaat213_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scalaat213_cmd_direct := exec.Command("./binary")
	err = scalaat213_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
