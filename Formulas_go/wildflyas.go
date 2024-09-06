package main

// WildflyAs - Managed application runtime for building applications
// Homepage: https://www.wildfly.org/

import (
	"fmt"
	
	"os/exec"
)

func installWildflyAs() {
	// Método 1: Descargar y extraer .tar.gz
	wildflyas_tar_url := "https://github.com/wildfly/wildfly/releases/download/32.0.0.Final/wildfly-32.0.0.Final.tar.gz"
	wildflyas_cmd_tar := exec.Command("curl", "-L", wildflyas_tar_url, "-o", "package.tar.gz")
	err := wildflyas_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wildflyas_zip_url := "https://github.com/wildfly/wildfly/releases/download/32.0.0.Final/wildfly-32.0.0.Final.zip"
	wildflyas_cmd_zip := exec.Command("curl", "-L", wildflyas_zip_url, "-o", "package.zip")
	err = wildflyas_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wildflyas_bin_url := "https://github.com/wildfly/wildfly/releases/download/32.0.0.Final/wildfly-32.0.0.Final.bin"
	wildflyas_cmd_bin := exec.Command("curl", "-L", wildflyas_bin_url, "-o", "binary.bin")
	err = wildflyas_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wildflyas_src_url := "https://github.com/wildfly/wildfly/releases/download/32.0.0.Final/wildfly-32.0.0.Final.src.tar.gz"
	wildflyas_cmd_src := exec.Command("curl", "-L", wildflyas_src_url, "-o", "source.tar.gz")
	err = wildflyas_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wildflyas_cmd_direct := exec.Command("./binary")
	err = wildflyas_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
