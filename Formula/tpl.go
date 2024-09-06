package main

// Tpl - Store and retrieve binary data in C
// Homepage: https://troydhanson.github.io/tpl/

import (
	"fmt"
	
	"os/exec"
)

func installTpl() {
	// Método 1: Descargar y extraer .tar.gz
	tpl_tar_url := "https://github.com/troydhanson/tpl/archive/refs/tags/v1.6.1.tar.gz"
	tpl_cmd_tar := exec.Command("curl", "-L", tpl_tar_url, "-o", "package.tar.gz")
	err := tpl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tpl_zip_url := "https://github.com/troydhanson/tpl/archive/refs/tags/v1.6.1.zip"
	tpl_cmd_zip := exec.Command("curl", "-L", tpl_zip_url, "-o", "package.zip")
	err = tpl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tpl_bin_url := "https://github.com/troydhanson/tpl/archive/refs/tags/v1.6.1.bin"
	tpl_cmd_bin := exec.Command("curl", "-L", tpl_bin_url, "-o", "binary.bin")
	err = tpl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tpl_src_url := "https://github.com/troydhanson/tpl/archive/refs/tags/v1.6.1.src.tar.gz"
	tpl_cmd_src := exec.Command("curl", "-L", tpl_src_url, "-o", "source.tar.gz")
	err = tpl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tpl_cmd_direct := exec.Command("./binary")
	err = tpl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
