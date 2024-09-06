package main

// Pwned - CLI for the 'Have I been pwned?' service
// Homepage: https://github.com/wKovacs64/pwned

import (
	"fmt"
	
	"os/exec"
)

func installPwned() {
	// Método 1: Descargar y extraer .tar.gz
	pwned_tar_url := "https://registry.npmjs.org/pwned/-/pwned-12.1.1.tgz"
	pwned_cmd_tar := exec.Command("curl", "-L", pwned_tar_url, "-o", "package.tar.gz")
	err := pwned_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pwned_zip_url := "https://registry.npmjs.org/pwned/-/pwned-12.1.1.tgz"
	pwned_cmd_zip := exec.Command("curl", "-L", pwned_zip_url, "-o", "package.zip")
	err = pwned_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pwned_bin_url := "https://registry.npmjs.org/pwned/-/pwned-12.1.1.tgz"
	pwned_cmd_bin := exec.Command("curl", "-L", pwned_bin_url, "-o", "binary.bin")
	err = pwned_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pwned_src_url := "https://registry.npmjs.org/pwned/-/pwned-12.1.1.tgz"
	pwned_cmd_src := exec.Command("curl", "-L", pwned_src_url, "-o", "source.tar.gz")
	err = pwned_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pwned_cmd_direct := exec.Command("./binary")
	err = pwned_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
