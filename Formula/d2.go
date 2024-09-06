package main

// D2 - Modern diagram scripting language that turns text to diagrams
// Homepage: https://d2lang.com/

import (
	"fmt"
	
	"os/exec"
)

func installD2() {
	// Método 1: Descargar y extraer .tar.gz
	d2_tar_url := "https://github.com/terrastruct/d2/archive/refs/tags/v0.6.6.tar.gz"
	d2_cmd_tar := exec.Command("curl", "-L", d2_tar_url, "-o", "package.tar.gz")
	err := d2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	d2_zip_url := "https://github.com/terrastruct/d2/archive/refs/tags/v0.6.6.zip"
	d2_cmd_zip := exec.Command("curl", "-L", d2_zip_url, "-o", "package.zip")
	err = d2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	d2_bin_url := "https://github.com/terrastruct/d2/archive/refs/tags/v0.6.6.bin"
	d2_cmd_bin := exec.Command("curl", "-L", d2_bin_url, "-o", "binary.bin")
	err = d2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	d2_src_url := "https://github.com/terrastruct/d2/archive/refs/tags/v0.6.6.src.tar.gz"
	d2_cmd_src := exec.Command("curl", "-L", d2_src_url, "-o", "source.tar.gz")
	err = d2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	d2_cmd_direct := exec.Command("./binary")
	err = d2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
