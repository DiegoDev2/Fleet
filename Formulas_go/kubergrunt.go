package main

// Kubergrunt - Collection of commands to fill in the gaps between Terraform, Helm, and Kubectl
// Homepage: https://github.com/gruntwork-io/kubergrunt

import (
	"fmt"
	
	"os/exec"
)

func installKubergrunt() {
	// Método 1: Descargar y extraer .tar.gz
	kubergrunt_tar_url := "https://github.com/gruntwork-io/kubergrunt/archive/refs/tags/v0.16.0.tar.gz"
	kubergrunt_cmd_tar := exec.Command("curl", "-L", kubergrunt_tar_url, "-o", "package.tar.gz")
	err := kubergrunt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubergrunt_zip_url := "https://github.com/gruntwork-io/kubergrunt/archive/refs/tags/v0.16.0.zip"
	kubergrunt_cmd_zip := exec.Command("curl", "-L", kubergrunt_zip_url, "-o", "package.zip")
	err = kubergrunt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubergrunt_bin_url := "https://github.com/gruntwork-io/kubergrunt/archive/refs/tags/v0.16.0.bin"
	kubergrunt_cmd_bin := exec.Command("curl", "-L", kubergrunt_bin_url, "-o", "binary.bin")
	err = kubergrunt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubergrunt_src_url := "https://github.com/gruntwork-io/kubergrunt/archive/refs/tags/v0.16.0.src.tar.gz"
	kubergrunt_cmd_src := exec.Command("curl", "-L", kubergrunt_src_url, "-o", "source.tar.gz")
	err = kubergrunt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubergrunt_cmd_direct := exec.Command("./binary")
	err = kubergrunt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
