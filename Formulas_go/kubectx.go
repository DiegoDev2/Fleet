package main

// Kubectx - Tool that can switch between kubectl contexts easily and create aliases
// Homepage: https://github.com/ahmetb/kubectx

import (
	"fmt"
	
	"os/exec"
)

func installKubectx() {
	// Método 1: Descargar y extraer .tar.gz
	kubectx_tar_url := "https://github.com/ahmetb/kubectx/archive/refs/tags/v0.9.5.tar.gz"
	kubectx_cmd_tar := exec.Command("curl", "-L", kubectx_tar_url, "-o", "package.tar.gz")
	err := kubectx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubectx_zip_url := "https://github.com/ahmetb/kubectx/archive/refs/tags/v0.9.5.zip"
	kubectx_cmd_zip := exec.Command("curl", "-L", kubectx_zip_url, "-o", "package.zip")
	err = kubectx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubectx_bin_url := "https://github.com/ahmetb/kubectx/archive/refs/tags/v0.9.5.bin"
	kubectx_cmd_bin := exec.Command("curl", "-L", kubectx_bin_url, "-o", "binary.bin")
	err = kubectx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubectx_src_url := "https://github.com/ahmetb/kubectx/archive/refs/tags/v0.9.5.src.tar.gz"
	kubectx_cmd_src := exec.Command("curl", "-L", kubectx_src_url, "-o", "source.tar.gz")
	err = kubectx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubectx_cmd_direct := exec.Command("./binary")
	err = kubectx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: kubernetes-cli")
exec.Command("latte", "install", "kubernetes-cli")
}
