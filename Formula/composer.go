package main

// Composer - Dependency Manager for PHP
// Homepage: https://getcomposer.org/

import (
	"fmt"
	
	"os/exec"
)

func installComposer() {
	// Método 1: Descargar y extraer .tar.gz
	composer_tar_url := "https://getcomposer.org/download/2.7.9/composer.phar"
	composer_cmd_tar := exec.Command("curl", "-L", composer_tar_url, "-o", "package.tar.gz")
	err := composer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	composer_zip_url := "https://getcomposer.org/download/2.7.9/composer.phar"
	composer_cmd_zip := exec.Command("curl", "-L", composer_zip_url, "-o", "package.zip")
	err = composer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	composer_bin_url := "https://getcomposer.org/download/2.7.9/composer.phar"
	composer_cmd_bin := exec.Command("curl", "-L", composer_bin_url, "-o", "binary.bin")
	err = composer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	composer_src_url := "https://getcomposer.org/download/2.7.9/composer.phar"
	composer_cmd_src := exec.Command("curl", "-L", composer_src_url, "-o", "source.tar.gz")
	err = composer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	composer_cmd_direct := exec.Command("./binary")
	err = composer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}
