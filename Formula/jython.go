package main

// Jython - Python implementation written in Java (successor to JPython)
// Homepage: https://www.jython.org/

import (
	"fmt"
	
	"os/exec"
)

func installJython() {
	// Método 1: Descargar y extraer .tar.gz
	jython_tar_url := "https://search.maven.org/remotecontent?filepath=org/python/jython-installer/2.7.4/jython-installer-2.7.4.jar"
	jython_cmd_tar := exec.Command("curl", "-L", jython_tar_url, "-o", "package.tar.gz")
	err := jython_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jython_zip_url := "https://search.maven.org/remotecontent?filepath=org/python/jython-installer/2.7.4/jython-installer-2.7.4.jar"
	jython_cmd_zip := exec.Command("curl", "-L", jython_zip_url, "-o", "package.zip")
	err = jython_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jython_bin_url := "https://search.maven.org/remotecontent?filepath=org/python/jython-installer/2.7.4/jython-installer-2.7.4.jar"
	jython_cmd_bin := exec.Command("curl", "-L", jython_bin_url, "-o", "binary.bin")
	err = jython_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jython_src_url := "https://search.maven.org/remotecontent?filepath=org/python/jython-installer/2.7.4/jython-installer-2.7.4.jar"
	jython_cmd_src := exec.Command("curl", "-L", jython_src_url, "-o", "source.tar.gz")
	err = jython_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jython_cmd_direct := exec.Command("./binary")
	err = jython_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
