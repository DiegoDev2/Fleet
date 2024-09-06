package main

// Pug - Drive terraform at terminal velocity
// Homepage: https://github.com/leg100/pug

import (
	"fmt"
	
	"os/exec"
)

func installPug() {
	// Método 1: Descargar y extraer .tar.gz
	pug_tar_url := "https://github.com/leg100/pug/archive/refs/tags/v0.5.1.tar.gz"
	pug_cmd_tar := exec.Command("curl", "-L", pug_tar_url, "-o", "package.tar.gz")
	err := pug_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pug_zip_url := "https://github.com/leg100/pug/archive/refs/tags/v0.5.1.zip"
	pug_cmd_zip := exec.Command("curl", "-L", pug_zip_url, "-o", "package.zip")
	err = pug_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pug_bin_url := "https://github.com/leg100/pug/archive/refs/tags/v0.5.1.bin"
	pug_cmd_bin := exec.Command("curl", "-L", pug_bin_url, "-o", "binary.bin")
	err = pug_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pug_src_url := "https://github.com/leg100/pug/archive/refs/tags/v0.5.1.src.tar.gz"
	pug_cmd_src := exec.Command("curl", "-L", pug_src_url, "-o", "source.tar.gz")
	err = pug_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pug_cmd_direct := exec.Command("./binary")
	err = pug_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
