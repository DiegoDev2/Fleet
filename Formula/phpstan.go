package main

// Phpstan - PHP Static Analysis Tool
// Homepage: https://github.com/phpstan/phpstan

import (
	"fmt"
	
	"os/exec"
)

func installPhpstan() {
	// Método 1: Descargar y extraer .tar.gz
	phpstan_tar_url := "https://github.com/phpstan/phpstan/releases/download/1.12.2/phpstan.phar"
	phpstan_cmd_tar := exec.Command("curl", "-L", phpstan_tar_url, "-o", "package.tar.gz")
	err := phpstan_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phpstan_zip_url := "https://github.com/phpstan/phpstan/releases/download/1.12.2/phpstan.phar"
	phpstan_cmd_zip := exec.Command("curl", "-L", phpstan_zip_url, "-o", "package.zip")
	err = phpstan_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phpstan_bin_url := "https://github.com/phpstan/phpstan/releases/download/1.12.2/phpstan.phar"
	phpstan_cmd_bin := exec.Command("curl", "-L", phpstan_bin_url, "-o", "binary.bin")
	err = phpstan_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phpstan_src_url := "https://github.com/phpstan/phpstan/releases/download/1.12.2/phpstan.phar"
	phpstan_cmd_src := exec.Command("curl", "-L", phpstan_src_url, "-o", "source.tar.gz")
	err = phpstan_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phpstan_cmd_direct := exec.Command("./binary")
	err = phpstan_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}
