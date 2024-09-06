package main

// Renovate - Automated dependency updates. Flexible so you don't need to be
// Homepage: https://github.com/renovatebot/renovate

import (
	"fmt"
	
	"os/exec"
)

func installRenovate() {
	// Método 1: Descargar y extraer .tar.gz
	renovate_tar_url := "https://registry.npmjs.org/renovate/-/renovate-38.68.0.tgz"
	renovate_cmd_tar := exec.Command("curl", "-L", renovate_tar_url, "-o", "package.tar.gz")
	err := renovate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	renovate_zip_url := "https://registry.npmjs.org/renovate/-/renovate-38.68.0.tgz"
	renovate_cmd_zip := exec.Command("curl", "-L", renovate_zip_url, "-o", "package.zip")
	err = renovate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	renovate_bin_url := "https://registry.npmjs.org/renovate/-/renovate-38.68.0.tgz"
	renovate_cmd_bin := exec.Command("curl", "-L", renovate_bin_url, "-o", "binary.bin")
	err = renovate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	renovate_src_url := "https://registry.npmjs.org/renovate/-/renovate-38.68.0.tgz"
	renovate_cmd_src := exec.Command("curl", "-L", renovate_src_url, "-o", "source.tar.gz")
	err = renovate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	renovate_cmd_direct := exec.Command("./binary")
	err = renovate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node@20")
	exec.Command("latte", "install", "node@20").Run()
}
