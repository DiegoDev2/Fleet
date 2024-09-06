package main

// Kwok - Kubernetes WithOut Kubelet - Simulates thousands of Nodes and Clusters
// Homepage: https://kwok.sigs.k8s.io

import (
	"fmt"
	
	"os/exec"
)

func installKwok() {
	// Método 1: Descargar y extraer .tar.gz
	kwok_tar_url := "https://github.com/kubernetes-sigs/kwok/archive/refs/tags/v0.6.0.tar.gz"
	kwok_cmd_tar := exec.Command("curl", "-L", kwok_tar_url, "-o", "package.tar.gz")
	err := kwok_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kwok_zip_url := "https://github.com/kubernetes-sigs/kwok/archive/refs/tags/v0.6.0.zip"
	kwok_cmd_zip := exec.Command("curl", "-L", kwok_zip_url, "-o", "package.zip")
	err = kwok_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kwok_bin_url := "https://github.com/kubernetes-sigs/kwok/archive/refs/tags/v0.6.0.bin"
	kwok_cmd_bin := exec.Command("curl", "-L", kwok_bin_url, "-o", "binary.bin")
	err = kwok_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kwok_src_url := "https://github.com/kubernetes-sigs/kwok/archive/refs/tags/v0.6.0.src.tar.gz"
	kwok_cmd_src := exec.Command("curl", "-L", kwok_src_url, "-o", "source.tar.gz")
	err = kwok_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kwok_cmd_direct := exec.Command("./binary")
	err = kwok_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: docker")
	exec.Command("latte", "install", "docker").Run()
}
