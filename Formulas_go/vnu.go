package main

// Vnu - Nu Markup Checker: command-line and server HTML validator
// Homepage: https://validator.github.io/validator/

import (
	"fmt"
	
	"os/exec"
)

func installVnu() {
	// Método 1: Descargar y extraer .tar.gz
	vnu_tar_url := "https://registry.npmjs.org/vnu-jar/-/vnu-jar-23.4.11.tgz"
	vnu_cmd_tar := exec.Command("curl", "-L", vnu_tar_url, "-o", "package.tar.gz")
	err := vnu_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vnu_zip_url := "https://registry.npmjs.org/vnu-jar/-/vnu-jar-23.4.11.tgz"
	vnu_cmd_zip := exec.Command("curl", "-L", vnu_zip_url, "-o", "package.zip")
	err = vnu_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vnu_bin_url := "https://registry.npmjs.org/vnu-jar/-/vnu-jar-23.4.11.tgz"
	vnu_cmd_bin := exec.Command("curl", "-L", vnu_bin_url, "-o", "binary.bin")
	err = vnu_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vnu_src_url := "https://registry.npmjs.org/vnu-jar/-/vnu-jar-23.4.11.tgz"
	vnu_cmd_src := exec.Command("curl", "-L", vnu_src_url, "-o", "source.tar.gz")
	err = vnu_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vnu_cmd_direct := exec.Command("./binary")
	err = vnu_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
