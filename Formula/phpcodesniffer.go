package main

// PhpCodeSniffer - Check coding standards in PHP, JavaScript and CSS
// Homepage: https://github.com/PHPCSStandards/PHP_CodeSniffer

import (
	"fmt"
	
	"os/exec"
)

func installPhpCodeSniffer() {
	// Método 1: Descargar y extraer .tar.gz
	phpcodesniffer_tar_url := "https://github.com/PHPCSStandards/PHP_CodeSniffer/releases/download/3.10.2/phpcs.phar"
	phpcodesniffer_cmd_tar := exec.Command("curl", "-L", phpcodesniffer_tar_url, "-o", "package.tar.gz")
	err := phpcodesniffer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phpcodesniffer_zip_url := "https://github.com/PHPCSStandards/PHP_CodeSniffer/releases/download/3.10.2/phpcs.phar"
	phpcodesniffer_cmd_zip := exec.Command("curl", "-L", phpcodesniffer_zip_url, "-o", "package.zip")
	err = phpcodesniffer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phpcodesniffer_bin_url := "https://github.com/PHPCSStandards/PHP_CodeSniffer/releases/download/3.10.2/phpcs.phar"
	phpcodesniffer_cmd_bin := exec.Command("curl", "-L", phpcodesniffer_bin_url, "-o", "binary.bin")
	err = phpcodesniffer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phpcodesniffer_src_url := "https://github.com/PHPCSStandards/PHP_CodeSniffer/releases/download/3.10.2/phpcs.phar"
	phpcodesniffer_cmd_src := exec.Command("curl", "-L", phpcodesniffer_src_url, "-o", "source.tar.gz")
	err = phpcodesniffer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phpcodesniffer_cmd_direct := exec.Command("./binary")
	err = phpcodesniffer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}
