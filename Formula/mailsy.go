package main

// Mailsy - Quickly generate a temporary email address
// Homepage: https://github.com/BalliAsghar/Mailsy

import (
	"fmt"
	
	"os/exec"
)

func installMailsy() {
	// Método 1: Descargar y extraer .tar.gz
	mailsy_tar_url := "https://registry.npmjs.org/mailsy/-/mailsy-5.0.0.tgz"
	mailsy_cmd_tar := exec.Command("curl", "-L", mailsy_tar_url, "-o", "package.tar.gz")
	err := mailsy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mailsy_zip_url := "https://registry.npmjs.org/mailsy/-/mailsy-5.0.0.tgz"
	mailsy_cmd_zip := exec.Command("curl", "-L", mailsy_zip_url, "-o", "package.zip")
	err = mailsy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mailsy_bin_url := "https://registry.npmjs.org/mailsy/-/mailsy-5.0.0.tgz"
	mailsy_cmd_bin := exec.Command("curl", "-L", mailsy_bin_url, "-o", "binary.bin")
	err = mailsy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mailsy_src_url := "https://registry.npmjs.org/mailsy/-/mailsy-5.0.0.tgz"
	mailsy_cmd_src := exec.Command("curl", "-L", mailsy_src_url, "-o", "source.tar.gz")
	err = mailsy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mailsy_cmd_direct := exec.Command("./binary")
	err = mailsy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
