package main

// Lerna - Tool for managing JavaScript projects with multiple packages
// Homepage: https://lerna.js.org

import (
	"fmt"
	
	"os/exec"
)

func installLerna() {
	// Método 1: Descargar y extraer .tar.gz
	lerna_tar_url := "https://registry.npmjs.org/lerna/-/lerna-8.1.8.tgz"
	lerna_cmd_tar := exec.Command("curl", "-L", lerna_tar_url, "-o", "package.tar.gz")
	err := lerna_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lerna_zip_url := "https://registry.npmjs.org/lerna/-/lerna-8.1.8.tgz"
	lerna_cmd_zip := exec.Command("curl", "-L", lerna_zip_url, "-o", "package.zip")
	err = lerna_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lerna_bin_url := "https://registry.npmjs.org/lerna/-/lerna-8.1.8.tgz"
	lerna_cmd_bin := exec.Command("curl", "-L", lerna_bin_url, "-o", "binary.bin")
	err = lerna_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lerna_src_url := "https://registry.npmjs.org/lerna/-/lerna-8.1.8.tgz"
	lerna_cmd_src := exec.Command("curl", "-L", lerna_src_url, "-o", "source.tar.gz")
	err = lerna_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lerna_cmd_direct := exec.Command("./binary")
	err = lerna_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
