package main

// KubePs1 - Kubernetes prompt info for bash and zsh
// Homepage: https://github.com/jonmosco/kube-ps1

import (
	"fmt"
	
	"os/exec"
)

func installKubePs1() {
	// Método 1: Descargar y extraer .tar.gz
	kubeps1_tar_url := "https://github.com/jonmosco/kube-ps1/archive/refs/tags/v0.9.0.tar.gz"
	kubeps1_cmd_tar := exec.Command("curl", "-L", kubeps1_tar_url, "-o", "package.tar.gz")
	err := kubeps1_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubeps1_zip_url := "https://github.com/jonmosco/kube-ps1/archive/refs/tags/v0.9.0.zip"
	kubeps1_cmd_zip := exec.Command("curl", "-L", kubeps1_zip_url, "-o", "package.zip")
	err = kubeps1_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubeps1_bin_url := "https://github.com/jonmosco/kube-ps1/archive/refs/tags/v0.9.0.bin"
	kubeps1_cmd_bin := exec.Command("curl", "-L", kubeps1_bin_url, "-o", "binary.bin")
	err = kubeps1_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubeps1_src_url := "https://github.com/jonmosco/kube-ps1/archive/refs/tags/v0.9.0.src.tar.gz"
	kubeps1_cmd_src := exec.Command("curl", "-L", kubeps1_src_url, "-o", "source.tar.gz")
	err = kubeps1_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubeps1_cmd_direct := exec.Command("./binary")
	err = kubeps1_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: kubernetes-cli")
exec.Command("latte", "install", "kubernetes-cli")
}
