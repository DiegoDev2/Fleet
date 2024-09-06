package main

// Autocode - Code automation for every language, library and framework
// Homepage: https://autocode.readme.io/

import (
	"fmt"
	
	"os/exec"
)

func installAutocode() {
	// Método 1: Descargar y extraer .tar.gz
	autocode_tar_url := "https://registry.npmjs.org/autocode/-/autocode-1.3.1.tgz"
	autocode_cmd_tar := exec.Command("curl", "-L", autocode_tar_url, "-o", "package.tar.gz")
	err := autocode_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autocode_zip_url := "https://registry.npmjs.org/autocode/-/autocode-1.3.1.tgz"
	autocode_cmd_zip := exec.Command("curl", "-L", autocode_zip_url, "-o", "package.zip")
	err = autocode_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autocode_bin_url := "https://registry.npmjs.org/autocode/-/autocode-1.3.1.tgz"
	autocode_cmd_bin := exec.Command("curl", "-L", autocode_bin_url, "-o", "binary.bin")
	err = autocode_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autocode_src_url := "https://registry.npmjs.org/autocode/-/autocode-1.3.1.tgz"
	autocode_cmd_src := exec.Command("curl", "-L", autocode_src_url, "-o", "source.tar.gz")
	err = autocode_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autocode_cmd_direct := exec.Command("./binary")
	err = autocode_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
