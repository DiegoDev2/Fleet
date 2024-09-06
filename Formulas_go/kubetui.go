package main

// Kubetui - TUI tool for monitoring and exploration of Kubernetes resources
// Homepage: https://github.com/sarub0b0/kubetui

import (
	"fmt"
	
	"os/exec"
)

func installKubetui() {
	// Método 1: Descargar y extraer .tar.gz
	kubetui_tar_url := "https://github.com/sarub0b0/kubetui/archive/refs/tags/v1.5.3.tar.gz"
	kubetui_cmd_tar := exec.Command("curl", "-L", kubetui_tar_url, "-o", "package.tar.gz")
	err := kubetui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kubetui_zip_url := "https://github.com/sarub0b0/kubetui/archive/refs/tags/v1.5.3.zip"
	kubetui_cmd_zip := exec.Command("curl", "-L", kubetui_zip_url, "-o", "package.zip")
	err = kubetui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kubetui_bin_url := "https://github.com/sarub0b0/kubetui/archive/refs/tags/v1.5.3.bin"
	kubetui_cmd_bin := exec.Command("curl", "-L", kubetui_bin_url, "-o", "binary.bin")
	err = kubetui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kubetui_src_url := "https://github.com/sarub0b0/kubetui/archive/refs/tags/v1.5.3.src.tar.gz"
	kubetui_cmd_src := exec.Command("curl", "-L", kubetui_src_url, "-o", "source.tar.gz")
	err = kubetui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kubetui_cmd_direct := exec.Command("./binary")
	err = kubetui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
