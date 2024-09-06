package main

// Gostatic - Fast static site generator
// Homepage: https://github.com/piranha/gostatic

import (
	"fmt"
	
	"os/exec"
)

func installGostatic() {
	// Método 1: Descargar y extraer .tar.gz
	gostatic_tar_url := "https://github.com/piranha/gostatic/archive/refs/tags/2.36.tar.gz"
	gostatic_cmd_tar := exec.Command("curl", "-L", gostatic_tar_url, "-o", "package.tar.gz")
	err := gostatic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gostatic_zip_url := "https://github.com/piranha/gostatic/archive/refs/tags/2.36.zip"
	gostatic_cmd_zip := exec.Command("curl", "-L", gostatic_zip_url, "-o", "package.zip")
	err = gostatic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gostatic_bin_url := "https://github.com/piranha/gostatic/archive/refs/tags/2.36.bin"
	gostatic_cmd_bin := exec.Command("curl", "-L", gostatic_bin_url, "-o", "binary.bin")
	err = gostatic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gostatic_src_url := "https://github.com/piranha/gostatic/archive/refs/tags/2.36.src.tar.gz"
	gostatic_cmd_src := exec.Command("curl", "-L", gostatic_src_url, "-o", "source.tar.gz")
	err = gostatic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gostatic_cmd_direct := exec.Command("./binary")
	err = gostatic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
