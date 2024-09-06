package main

// Kubecolor - Colorize your kubectl output
// Homepage: https://kubecolor.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installKubecolor() {
	// Método 1: Descargar y extraer .tar.gz
	kubecolor_tar_url := "https://github.com/kubecolor/kubecolor/archive/refs/tags/v0.4.0.tar.gz"
	kubecolor_cmd_tar := exec.Command("curl", "-L", kubecolor_tar_url, "-o", "package.tar.gz")
	err := kubecolor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubecolor_zip_url := "https://github.com/kubecolor/kubecolor/archive/refs/tags/v0.4.0.zip"
	kubecolor_cmd_zip := exec.Command("curl", "-L", kubecolor_zip_url, "-o", "package.zip")
	err = kubecolor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubecolor_bin_url := "https://github.com/kubecolor/kubecolor/archive/refs/tags/v0.4.0.bin"
	kubecolor_cmd_bin := exec.Command("curl", "-L", kubecolor_bin_url, "-o", "binary.bin")
	err = kubecolor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubecolor_src_url := "https://github.com/kubecolor/kubecolor/archive/refs/tags/v0.4.0.src.tar.gz"
	kubecolor_cmd_src := exec.Command("curl", "-L", kubecolor_src_url, "-o", "source.tar.gz")
	err = kubecolor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubecolor_cmd_direct := exec.Command("./binary")
	err = kubecolor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: kubernetes-cli")
exec.Command("latte", "install", "kubernetes-cli")
}
