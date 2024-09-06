package main

// VertX - Toolkit for building reactive applications on the JVM
// Homepage: https://vertx.io/

import (
	"fmt"
	
	"os/exec"
)

func installVertX() {
	// Método 1: Descargar y extraer .tar.gz
	vertx_tar_url := "https://search.maven.org/remotecontent?filepath=io/vertx/vertx-stack-manager/4.1.5/vertx-stack-manager-4.1.5-full.tar.gz"
	vertx_cmd_tar := exec.Command("curl", "-L", vertx_tar_url, "-o", "package.tar.gz")
	err := vertx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vertx_zip_url := "https://search.maven.org/remotecontent?filepath=io/vertx/vertx-stack-manager/4.1.5/vertx-stack-manager-4.1.5-full.zip"
	vertx_cmd_zip := exec.Command("curl", "-L", vertx_zip_url, "-o", "package.zip")
	err = vertx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vertx_bin_url := "https://search.maven.org/remotecontent?filepath=io/vertx/vertx-stack-manager/4.1.5/vertx-stack-manager-4.1.5-full.bin"
	vertx_cmd_bin := exec.Command("curl", "-L", vertx_bin_url, "-o", "binary.bin")
	err = vertx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vertx_src_url := "https://search.maven.org/remotecontent?filepath=io/vertx/vertx-stack-manager/4.1.5/vertx-stack-manager-4.1.5-full.src.tar.gz"
	vertx_cmd_src := exec.Command("curl", "-L", vertx_src_url, "-o", "source.tar.gz")
	err = vertx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vertx_cmd_direct := exec.Command("./binary")
	err = vertx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@17")
exec.Command("latte", "install", "openjdk@17")
}
