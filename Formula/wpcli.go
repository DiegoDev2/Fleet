package main

// WpCli - Command-line interface for WordPress
// Homepage: https://wp-cli.org/

import (
	"fmt"
	
	"os/exec"
)

func installWpCli() {
	// Método 1: Descargar y extraer .tar.gz
	wpcli_tar_url := "https://github.com/wp-cli/wp-cli/releases/download/v2.11.0/wp-cli-2.11.0.phar"
	wpcli_cmd_tar := exec.Command("curl", "-L", wpcli_tar_url, "-o", "package.tar.gz")
	err := wpcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wpcli_zip_url := "https://github.com/wp-cli/wp-cli/releases/download/v2.11.0/wp-cli-2.11.0.phar"
	wpcli_cmd_zip := exec.Command("curl", "-L", wpcli_zip_url, "-o", "package.zip")
	err = wpcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wpcli_bin_url := "https://github.com/wp-cli/wp-cli/releases/download/v2.11.0/wp-cli-2.11.0.phar"
	wpcli_cmd_bin := exec.Command("curl", "-L", wpcli_bin_url, "-o", "binary.bin")
	err = wpcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wpcli_src_url := "https://github.com/wp-cli/wp-cli/releases/download/v2.11.0/wp-cli-2.11.0.phar"
	wpcli_cmd_src := exec.Command("curl", "-L", wpcli_src_url, "-o", "source.tar.gz")
	err = wpcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wpcli_cmd_direct := exec.Command("./binary")
	err = wpcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}
