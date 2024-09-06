package main

// CriTools - CLI and validation tools for Kubelet Container Runtime Interface (CRI)
// Homepage: https://github.com/kubernetes-sigs/cri-tools

import (
	"fmt"
	
	"os/exec"
)

func installCriTools() {
	// Método 1: Descargar y extraer .tar.gz
	critools_tar_url := "https://github.com/kubernetes-sigs/cri-tools/archive/refs/tags/v1.31.1.tar.gz"
	critools_cmd_tar := exec.Command("curl", "-L", critools_tar_url, "-o", "package.tar.gz")
	err := critools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	critools_zip_url := "https://github.com/kubernetes-sigs/cri-tools/archive/refs/tags/v1.31.1.zip"
	critools_cmd_zip := exec.Command("curl", "-L", critools_zip_url, "-o", "package.zip")
	err = critools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	critools_bin_url := "https://github.com/kubernetes-sigs/cri-tools/archive/refs/tags/v1.31.1.bin"
	critools_cmd_bin := exec.Command("curl", "-L", critools_bin_url, "-o", "binary.bin")
	err = critools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	critools_src_url := "https://github.com/kubernetes-sigs/cri-tools/archive/refs/tags/v1.31.1.src.tar.gz"
	critools_cmd_src := exec.Command("curl", "-L", critools_src_url, "-o", "source.tar.gz")
	err = critools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	critools_cmd_direct := exec.Command("./binary")
	err = critools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
