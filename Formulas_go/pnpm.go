package main

// Pnpm - Fast, disk space efficient package manager
// Homepage: https://pnpm.io/

import (
	"fmt"
	
	"os/exec"
)

func installPnpm() {
	// Método 1: Descargar y extraer .tar.gz
	pnpm_tar_url := "https://registry.npmjs.org/pnpm/-/pnpm-9.9.0.tgz"
	pnpm_cmd_tar := exec.Command("curl", "-L", pnpm_tar_url, "-o", "package.tar.gz")
	err := pnpm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pnpm_zip_url := "https://registry.npmjs.org/pnpm/-/pnpm-9.9.0.tgz"
	pnpm_cmd_zip := exec.Command("curl", "-L", pnpm_zip_url, "-o", "package.zip")
	err = pnpm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pnpm_bin_url := "https://registry.npmjs.org/pnpm/-/pnpm-9.9.0.tgz"
	pnpm_cmd_bin := exec.Command("curl", "-L", pnpm_bin_url, "-o", "binary.bin")
	err = pnpm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pnpm_src_url := "https://registry.npmjs.org/pnpm/-/pnpm-9.9.0.tgz"
	pnpm_cmd_src := exec.Command("curl", "-L", pnpm_src_url, "-o", "source.tar.gz")
	err = pnpm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pnpm_cmd_direct := exec.Command("./binary")
	err = pnpm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
