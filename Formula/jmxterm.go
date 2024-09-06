package main

// Jmxterm - Open source, command-line based interactive JMX client
// Homepage: https://docs.cyclopsgroup.org/jmxterm

import (
	"fmt"
	
	"os/exec"
)

func installJmxterm() {
	// Método 1: Descargar y extraer .tar.gz
	jmxterm_tar_url := "https://github.com/jiaqi/jmxterm/releases/download/v1.0.4/jmxterm-1.0.4-uber.jar"
	jmxterm_cmd_tar := exec.Command("curl", "-L", jmxterm_tar_url, "-o", "package.tar.gz")
	err := jmxterm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jmxterm_zip_url := "https://github.com/jiaqi/jmxterm/releases/download/v1.0.4/jmxterm-1.0.4-uber.jar"
	jmxterm_cmd_zip := exec.Command("curl", "-L", jmxterm_zip_url, "-o", "package.zip")
	err = jmxterm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jmxterm_bin_url := "https://github.com/jiaqi/jmxterm/releases/download/v1.0.4/jmxterm-1.0.4-uber.jar"
	jmxterm_cmd_bin := exec.Command("curl", "-L", jmxterm_bin_url, "-o", "binary.bin")
	err = jmxterm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jmxterm_src_url := "https://github.com/jiaqi/jmxterm/releases/download/v1.0.4/jmxterm-1.0.4-uber.jar"
	jmxterm_cmd_src := exec.Command("curl", "-L", jmxterm_src_url, "-o", "source.tar.gz")
	err = jmxterm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jmxterm_cmd_direct := exec.Command("./binary")
	err = jmxterm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
