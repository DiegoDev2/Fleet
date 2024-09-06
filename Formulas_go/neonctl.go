package main

// Neonctl - Neon CLI tool
// Homepage: https://neon.tech/docs/reference/neon-cli

import (
	"fmt"
	
	"os/exec"
)

func installNeonctl() {
	// Método 1: Descargar y extraer .tar.gz
	neonctl_tar_url := "https://registry.npmjs.org/neonctl/-/neonctl-1.36.0.tgz"
	neonctl_cmd_tar := exec.Command("curl", "-L", neonctl_tar_url, "-o", "package.tar.gz")
	err := neonctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	neonctl_zip_url := "https://registry.npmjs.org/neonctl/-/neonctl-1.36.0.tgz"
	neonctl_cmd_zip := exec.Command("curl", "-L", neonctl_zip_url, "-o", "package.zip")
	err = neonctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	neonctl_bin_url := "https://registry.npmjs.org/neonctl/-/neonctl-1.36.0.tgz"
	neonctl_cmd_bin := exec.Command("curl", "-L", neonctl_bin_url, "-o", "binary.bin")
	err = neonctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	neonctl_src_url := "https://registry.npmjs.org/neonctl/-/neonctl-1.36.0.tgz"
	neonctl_cmd_src := exec.Command("curl", "-L", neonctl_src_url, "-o", "source.tar.gz")
	err = neonctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	neonctl_cmd_direct := exec.Command("./binary")
	err = neonctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
}
