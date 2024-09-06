package main

// Byteman - Java bytecode manipulation tool for testing, monitoring and tracing
// Homepage: https://byteman.jboss.org/

import (
	"fmt"
	
	"os/exec"
)

func installByteman() {
	// Método 1: Descargar y extraer .tar.gz
	byteman_tar_url := "https://downloads.jboss.org/byteman/4.0.23/byteman-download-4.0.23-bin.zip"
	byteman_cmd_tar := exec.Command("curl", "-L", byteman_tar_url, "-o", "package.tar.gz")
	err := byteman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	byteman_zip_url := "https://downloads.jboss.org/byteman/4.0.23/byteman-download-4.0.23-bin.zip"
	byteman_cmd_zip := exec.Command("curl", "-L", byteman_zip_url, "-o", "package.zip")
	err = byteman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	byteman_bin_url := "https://downloads.jboss.org/byteman/4.0.23/byteman-download-4.0.23-bin.zip"
	byteman_cmd_bin := exec.Command("curl", "-L", byteman_bin_url, "-o", "binary.bin")
	err = byteman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	byteman_src_url := "https://downloads.jboss.org/byteman/4.0.23/byteman-download-4.0.23-bin.zip"
	byteman_cmd_src := exec.Command("curl", "-L", byteman_src_url, "-o", "source.tar.gz")
	err = byteman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	byteman_cmd_direct := exec.Command("./binary")
	err = byteman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}
