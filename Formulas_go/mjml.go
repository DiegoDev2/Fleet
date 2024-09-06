package main

// Mjml - JavaScript framework that makes responsive-email easy
// Homepage: https://mjml.io

import (
	"fmt"
	
	"os/exec"
)

func installMjml() {
	// Método 1: Descargar y extraer .tar.gz
	mjml_tar_url := "https://registry.npmjs.org/mjml/-/mjml-4.15.3.tgz"
	mjml_cmd_tar := exec.Command("curl", "-L", mjml_tar_url, "-o", "package.tar.gz")
	err := mjml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mjml_zip_url := "https://registry.npmjs.org/mjml/-/mjml-4.15.3.tgz"
	mjml_cmd_zip := exec.Command("curl", "-L", mjml_zip_url, "-o", "package.zip")
	err = mjml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mjml_bin_url := "https://registry.npmjs.org/mjml/-/mjml-4.15.3.tgz"
	mjml_cmd_bin := exec.Command("curl", "-L", mjml_bin_url, "-o", "binary.bin")
	err = mjml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mjml_src_url := "https://registry.npmjs.org/mjml/-/mjml-4.15.3.tgz"
	mjml_cmd_src := exec.Command("curl", "-L", mjml_src_url, "-o", "source.tar.gz")
	err = mjml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mjml_cmd_direct := exec.Command("./binary")
	err = mjml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
