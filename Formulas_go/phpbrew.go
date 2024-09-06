package main

// Phpbrew - Brew & manage PHP versions in pure PHP at HOME
// Homepage: https://phpbrew.github.io/phpbrew

import (
	"fmt"
	
	"os/exec"
)

func installPhpbrew() {
	// Método 1: Descargar y extraer .tar.gz
	phpbrew_tar_url := "https://github.com/phpbrew/phpbrew/releases/download/2.2.0/phpbrew.phar"
	phpbrew_cmd_tar := exec.Command("curl", "-L", phpbrew_tar_url, "-o", "package.tar.gz")
	err := phpbrew_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phpbrew_zip_url := "https://github.com/phpbrew/phpbrew/releases/download/2.2.0/phpbrew.phar"
	phpbrew_cmd_zip := exec.Command("curl", "-L", phpbrew_zip_url, "-o", "package.zip")
	err = phpbrew_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phpbrew_bin_url := "https://github.com/phpbrew/phpbrew/releases/download/2.2.0/phpbrew.phar"
	phpbrew_cmd_bin := exec.Command("curl", "-L", phpbrew_bin_url, "-o", "binary.bin")
	err = phpbrew_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phpbrew_src_url := "https://github.com/phpbrew/phpbrew/releases/download/2.2.0/phpbrew.phar"
	phpbrew_cmd_src := exec.Command("curl", "-L", phpbrew_src_url, "-o", "source.tar.gz")
	err = phpbrew_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phpbrew_cmd_direct := exec.Command("./binary")
	err = phpbrew_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
exec.Command("latte", "install", "php")
}
