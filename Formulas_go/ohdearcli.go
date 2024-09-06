package main

// OhdearCli - Tool to manage your Oh Dear sites
// Homepage: https://github.com/ohdearapp/ohdear-cli

import (
	"fmt"
	
	"os/exec"
)

func installOhdearCli() {
	// Método 1: Descargar y extraer .tar.gz
	ohdearcli_tar_url := "https://github.com/ohdearapp/ohdear-cli/releases/download/v4.3.0/ohdear.phar"
	ohdearcli_cmd_tar := exec.Command("curl", "-L", ohdearcli_tar_url, "-o", "package.tar.gz")
	err := ohdearcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ohdearcli_zip_url := "https://github.com/ohdearapp/ohdear-cli/releases/download/v4.3.0/ohdear.phar"
	ohdearcli_cmd_zip := exec.Command("curl", "-L", ohdearcli_zip_url, "-o", "package.zip")
	err = ohdearcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ohdearcli_bin_url := "https://github.com/ohdearapp/ohdear-cli/releases/download/v4.3.0/ohdear.phar"
	ohdearcli_cmd_bin := exec.Command("curl", "-L", ohdearcli_bin_url, "-o", "binary.bin")
	err = ohdearcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ohdearcli_src_url := "https://github.com/ohdearapp/ohdear-cli/releases/download/v4.3.0/ohdear.phar"
	ohdearcli_cmd_src := exec.Command("curl", "-L", ohdearcli_src_url, "-o", "source.tar.gz")
	err = ohdearcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ohdearcli_cmd_direct := exec.Command("./binary")
	err = ohdearcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: php")
exec.Command("latte", "install", "php")
}
