package main

// Kustomize - Template-free customization of Kubernetes YAML manifests
// Homepage: https://github.com/kubernetes-sigs/kustomize

import (
	"fmt"
	
	"os/exec"
)

func installKustomize() {
	// Método 1: Descargar y extraer .tar.gz
	kustomize_tar_url := "https://github.com/kubernetes-sigs/kustomize/archive/refs/tags/kustomize/v5.4.3.tar.gz"
	kustomize_cmd_tar := exec.Command("curl", "-L", kustomize_tar_url, "-o", "package.tar.gz")
	err := kustomize_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kustomize_zip_url := "https://github.com/kubernetes-sigs/kustomize/archive/refs/tags/kustomize/v5.4.3.zip"
	kustomize_cmd_zip := exec.Command("curl", "-L", kustomize_zip_url, "-o", "package.zip")
	err = kustomize_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kustomize_bin_url := "https://github.com/kubernetes-sigs/kustomize/archive/refs/tags/kustomize/v5.4.3.bin"
	kustomize_cmd_bin := exec.Command("curl", "-L", kustomize_bin_url, "-o", "binary.bin")
	err = kustomize_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kustomize_src_url := "https://github.com/kubernetes-sigs/kustomize/archive/refs/tags/kustomize/v5.4.3.src.tar.gz"
	kustomize_cmd_src := exec.Command("curl", "-L", kustomize_src_url, "-o", "source.tar.gz")
	err = kustomize_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kustomize_cmd_direct := exec.Command("./binary")
	err = kustomize_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
