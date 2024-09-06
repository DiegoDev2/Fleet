package main

// Psalm - PHP Static Analysis Tool
// Homepage: https://psalm.dev

import (
	"fmt"
	
	"os/exec"
)

func installPsalm() {
	// Método 1: Descargar y extraer .tar.gz
	psalm_tar_url := "https://github.com/vimeo/psalm/releases/download/5.25.0/psalm.phar"
	psalm_cmd_tar := exec.Command("curl", "-L", psalm_tar_url, "-o", "package.tar.gz")
	err := psalm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	psalm_zip_url := "https://github.com/vimeo/psalm/releases/download/5.25.0/psalm.phar"
	psalm_cmd_zip := exec.Command("curl", "-L", psalm_zip_url, "-o", "package.zip")
	err = psalm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	psalm_bin_url := "https://github.com/vimeo/psalm/releases/download/5.25.0/psalm.phar"
	psalm_cmd_bin := exec.Command("curl", "-L", psalm_bin_url, "-o", "binary.bin")
	err = psalm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	psalm_src_url := "https://github.com/vimeo/psalm/releases/download/5.25.0/psalm.phar"
	psalm_cmd_src := exec.Command("curl", "-L", psalm_src_url, "-o", "source.tar.gz")
	err = psalm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	psalm_cmd_direct := exec.Command("./binary")
	err = psalm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: composer")
	exec.Command("latte", "install", "composer").Run()
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}
