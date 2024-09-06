package main

// Gistit - Command-line utility for creating Gists
// Homepage: https://github.com/jrbasso/gistit

import (
	"fmt"
	
	"os/exec"
)

func installGistit() {
	// Método 1: Descargar y extraer .tar.gz
	gistit_tar_url := "https://github.com/jrbasso/gistit/archive/refs/tags/v0.1.4.tar.gz"
	gistit_cmd_tar := exec.Command("curl", "-L", gistit_tar_url, "-o", "package.tar.gz")
	err := gistit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gistit_zip_url := "https://github.com/jrbasso/gistit/archive/refs/tags/v0.1.4.zip"
	gistit_cmd_zip := exec.Command("curl", "-L", gistit_zip_url, "-o", "package.zip")
	err = gistit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gistit_bin_url := "https://github.com/jrbasso/gistit/archive/refs/tags/v0.1.4.bin"
	gistit_cmd_bin := exec.Command("curl", "-L", gistit_bin_url, "-o", "binary.bin")
	err = gistit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gistit_src_url := "https://github.com/jrbasso/gistit/archive/refs/tags/v0.1.4.src.tar.gz"
	gistit_cmd_src := exec.Command("curl", "-L", gistit_src_url, "-o", "source.tar.gz")
	err = gistit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gistit_cmd_direct := exec.Command("./binary")
	err = gistit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
}
