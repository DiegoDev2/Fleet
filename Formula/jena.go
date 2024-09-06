package main

// Jena - Framework for building semantic web and linked data apps
// Homepage: https://jena.apache.org/

import (
	"fmt"
	
	"os/exec"
)

func installJena() {
	// Método 1: Descargar y extraer .tar.gz
	jena_tar_url := "https://www.apache.org/dyn/closer.lua?path=jena/binaries/apache-jena-5.1.0.tar.gz"
	jena_cmd_tar := exec.Command("curl", "-L", jena_tar_url, "-o", "package.tar.gz")
	err := jena_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jena_zip_url := "https://www.apache.org/dyn/closer.lua?path=jena/binaries/apache-jena-5.1.0.zip"
	jena_cmd_zip := exec.Command("curl", "-L", jena_zip_url, "-o", "package.zip")
	err = jena_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jena_bin_url := "https://www.apache.org/dyn/closer.lua?path=jena/binaries/apache-jena-5.1.0.bin"
	jena_cmd_bin := exec.Command("curl", "-L", jena_bin_url, "-o", "binary.bin")
	err = jena_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jena_src_url := "https://www.apache.org/dyn/closer.lua?path=jena/binaries/apache-jena-5.1.0.src.tar.gz"
	jena_cmd_src := exec.Command("curl", "-L", jena_src_url, "-o", "source.tar.gz")
	err = jena_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jena_cmd_direct := exec.Command("./binary")
	err = jena_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
