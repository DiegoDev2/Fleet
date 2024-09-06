package main

// Cortexso - Drop-in, local AI alternative to the OpenAI stack
// Homepage: https://jan.ai/cortex

import (
	"fmt"
	
	"os/exec"
)

func installCortexso() {
	// Método 1: Descargar y extraer .tar.gz
	cortexso_tar_url := "https://registry.npmjs.org/cortexso/-/cortexso-0.1.1.tgz"
	cortexso_cmd_tar := exec.Command("curl", "-L", cortexso_tar_url, "-o", "package.tar.gz")
	err := cortexso_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cortexso_zip_url := "https://registry.npmjs.org/cortexso/-/cortexso-0.1.1.tgz"
	cortexso_cmd_zip := exec.Command("curl", "-L", cortexso_zip_url, "-o", "package.zip")
	err = cortexso_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cortexso_bin_url := "https://registry.npmjs.org/cortexso/-/cortexso-0.1.1.tgz"
	cortexso_cmd_bin := exec.Command("curl", "-L", cortexso_bin_url, "-o", "binary.bin")
	err = cortexso_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cortexso_src_url := "https://registry.npmjs.org/cortexso/-/cortexso-0.1.1.tgz"
	cortexso_cmd_src := exec.Command("curl", "-L", cortexso_src_url, "-o", "source.tar.gz")
	err = cortexso_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cortexso_cmd_direct := exec.Command("./binary")
	err = cortexso_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
}
