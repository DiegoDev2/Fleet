package main

// K8sgpt - Scanning your k8s clusters, diagnosing, and triaging issues in simple English
// Homepage: https://k8sgpt.ai/

import (
	"fmt"
	
	"os/exec"
)

func installK8sgpt() {
	// Método 1: Descargar y extraer .tar.gz
	k8sgpt_tar_url := "https://github.com/k8sgpt-ai/k8sgpt/archive/refs/tags/v0.3.40.tar.gz"
	k8sgpt_cmd_tar := exec.Command("curl", "-L", k8sgpt_tar_url, "-o", "package.tar.gz")
	err := k8sgpt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	k8sgpt_zip_url := "https://github.com/k8sgpt-ai/k8sgpt/archive/refs/tags/v0.3.40.zip"
	k8sgpt_cmd_zip := exec.Command("curl", "-L", k8sgpt_zip_url, "-o", "package.zip")
	err = k8sgpt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	k8sgpt_bin_url := "https://github.com/k8sgpt-ai/k8sgpt/archive/refs/tags/v0.3.40.bin"
	k8sgpt_cmd_bin := exec.Command("curl", "-L", k8sgpt_bin_url, "-o", "binary.bin")
	err = k8sgpt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	k8sgpt_src_url := "https://github.com/k8sgpt-ai/k8sgpt/archive/refs/tags/v0.3.40.src.tar.gz"
	k8sgpt_cmd_src := exec.Command("curl", "-L", k8sgpt_src_url, "-o", "source.tar.gz")
	err = k8sgpt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	k8sgpt_cmd_direct := exec.Command("./binary")
	err = k8sgpt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
