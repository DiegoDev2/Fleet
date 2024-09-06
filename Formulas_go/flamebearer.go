package main

// Flamebearer - Blazing fast flame graph tool for V8 and Node
// Homepage: https://github.com/mapbox/flamebearer

import (
	"fmt"
	
	"os/exec"
)

func installFlamebearer() {
	// Método 1: Descargar y extraer .tar.gz
	flamebearer_tar_url := "https://registry.npmjs.org/flamebearer/-/flamebearer-1.1.3.tgz"
	flamebearer_cmd_tar := exec.Command("curl", "-L", flamebearer_tar_url, "-o", "package.tar.gz")
	err := flamebearer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flamebearer_zip_url := "https://registry.npmjs.org/flamebearer/-/flamebearer-1.1.3.tgz"
	flamebearer_cmd_zip := exec.Command("curl", "-L", flamebearer_zip_url, "-o", "package.zip")
	err = flamebearer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flamebearer_bin_url := "https://registry.npmjs.org/flamebearer/-/flamebearer-1.1.3.tgz"
	flamebearer_cmd_bin := exec.Command("curl", "-L", flamebearer_bin_url, "-o", "binary.bin")
	err = flamebearer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flamebearer_src_url := "https://registry.npmjs.org/flamebearer/-/flamebearer-1.1.3.tgz"
	flamebearer_cmd_src := exec.Command("curl", "-L", flamebearer_src_url, "-o", "source.tar.gz")
	err = flamebearer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flamebearer_cmd_direct := exec.Command("./binary")
	err = flamebearer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
