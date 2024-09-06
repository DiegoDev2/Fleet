package main

// Kubecm - KubeConfig Manager
// Homepage: https://kubecm.cloud

import (
	"fmt"
	
	"os/exec"
)

func installKubecm() {
	// Método 1: Descargar y extraer .tar.gz
	kubecm_tar_url := "https://github.com/sunny0826/kubecm/archive/refs/tags/v0.31.0.tar.gz"
	kubecm_cmd_tar := exec.Command("curl", "-L", kubecm_tar_url, "-o", "package.tar.gz")
	err := kubecm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubecm_zip_url := "https://github.com/sunny0826/kubecm/archive/refs/tags/v0.31.0.zip"
	kubecm_cmd_zip := exec.Command("curl", "-L", kubecm_zip_url, "-o", "package.zip")
	err = kubecm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubecm_bin_url := "https://github.com/sunny0826/kubecm/archive/refs/tags/v0.31.0.bin"
	kubecm_cmd_bin := exec.Command("curl", "-L", kubecm_bin_url, "-o", "binary.bin")
	err = kubecm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubecm_src_url := "https://github.com/sunny0826/kubecm/archive/refs/tags/v0.31.0.src.tar.gz"
	kubecm_cmd_src := exec.Command("curl", "-L", kubecm_src_url, "-o", "source.tar.gz")
	err = kubecm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubecm_cmd_direct := exec.Command("./binary")
	err = kubecm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
