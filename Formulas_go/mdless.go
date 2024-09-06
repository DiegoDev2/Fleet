package main

// Mdless - Provides a formatted and highlighted view of Markdown files in Terminal
// Homepage: https://github.com/ttscoff/mdless

import (
	"fmt"
	
	"os/exec"
)

func installMdless() {
	// Método 1: Descargar y extraer .tar.gz
	mdless_tar_url := "https://github.com/ttscoff/mdless/archive/refs/tags/2.1.44.tar.gz"
	mdless_cmd_tar := exec.Command("curl", "-L", mdless_tar_url, "-o", "package.tar.gz")
	err := mdless_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdless_zip_url := "https://github.com/ttscoff/mdless/archive/refs/tags/2.1.44.zip"
	mdless_cmd_zip := exec.Command("curl", "-L", mdless_zip_url, "-o", "package.zip")
	err = mdless_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdless_bin_url := "https://github.com/ttscoff/mdless/archive/refs/tags/2.1.44.bin"
	mdless_cmd_bin := exec.Command("curl", "-L", mdless_bin_url, "-o", "binary.bin")
	err = mdless_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdless_src_url := "https://github.com/ttscoff/mdless/archive/refs/tags/2.1.44.src.tar.gz"
	mdless_cmd_src := exec.Command("curl", "-L", mdless_src_url, "-o", "source.tar.gz")
	err = mdless_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdless_cmd_direct := exec.Command("./binary")
	err = mdless_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
}
