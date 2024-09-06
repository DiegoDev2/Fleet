package main

// Kubecfg - Manage complex enterprise Kubernetes environments as code
// Homepage: https://github.com/kubecfg/kubecfg

import (
	"fmt"
	
	"os/exec"
)

func installKubecfg() {
	// Método 1: Descargar y extraer .tar.gz
	kubecfg_tar_url := "https://github.com/kubecfg/kubecfg/archive/refs/tags/v0.34.3.tar.gz"
	kubecfg_cmd_tar := exec.Command("curl", "-L", kubecfg_tar_url, "-o", "package.tar.gz")
	err := kubecfg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubecfg_zip_url := "https://github.com/kubecfg/kubecfg/archive/refs/tags/v0.34.3.zip"
	kubecfg_cmd_zip := exec.Command("curl", "-L", kubecfg_zip_url, "-o", "package.zip")
	err = kubecfg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubecfg_bin_url := "https://github.com/kubecfg/kubecfg/archive/refs/tags/v0.34.3.bin"
	kubecfg_cmd_bin := exec.Command("curl", "-L", kubecfg_bin_url, "-o", "binary.bin")
	err = kubecfg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubecfg_src_url := "https://github.com/kubecfg/kubecfg/archive/refs/tags/v0.34.3.src.tar.gz"
	kubecfg_cmd_src := exec.Command("curl", "-L", kubecfg_src_url, "-o", "source.tar.gz")
	err = kubecfg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubecfg_cmd_direct := exec.Command("./binary")
	err = kubecfg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
