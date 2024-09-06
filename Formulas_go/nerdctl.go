package main

// Nerdctl - ContaiNERD CTL - Docker-compatible CLI for containerd
// Homepage: https://github.com/containerd/nerdctl

import (
	"fmt"
	
	"os/exec"
)

func installNerdctl() {
	// Método 1: Descargar y extraer .tar.gz
	nerdctl_tar_url := "https://github.com/containerd/nerdctl/archive/refs/tags/v1.7.6.tar.gz"
	nerdctl_cmd_tar := exec.Command("curl", "-L", nerdctl_tar_url, "-o", "package.tar.gz")
	err := nerdctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nerdctl_zip_url := "https://github.com/containerd/nerdctl/archive/refs/tags/v1.7.6.zip"
	nerdctl_cmd_zip := exec.Command("curl", "-L", nerdctl_zip_url, "-o", "package.zip")
	err = nerdctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nerdctl_bin_url := "https://github.com/containerd/nerdctl/archive/refs/tags/v1.7.6.bin"
	nerdctl_cmd_bin := exec.Command("curl", "-L", nerdctl_bin_url, "-o", "binary.bin")
	err = nerdctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nerdctl_src_url := "https://github.com/containerd/nerdctl/archive/refs/tags/v1.7.6.src.tar.gz"
	nerdctl_cmd_src := exec.Command("curl", "-L", nerdctl_src_url, "-o", "source.tar.gz")
	err = nerdctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nerdctl_cmd_direct := exec.Command("./binary")
	err = nerdctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
