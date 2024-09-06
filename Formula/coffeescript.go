package main

// Coffeescript - Unfancy JavaScript
// Homepage: https://coffeescript.org/

import (
	"fmt"
	
	"os/exec"
)

func installCoffeescript() {
	// Método 1: Descargar y extraer .tar.gz
	coffeescript_tar_url := "https://registry.npmjs.org/coffeescript/-/coffeescript-2.7.0.tgz"
	coffeescript_cmd_tar := exec.Command("curl", "-L", coffeescript_tar_url, "-o", "package.tar.gz")
	err := coffeescript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coffeescript_zip_url := "https://registry.npmjs.org/coffeescript/-/coffeescript-2.7.0.tgz"
	coffeescript_cmd_zip := exec.Command("curl", "-L", coffeescript_zip_url, "-o", "package.zip")
	err = coffeescript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coffeescript_bin_url := "https://registry.npmjs.org/coffeescript/-/coffeescript-2.7.0.tgz"
	coffeescript_cmd_bin := exec.Command("curl", "-L", coffeescript_bin_url, "-o", "binary.bin")
	err = coffeescript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coffeescript_src_url := "https://registry.npmjs.org/coffeescript/-/coffeescript-2.7.0.tgz"
	coffeescript_cmd_src := exec.Command("curl", "-L", coffeescript_src_url, "-o", "source.tar.gz")
	err = coffeescript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coffeescript_cmd_direct := exec.Command("./binary")
	err = coffeescript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
