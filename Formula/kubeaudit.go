package main

// Kubeaudit - Helps audit your Kubernetes clusters against common security controls
// Homepage: https://github.com/Shopify/kubeaudit

import (
	"fmt"
	
	"os/exec"
)

func installKubeaudit() {
	// Método 1: Descargar y extraer .tar.gz
	kubeaudit_tar_url := "https://github.com/Shopify/kubeaudit/archive/refs/tags/v0.22.2.tar.gz"
	kubeaudit_cmd_tar := exec.Command("curl", "-L", kubeaudit_tar_url, "-o", "package.tar.gz")
	err := kubeaudit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubeaudit_zip_url := "https://github.com/Shopify/kubeaudit/archive/refs/tags/v0.22.2.zip"
	kubeaudit_cmd_zip := exec.Command("curl", "-L", kubeaudit_zip_url, "-o", "package.zip")
	err = kubeaudit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubeaudit_bin_url := "https://github.com/Shopify/kubeaudit/archive/refs/tags/v0.22.2.bin"
	kubeaudit_cmd_bin := exec.Command("curl", "-L", kubeaudit_bin_url, "-o", "binary.bin")
	err = kubeaudit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubeaudit_src_url := "https://github.com/Shopify/kubeaudit/archive/refs/tags/v0.22.2.src.tar.gz"
	kubeaudit_cmd_src := exec.Command("curl", "-L", kubeaudit_src_url, "-o", "source.tar.gz")
	err = kubeaudit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubeaudit_cmd_direct := exec.Command("./binary")
	err = kubeaudit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
