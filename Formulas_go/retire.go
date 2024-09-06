package main

// Retire - Scanner detecting the use of JavaScript libraries with known vulnerabilities
// Homepage: https://retirejs.github.io/retire.js/

import (
	"fmt"
	
	"os/exec"
)

func installRetire() {
	// Método 1: Descargar y extraer .tar.gz
	retire_tar_url := "https://registry.npmjs.org/retire/-/retire-5.2.2.tgz"
	retire_cmd_tar := exec.Command("curl", "-L", retire_tar_url, "-o", "package.tar.gz")
	err := retire_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	retire_zip_url := "https://registry.npmjs.org/retire/-/retire-5.2.2.tgz"
	retire_cmd_zip := exec.Command("curl", "-L", retire_zip_url, "-o", "package.zip")
	err = retire_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	retire_bin_url := "https://registry.npmjs.org/retire/-/retire-5.2.2.tgz"
	retire_cmd_bin := exec.Command("curl", "-L", retire_bin_url, "-o", "binary.bin")
	err = retire_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	retire_src_url := "https://registry.npmjs.org/retire/-/retire-5.2.2.tgz"
	retire_cmd_src := exec.Command("curl", "-L", retire_src_url, "-o", "source.tar.gz")
	err = retire_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	retire_cmd_direct := exec.Command("./binary")
	err = retire_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
