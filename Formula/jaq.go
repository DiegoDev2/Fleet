package main

// Jaq - JQ clone focussed on correctness, speed, and simplicity
// Homepage: https://github.com/01mf02/jaq

import (
	"fmt"
	
	"os/exec"
)

func installJaq() {
	// Método 1: Descargar y extraer .tar.gz
	jaq_tar_url := "https://github.com/01mf02/jaq/archive/refs/tags/v1.6.0.tar.gz"
	jaq_cmd_tar := exec.Command("curl", "-L", jaq_tar_url, "-o", "package.tar.gz")
	err := jaq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jaq_zip_url := "https://github.com/01mf02/jaq/archive/refs/tags/v1.6.0.zip"
	jaq_cmd_zip := exec.Command("curl", "-L", jaq_zip_url, "-o", "package.zip")
	err = jaq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jaq_bin_url := "https://github.com/01mf02/jaq/archive/refs/tags/v1.6.0.bin"
	jaq_cmd_bin := exec.Command("curl", "-L", jaq_bin_url, "-o", "binary.bin")
	err = jaq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jaq_src_url := "https://github.com/01mf02/jaq/archive/refs/tags/v1.6.0.src.tar.gz"
	jaq_cmd_src := exec.Command("curl", "-L", jaq_src_url, "-o", "source.tar.gz")
	err = jaq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jaq_cmd_direct := exec.Command("./binary")
	err = jaq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
