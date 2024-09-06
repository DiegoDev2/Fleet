package main

// Istioctl - Istio configuration command-line utility
// Homepage: https://istio.io/

import (
	"fmt"
	
	"os/exec"
)

func installIstioctl() {
	// Método 1: Descargar y extraer .tar.gz
	istioctl_tar_url := "https://github.com/istio/istio/archive/refs/tags/1.23.0.tar.gz"
	istioctl_cmd_tar := exec.Command("curl", "-L", istioctl_tar_url, "-o", "package.tar.gz")
	err := istioctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	istioctl_zip_url := "https://github.com/istio/istio/archive/refs/tags/1.23.0.zip"
	istioctl_cmd_zip := exec.Command("curl", "-L", istioctl_zip_url, "-o", "package.zip")
	err = istioctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	istioctl_bin_url := "https://github.com/istio/istio/archive/refs/tags/1.23.0.bin"
	istioctl_cmd_bin := exec.Command("curl", "-L", istioctl_bin_url, "-o", "binary.bin")
	err = istioctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	istioctl_src_url := "https://github.com/istio/istio/archive/refs/tags/1.23.0.src.tar.gz"
	istioctl_cmd_src := exec.Command("curl", "-L", istioctl_src_url, "-o", "source.tar.gz")
	err = istioctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	istioctl_cmd_direct := exec.Command("./binary")
	err = istioctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
