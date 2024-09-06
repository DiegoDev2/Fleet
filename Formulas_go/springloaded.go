package main

// SpringLoaded - Java agent to enable class reloading in a running JVM
// Homepage: https://github.com/spring-projects/spring-loaded

import (
	"fmt"
	
	"os/exec"
)

func installSpringLoaded() {
	// Método 1: Descargar y extraer .tar.gz
	springloaded_tar_url := "https://search.maven.org/remotecontent?filepath=org/springframework/springloaded/1.2.6.RELEASE/springloaded-1.2.6.RELEASE.jar"
	springloaded_cmd_tar := exec.Command("curl", "-L", springloaded_tar_url, "-o", "package.tar.gz")
	err := springloaded_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	springloaded_zip_url := "https://search.maven.org/remotecontent?filepath=org/springframework/springloaded/1.2.6.RELEASE/springloaded-1.2.6.RELEASE.jar"
	springloaded_cmd_zip := exec.Command("curl", "-L", springloaded_zip_url, "-o", "package.zip")
	err = springloaded_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	springloaded_bin_url := "https://search.maven.org/remotecontent?filepath=org/springframework/springloaded/1.2.6.RELEASE/springloaded-1.2.6.RELEASE.jar"
	springloaded_cmd_bin := exec.Command("curl", "-L", springloaded_bin_url, "-o", "binary.bin")
	err = springloaded_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	springloaded_src_url := "https://search.maven.org/remotecontent?filepath=org/springframework/springloaded/1.2.6.RELEASE/springloaded-1.2.6.RELEASE.jar"
	springloaded_cmd_src := exec.Command("curl", "-L", springloaded_src_url, "-o", "source.tar.gz")
	err = springloaded_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	springloaded_cmd_direct := exec.Command("./binary")
	err = springloaded_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
