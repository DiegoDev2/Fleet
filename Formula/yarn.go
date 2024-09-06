package main

// Yarn - JavaScript package manager
// Homepage: https://yarnpkg.com/

import (
	"fmt"
	
	"os/exec"
)

func installYarn() {
	// Método 1: Descargar y extraer .tar.gz
	yarn_tar_url := "https://yarnpkg.com/downloads/1.22.22/yarn-v1.22.22.tar.gz"
	yarn_cmd_tar := exec.Command("curl", "-L", yarn_tar_url, "-o", "package.tar.gz")
	err := yarn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	yarn_zip_url := "https://yarnpkg.com/downloads/1.22.22/yarn-v1.22.22.zip"
	yarn_cmd_zip := exec.Command("curl", "-L", yarn_zip_url, "-o", "package.zip")
	err = yarn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	yarn_bin_url := "https://yarnpkg.com/downloads/1.22.22/yarn-v1.22.22.bin"
	yarn_cmd_bin := exec.Command("curl", "-L", yarn_bin_url, "-o", "binary.bin")
	err = yarn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	yarn_src_url := "https://yarnpkg.com/downloads/1.22.22/yarn-v1.22.22.src.tar.gz"
	yarn_cmd_src := exec.Command("curl", "-L", yarn_src_url, "-o", "source.tar.gz")
	err = yarn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	yarn_cmd_direct := exec.Command("./binary")
	err = yarn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
