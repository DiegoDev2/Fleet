package main

// Phpunit - Programmer-oriented testing framework for PHP
// Homepage: https://phpunit.de

import (
	"fmt"
	
	"os/exec"
)

func installPhpunit() {
	// Método 1: Descargar y extraer .tar.gz
	phpunit_tar_url := "https://phar.phpunit.de/phpunit-11.3.3.phar"
	phpunit_cmd_tar := exec.Command("curl", "-L", phpunit_tar_url, "-o", "package.tar.gz")
	err := phpunit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phpunit_zip_url := "https://phar.phpunit.de/phpunit-11.3.3.phar"
	phpunit_cmd_zip := exec.Command("curl", "-L", phpunit_zip_url, "-o", "package.zip")
	err = phpunit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phpunit_bin_url := "https://phar.phpunit.de/phpunit-11.3.3.phar"
	phpunit_cmd_bin := exec.Command("curl", "-L", phpunit_bin_url, "-o", "binary.bin")
	err = phpunit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phpunit_src_url := "https://phar.phpunit.de/phpunit-11.3.3.phar"
	phpunit_cmd_src := exec.Command("curl", "-L", phpunit_src_url, "-o", "source.tar.gz")
	err = phpunit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phpunit_cmd_direct := exec.Command("./binary")
	err = phpunit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
exec.Command("latte", "install", "php")
}
