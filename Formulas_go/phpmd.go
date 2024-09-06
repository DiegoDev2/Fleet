package main

// Phpmd - PHP Mess Detector
// Homepage: https://phpmd.org

import (
	"fmt"
	
	"os/exec"
)

func installPhpmd() {
	// Método 1: Descargar y extraer .tar.gz
	phpmd_tar_url := "https://github.com/phpmd/phpmd/releases/download/2.15.0/phpmd.phar"
	phpmd_cmd_tar := exec.Command("curl", "-L", phpmd_tar_url, "-o", "package.tar.gz")
	err := phpmd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	phpmd_zip_url := "https://github.com/phpmd/phpmd/releases/download/2.15.0/phpmd.phar"
	phpmd_cmd_zip := exec.Command("curl", "-L", phpmd_zip_url, "-o", "package.zip")
	err = phpmd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	phpmd_bin_url := "https://github.com/phpmd/phpmd/releases/download/2.15.0/phpmd.phar"
	phpmd_cmd_bin := exec.Command("curl", "-L", phpmd_bin_url, "-o", "binary.bin")
	err = phpmd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	phpmd_src_url := "https://github.com/phpmd/phpmd/releases/download/2.15.0/phpmd.phar"
	phpmd_cmd_src := exec.Command("curl", "-L", phpmd_src_url, "-o", "source.tar.gz")
	err = phpmd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	phpmd_cmd_direct := exec.Command("./binary")
	err = phpmd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
exec.Command("latte", "install", "php")
}
