package main

// BrewPhpSwitcher - Switch Apache / Valet / CLI configs between PHP versions
// Homepage: https://github.com/philcook/brew-php-switcher

import (
	"fmt"
	
	"os/exec"
)

func installBrewPhpSwitcher() {
	// Método 1: Descargar y extraer .tar.gz
	brewphpswitcher_tar_url := "https://github.com/philcook/brew-php-switcher/archive/refs/tags/v2.5.tar.gz"
	brewphpswitcher_cmd_tar := exec.Command("curl", "-L", brewphpswitcher_tar_url, "-o", "package.tar.gz")
	err := brewphpswitcher_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brewphpswitcher_zip_url := "https://github.com/philcook/brew-php-switcher/archive/refs/tags/v2.5.zip"
	brewphpswitcher_cmd_zip := exec.Command("curl", "-L", brewphpswitcher_zip_url, "-o", "package.zip")
	err = brewphpswitcher_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brewphpswitcher_bin_url := "https://github.com/philcook/brew-php-switcher/archive/refs/tags/v2.5.bin"
	brewphpswitcher_cmd_bin := exec.Command("curl", "-L", brewphpswitcher_bin_url, "-o", "binary.bin")
	err = brewphpswitcher_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brewphpswitcher_src_url := "https://github.com/philcook/brew-php-switcher/archive/refs/tags/v2.5.src.tar.gz"
	brewphpswitcher_cmd_src := exec.Command("curl", "-L", brewphpswitcher_src_url, "-o", "source.tar.gz")
	err = brewphpswitcher_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brewphpswitcher_cmd_direct := exec.Command("./binary")
	err = brewphpswitcher_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
	exec.Command("latte", "install", "php").Run()
}
