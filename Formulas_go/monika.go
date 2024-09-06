package main

// Monika - Synthetic monitoring made easy
// Homepage: https://monika.hyperjump.tech

import (
	"fmt"
	
	"os/exec"
)

func installMonika() {
	// Método 1: Descargar y extraer .tar.gz
	monika_tar_url := "https://registry.npmjs.org/@hyperjumptech/monika/-/monika-1.21.0.tgz"
	monika_cmd_tar := exec.Command("curl", "-L", monika_tar_url, "-o", "package.tar.gz")
	err := monika_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	monika_zip_url := "https://registry.npmjs.org/@hyperjumptech/monika/-/monika-1.21.0.tgz"
	monika_cmd_zip := exec.Command("curl", "-L", monika_zip_url, "-o", "package.zip")
	err = monika_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	monika_bin_url := "https://registry.npmjs.org/@hyperjumptech/monika/-/monika-1.21.0.tgz"
	monika_cmd_bin := exec.Command("curl", "-L", monika_bin_url, "-o", "binary.bin")
	err = monika_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	monika_src_url := "https://registry.npmjs.org/@hyperjumptech/monika/-/monika-1.21.0.tgz"
	monika_cmd_src := exec.Command("curl", "-L", monika_src_url, "-o", "source.tar.gz")
	err = monika_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	monika_cmd_direct := exec.Command("./binary")
	err = monika_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
}
