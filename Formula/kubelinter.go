package main

// KubeLinter - Static analysis tool for Kubernetes YAML files and Helm charts
// Homepage: https://github.com/stackrox/kube-linter

import (
	"fmt"
	
	"os/exec"
)

func installKubeLinter() {
	// Método 1: Descargar y extraer .tar.gz
	kubelinter_tar_url := "https://github.com/stackrox/kube-linter/archive/refs/tags/v0.6.8.tar.gz"
	kubelinter_cmd_tar := exec.Command("curl", "-L", kubelinter_tar_url, "-o", "package.tar.gz")
	err := kubelinter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubelinter_zip_url := "https://github.com/stackrox/kube-linter/archive/refs/tags/v0.6.8.zip"
	kubelinter_cmd_zip := exec.Command("curl", "-L", kubelinter_zip_url, "-o", "package.zip")
	err = kubelinter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubelinter_bin_url := "https://github.com/stackrox/kube-linter/archive/refs/tags/v0.6.8.bin"
	kubelinter_cmd_bin := exec.Command("curl", "-L", kubelinter_bin_url, "-o", "binary.bin")
	err = kubelinter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubelinter_src_url := "https://github.com/stackrox/kube-linter/archive/refs/tags/v0.6.8.src.tar.gz"
	kubelinter_cmd_src := exec.Command("curl", "-L", kubelinter_src_url, "-o", "source.tar.gz")
	err = kubelinter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubelinter_cmd_direct := exec.Command("./binary")
	err = kubelinter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
