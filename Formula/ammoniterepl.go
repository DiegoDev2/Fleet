package main

// AmmoniteRepl - Ammonite is a cleanroom re-implementation of the Scala REPL
// Homepage: https://ammonite.io/

import (
	"fmt"
	
	"os/exec"
)

func installAmmoniteRepl() {
	// Método 1: Descargar y extraer .tar.gz
	ammoniterepl_tar_url := "https://github.com/com-lihaoyi/Ammonite/releases/download/3.0.0-M2/3.3-3.0.0-M2"
	ammoniterepl_cmd_tar := exec.Command("curl", "-L", ammoniterepl_tar_url, "-o", "package.tar.gz")
	err := ammoniterepl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ammoniterepl_zip_url := "https://github.com/com-lihaoyi/Ammonite/releases/download/3.0.0-M2/3.3-3.0.0-M2"
	ammoniterepl_cmd_zip := exec.Command("curl", "-L", ammoniterepl_zip_url, "-o", "package.zip")
	err = ammoniterepl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ammoniterepl_bin_url := "https://github.com/com-lihaoyi/Ammonite/releases/download/3.0.0-M2/3.3-3.0.0-M2"
	ammoniterepl_cmd_bin := exec.Command("curl", "-L", ammoniterepl_bin_url, "-o", "binary.bin")
	err = ammoniterepl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ammoniterepl_src_url := "https://github.com/com-lihaoyi/Ammonite/releases/download/3.0.0-M2/3.3-3.0.0-M2"
	ammoniterepl_cmd_src := exec.Command("curl", "-L", ammoniterepl_src_url, "-o", "source.tar.gz")
	err = ammoniterepl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ammoniterepl_cmd_direct := exec.Command("./binary")
	err = ammoniterepl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
