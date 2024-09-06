package main

// Sbt - Build tool for Scala projects
// Homepage: https://www.scala-sbt.org/

import (
	"fmt"
	
	"os/exec"
)

func installSbt() {
	// Método 1: Descargar y extraer .tar.gz
	sbt_tar_url := "https://github.com/sbt/sbt/releases/download/v1.10.1/sbt-1.10.1.tgz"
	sbt_cmd_tar := exec.Command("curl", "-L", sbt_tar_url, "-o", "package.tar.gz")
	err := sbt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sbt_zip_url := "https://github.com/sbt/sbt/releases/download/v1.10.1/sbt-1.10.1.tgz"
	sbt_cmd_zip := exec.Command("curl", "-L", sbt_zip_url, "-o", "package.zip")
	err = sbt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sbt_bin_url := "https://github.com/sbt/sbt/releases/download/v1.10.1/sbt-1.10.1.tgz"
	sbt_cmd_bin := exec.Command("curl", "-L", sbt_bin_url, "-o", "binary.bin")
	err = sbt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sbt_src_url := "https://github.com/sbt/sbt/releases/download/v1.10.1/sbt-1.10.1.tgz"
	sbt_cmd_src := exec.Command("curl", "-L", sbt_src_url, "-o", "source.tar.gz")
	err = sbt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sbt_cmd_direct := exec.Command("./binary")
	err = sbt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
