package main

// Ingress2gateway - Convert Kubernetes Ingress resources to Kubernetes Gateway API resources
// Homepage: https://github.com/kubernetes-sigs/ingress2gateway

import (
	"fmt"
	
	"os/exec"
)

func installIngress2gateway() {
	// Método 1: Descargar y extraer .tar.gz
	ingress2gateway_tar_url := "https://github.com/kubernetes-sigs/ingress2gateway/archive/refs/tags/v0.3.0.tar.gz"
	ingress2gateway_cmd_tar := exec.Command("curl", "-L", ingress2gateway_tar_url, "-o", "package.tar.gz")
	err := ingress2gateway_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ingress2gateway_zip_url := "https://github.com/kubernetes-sigs/ingress2gateway/archive/refs/tags/v0.3.0.zip"
	ingress2gateway_cmd_zip := exec.Command("curl", "-L", ingress2gateway_zip_url, "-o", "package.zip")
	err = ingress2gateway_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ingress2gateway_bin_url := "https://github.com/kubernetes-sigs/ingress2gateway/archive/refs/tags/v0.3.0.bin"
	ingress2gateway_cmd_bin := exec.Command("curl", "-L", ingress2gateway_bin_url, "-o", "binary.bin")
	err = ingress2gateway_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ingress2gateway_src_url := "https://github.com/kubernetes-sigs/ingress2gateway/archive/refs/tags/v0.3.0.src.tar.gz"
	ingress2gateway_cmd_src := exec.Command("curl", "-L", ingress2gateway_src_url, "-o", "source.tar.gz")
	err = ingress2gateway_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ingress2gateway_cmd_direct := exec.Command("./binary")
	err = ingress2gateway_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
