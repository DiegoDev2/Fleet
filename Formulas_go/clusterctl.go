package main

// Clusterctl - Home for the Cluster Management API work, a subproject of sig-cluster-lifecycle
// Homepage: https://cluster-api.sigs.k8s.io

import (
	"fmt"
	
	"os/exec"
)

func installClusterctl() {
	// Método 1: Descargar y extraer .tar.gz
	clusterctl_tar_url := "https://github.com/kubernetes-sigs/cluster-api/archive/refs/tags/v1.8.2.tar.gz"
	clusterctl_cmd_tar := exec.Command("curl", "-L", clusterctl_tar_url, "-o", "package.tar.gz")
	err := clusterctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clusterctl_zip_url := "https://github.com/kubernetes-sigs/cluster-api/archive/refs/tags/v1.8.2.zip"
	clusterctl_cmd_zip := exec.Command("curl", "-L", clusterctl_zip_url, "-o", "package.zip")
	err = clusterctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clusterctl_bin_url := "https://github.com/kubernetes-sigs/cluster-api/archive/refs/tags/v1.8.2.bin"
	clusterctl_cmd_bin := exec.Command("curl", "-L", clusterctl_bin_url, "-o", "binary.bin")
	err = clusterctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clusterctl_src_url := "https://github.com/kubernetes-sigs/cluster-api/archive/refs/tags/v1.8.2.src.tar.gz"
	clusterctl_cmd_src := exec.Command("curl", "-L", clusterctl_src_url, "-o", "source.tar.gz")
	err = clusterctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clusterctl_cmd_direct := exec.Command("./binary")
	err = clusterctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
