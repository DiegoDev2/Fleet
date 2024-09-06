package main

// Kubeconform - FAST Kubernetes manifests validator, with support for Custom Resources!
// Homepage: https://github.com/yannh/kubeconform

import (
	"fmt"
	
	"os/exec"
)

func installKubeconform() {
	// Método 1: Descargar y extraer .tar.gz
	kubeconform_tar_url := "https://github.com/yannh/kubeconform/archive/refs/tags/v0.6.7.tar.gz"
	kubeconform_cmd_tar := exec.Command("curl", "-L", kubeconform_tar_url, "-o", "package.tar.gz")
	err := kubeconform_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubeconform_zip_url := "https://github.com/yannh/kubeconform/archive/refs/tags/v0.6.7.zip"
	kubeconform_cmd_zip := exec.Command("curl", "-L", kubeconform_zip_url, "-o", "package.zip")
	err = kubeconform_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubeconform_bin_url := "https://github.com/yannh/kubeconform/archive/refs/tags/v0.6.7.bin"
	kubeconform_cmd_bin := exec.Command("curl", "-L", kubeconform_bin_url, "-o", "binary.bin")
	err = kubeconform_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubeconform_src_url := "https://github.com/yannh/kubeconform/archive/refs/tags/v0.6.7.src.tar.gz"
	kubeconform_cmd_src := exec.Command("curl", "-L", kubeconform_src_url, "-o", "source.tar.gz")
	err = kubeconform_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubeconform_cmd_direct := exec.Command("./binary")
	err = kubeconform_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
