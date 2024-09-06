package main

// PhpCsFixer - Tool to automatically fix PHP coding standards issues
// Homepage: https://cs.symfony.com/

import (
	"fmt"
	
	"os/exec"
)

func installPhpCsFixer() {
	// Método 1: Descargar y extraer .tar.gz
	phpcsfixer_tar_url := "https://github.com/PHP-CS-Fixer/PHP-CS-Fixer/releases/download/v3.64.0/php-cs-fixer.phar"
	phpcsfixer_cmd_tar := exec.Command("curl", "-L", phpcsfixer_tar_url, "-o", "package.tar.gz")
	err := phpcsfixer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phpcsfixer_zip_url := "https://github.com/PHP-CS-Fixer/PHP-CS-Fixer/releases/download/v3.64.0/php-cs-fixer.phar"
	phpcsfixer_cmd_zip := exec.Command("curl", "-L", phpcsfixer_zip_url, "-o", "package.zip")
	err = phpcsfixer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phpcsfixer_bin_url := "https://github.com/PHP-CS-Fixer/PHP-CS-Fixer/releases/download/v3.64.0/php-cs-fixer.phar"
	phpcsfixer_cmd_bin := exec.Command("curl", "-L", phpcsfixer_bin_url, "-o", "binary.bin")
	err = phpcsfixer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phpcsfixer_src_url := "https://github.com/PHP-CS-Fixer/PHP-CS-Fixer/releases/download/v3.64.0/php-cs-fixer.phar"
	phpcsfixer_cmd_src := exec.Command("curl", "-L", phpcsfixer_src_url, "-o", "source.tar.gz")
	err = phpcsfixer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phpcsfixer_cmd_direct := exec.Command("./binary")
	err = phpcsfixer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
exec.Command("latte", "install", "php")
}
