package main

// Kops - Production Grade K8s Installation, Upgrades, and Management
// Homepage: https://kops.sigs.k8s.io/

import (
	"fmt"
	
	"os/exec"
)

func installKops() {
	// Método 1: Descargar y extraer .tar.gz
	kops_tar_url := "https://github.com/kubernetes/kops/archive/refs/tags/v1.30.0.tar.gz"
	kops_cmd_tar := exec.Command("curl", "-L", kops_tar_url, "-o", "package.tar.gz")
	err := kops_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kops_zip_url := "https://github.com/kubernetes/kops/archive/refs/tags/v1.30.0.zip"
	kops_cmd_zip := exec.Command("curl", "-L", kops_zip_url, "-o", "package.zip")
	err = kops_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kops_bin_url := "https://github.com/kubernetes/kops/archive/refs/tags/v1.30.0.bin"
	kops_cmd_bin := exec.Command("curl", "-L", kops_bin_url, "-o", "binary.bin")
	err = kops_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kops_src_url := "https://github.com/kubernetes/kops/archive/refs/tags/v1.30.0.src.tar.gz"
	kops_cmd_src := exec.Command("curl", "-L", kops_src_url, "-o", "source.tar.gz")
	err = kops_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kops_cmd_direct := exec.Command("./binary")
	err = kops_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: kubernetes-cli")
	exec.Command("latte", "install", "kubernetes-cli").Run()
}
