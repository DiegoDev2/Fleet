package main

// Devspace - CLI helps develop/deploy/debug apps with Docker and k8s
// Homepage: https://devspace.sh/

import (
	"fmt"
	
	"os/exec"
)

func installDevspace() {
	// Método 1: Descargar y extraer .tar.gz
	devspace_tar_url := "https://github.com/devspace-sh/devspace/archive/refs/tags/v6.3.12.tar.gz"
	devspace_cmd_tar := exec.Command("curl", "-L", devspace_tar_url, "-o", "package.tar.gz")
	err := devspace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	devspace_zip_url := "https://github.com/devspace-sh/devspace/archive/refs/tags/v6.3.12.zip"
	devspace_cmd_zip := exec.Command("curl", "-L", devspace_zip_url, "-o", "package.zip")
	err = devspace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	devspace_bin_url := "https://github.com/devspace-sh/devspace/archive/refs/tags/v6.3.12.bin"
	devspace_cmd_bin := exec.Command("curl", "-L", devspace_bin_url, "-o", "binary.bin")
	err = devspace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	devspace_src_url := "https://github.com/devspace-sh/devspace/archive/refs/tags/v6.3.12.src.tar.gz"
	devspace_cmd_src := exec.Command("curl", "-L", devspace_src_url, "-o", "source.tar.gz")
	err = devspace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	devspace_cmd_direct := exec.Command("./binary")
	err = devspace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: kubernetes-cli")
exec.Command("latte", "install", "kubernetes-cli")
}
