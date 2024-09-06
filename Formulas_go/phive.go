package main

// Phive - Phar Installation and Verification Environment (PHIVE)
// Homepage: https://phar.io

import (
	"fmt"
	
	"os/exec"
)

func installPhive() {
	// Método 1: Descargar y extraer .tar.gz
	phive_tar_url := "https://github.com/phar-io/phive/releases/download/0.15.3/phive-0.15.3.phar"
	phive_cmd_tar := exec.Command("curl", "-L", phive_tar_url, "-o", "package.tar.gz")
	err := phive_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phive_zip_url := "https://github.com/phar-io/phive/releases/download/0.15.3/phive-0.15.3.phar"
	phive_cmd_zip := exec.Command("curl", "-L", phive_zip_url, "-o", "package.zip")
	err = phive_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phive_bin_url := "https://github.com/phar-io/phive/releases/download/0.15.3/phive-0.15.3.phar"
	phive_cmd_bin := exec.Command("curl", "-L", phive_bin_url, "-o", "binary.bin")
	err = phive_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phive_src_url := "https://github.com/phar-io/phive/releases/download/0.15.3/phive-0.15.3.phar"
	phive_cmd_src := exec.Command("curl", "-L", phive_src_url, "-o", "source.tar.gz")
	err = phive_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phive_cmd_direct := exec.Command("./binary")
	err = phive_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
exec.Command("latte", "install", "php")
}
