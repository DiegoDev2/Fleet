package main

// Cyctl - Customizable UI for Kubernetes workloads
// Homepage: https://cyclops-ui.com/

import (
	"fmt"
	
	"os/exec"
)

func installCyctl() {
	// Método 1: Descargar y extraer .tar.gz
	cyctl_tar_url := "https://github.com/cyclops-ui/cyclops/archive/refs/tags/v0.10.0.tar.gz"
	cyctl_cmd_tar := exec.Command("curl", "-L", cyctl_tar_url, "-o", "package.tar.gz")
	err := cyctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cyctl_zip_url := "https://github.com/cyclops-ui/cyclops/archive/refs/tags/v0.10.0.zip"
	cyctl_cmd_zip := exec.Command("curl", "-L", cyctl_zip_url, "-o", "package.zip")
	err = cyctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cyctl_bin_url := "https://github.com/cyclops-ui/cyclops/archive/refs/tags/v0.10.0.bin"
	cyctl_cmd_bin := exec.Command("curl", "-L", cyctl_bin_url, "-o", "binary.bin")
	err = cyctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cyctl_src_url := "https://github.com/cyclops-ui/cyclops/archive/refs/tags/v0.10.0.src.tar.gz"
	cyctl_cmd_src := exec.Command("curl", "-L", cyctl_src_url, "-o", "source.tar.gz")
	err = cyctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cyctl_cmd_direct := exec.Command("./binary")
	err = cyctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
