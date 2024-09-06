package main

// Jbake - Java based static site/blog generator
// Homepage: https://jbake.org/

import (
	"fmt"
	
	"os/exec"
)

func installJbake() {
	// Método 1: Descargar y extraer .tar.gz
	jbake_tar_url := "https://github.com/jbake-org/jbake/releases/download/v2.6.7/jbake-2.6.7-bin.zip"
	jbake_cmd_tar := exec.Command("curl", "-L", jbake_tar_url, "-o", "package.tar.gz")
	err := jbake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jbake_zip_url := "https://github.com/jbake-org/jbake/releases/download/v2.6.7/jbake-2.6.7-bin.zip"
	jbake_cmd_zip := exec.Command("curl", "-L", jbake_zip_url, "-o", "package.zip")
	err = jbake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jbake_bin_url := "https://github.com/jbake-org/jbake/releases/download/v2.6.7/jbake-2.6.7-bin.zip"
	jbake_cmd_bin := exec.Command("curl", "-L", jbake_bin_url, "-o", "binary.bin")
	err = jbake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jbake_src_url := "https://github.com/jbake-org/jbake/releases/download/v2.6.7/jbake-2.6.7-bin.zip"
	jbake_cmd_src := exec.Command("curl", "-L", jbake_src_url, "-o", "source.tar.gz")
	err = jbake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jbake_cmd_direct := exec.Command("./binary")
	err = jbake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
