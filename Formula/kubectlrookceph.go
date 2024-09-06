package main

// KubectlRookCeph - Rook plugin for Ceph management
// Homepage: https://rook.io/

import (
	"fmt"
	
	"os/exec"
)

func installKubectlRookCeph() {
	// Método 1: Descargar y extraer .tar.gz
	kubectlrookceph_tar_url := "https://github.com/rook/kubectl-rook-ceph/archive/refs/tags/v0.9.1.tar.gz"
	kubectlrookceph_cmd_tar := exec.Command("curl", "-L", kubectlrookceph_tar_url, "-o", "package.tar.gz")
	err := kubectlrookceph_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubectlrookceph_zip_url := "https://github.com/rook/kubectl-rook-ceph/archive/refs/tags/v0.9.1.zip"
	kubectlrookceph_cmd_zip := exec.Command("curl", "-L", kubectlrookceph_zip_url, "-o", "package.zip")
	err = kubectlrookceph_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubectlrookceph_bin_url := "https://github.com/rook/kubectl-rook-ceph/archive/refs/tags/v0.9.1.bin"
	kubectlrookceph_cmd_bin := exec.Command("curl", "-L", kubectlrookceph_bin_url, "-o", "binary.bin")
	err = kubectlrookceph_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubectlrookceph_src_url := "https://github.com/rook/kubectl-rook-ceph/archive/refs/tags/v0.9.1.src.tar.gz"
	kubectlrookceph_cmd_src := exec.Command("curl", "-L", kubectlrookceph_src_url, "-o", "source.tar.gz")
	err = kubectlrookceph_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubectlrookceph_cmd_direct := exec.Command("./binary")
	err = kubectlrookceph_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
