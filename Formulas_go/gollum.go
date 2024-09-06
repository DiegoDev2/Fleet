package main

// Gollum - Go n:m message multiplexer
// Homepage: https://gollum.readthedocs.io/en/latest/

import (
	"fmt"
	
	"os/exec"
)

func installGollum() {
	// Método 1: Descargar y extraer .tar.gz
	gollum_tar_url := "https://github.com/trivago/gollum/archive/refs/tags/0.6.0.tar.gz"
	gollum_cmd_tar := exec.Command("curl", "-L", gollum_tar_url, "-o", "package.tar.gz")
	err := gollum_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gollum_zip_url := "https://github.com/trivago/gollum/archive/refs/tags/0.6.0.zip"
	gollum_cmd_zip := exec.Command("curl", "-L", gollum_zip_url, "-o", "package.zip")
	err = gollum_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gollum_bin_url := "https://github.com/trivago/gollum/archive/refs/tags/0.6.0.bin"
	gollum_cmd_bin := exec.Command("curl", "-L", gollum_bin_url, "-o", "binary.bin")
	err = gollum_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gollum_src_url := "https://github.com/trivago/gollum/archive/refs/tags/0.6.0.src.tar.gz"
	gollum_cmd_src := exec.Command("curl", "-L", gollum_src_url, "-o", "source.tar.gz")
	err = gollum_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gollum_cmd_direct := exec.Command("./binary")
	err = gollum_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
