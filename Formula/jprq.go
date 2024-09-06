package main

// Jprq - Join Public Router, Quickly
// Homepage: https://jprq.io/

import (
	"fmt"
	
	"os/exec"
)

func installJprq() {
	// Método 1: Descargar y extraer .tar.gz
	jprq_tar_url := "https://github.com/azimjohn/jprq/archive/refs/tags/2.3.tar.gz"
	jprq_cmd_tar := exec.Command("curl", "-L", jprq_tar_url, "-o", "package.tar.gz")
	err := jprq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jprq_zip_url := "https://github.com/azimjohn/jprq/archive/refs/tags/2.3.zip"
	jprq_cmd_zip := exec.Command("curl", "-L", jprq_zip_url, "-o", "package.zip")
	err = jprq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jprq_bin_url := "https://github.com/azimjohn/jprq/archive/refs/tags/2.3.bin"
	jprq_cmd_bin := exec.Command("curl", "-L", jprq_bin_url, "-o", "binary.bin")
	err = jprq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jprq_src_url := "https://github.com/azimjohn/jprq/archive/refs/tags/2.3.src.tar.gz"
	jprq_cmd_src := exec.Command("curl", "-L", jprq_src_url, "-o", "source.tar.gz")
	err = jprq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jprq_cmd_direct := exec.Command("./binary")
	err = jprq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
